package errors

import (
	"errors"
	"fmt"
)

// Sentinel is a sentinel error type that can be defined as a constant
type Sentinel string

func (e Sentinel) Error() string {
	return string(e)
}

// New returns a composite error including both the receiver and a provided
// error message.
func (e Sentinel) New(text string) error {
	return Compose(e, errors.New(text))
}

// Wrap returns a composite error including both the receiver and a provided
// error.
func (e Sentinel) Wrap(err error) error {
	return Compose(e, err)
}

// Errorf returns a composite error including both the receiver and a formated
// error
func (e Sentinel) Errorf(format string, a ...interface{}) error {
	return Compose(e, fmt.Errorf(format, a...))
}
