package errors

import (
	"errors"
	"fmt"
)

type compositeError struct {
	tag, err error
}

func (e *compositeError) Error() string {
	return fmt.Sprintf("%v: %v", e.tag, e.err)
}

func (e *compositeError) Unwrap() error {
	return e.err
}

func (e *compositeError) Is(err error) bool {
	return errors.Is(e.tag, err) || errors.Is(e.err, err)
}

func (e *compositeError) As(target interface{}) bool {
	return errors.As(e.tag, target) || errors.As(e.err, target)
}

// Compose takes two or more errors and return a composite error representing
// both or all of them.
func Compose(tag, err error, errs ...error) error {

	if len(errs) == 0 {
		return &compositeError{
			tag: tag,
			err: err,
		}
	}
	return &compositeError{
		tag: tag,
		err: Compose(err, errs[0], errs[1:]...),
	}
}
