package errors_test

import (
	"testing"

	"github.com/maargenton/go-testpredicate/pkg/verify"

	"github.com/maargenton/go-errors"
)

var err1 = errors.New("err1")
var err2 = errors.New("err2")
var err3 = errors.New("err3")

func TestCompose(t *testing.T) {
	var err = errors.Compose(err1, err2)

	verify.That(t, err.Error()).Eq("err1: err2")
	verify.That(t, errors.Unwrap(err)).Eq(err2)
	verify.That(t, err).IsError(err1)
	verify.That(t, err).IsError(err2)
}

func TestComposeMultiple(t *testing.T) {
	var err = errors.Compose(err1, err2, err3)

	verify.That(t, err.Error()).Eq("err1: err2: err3")
	verify.That(t, err).IsError(err1)
	verify.That(t, err).IsError(err2)
	verify.That(t, err).IsError(err3)
}

type testError struct{}
type testError2 struct{}

func (*testError) Error() string  { return "testError" }
func (*testError2) Error() string { return "testError2" }

func TestComposeAs(t *testing.T) {
	var err = errors.Compose(&testError{}, &testError2{})
	var err1 *testError
	var err2 *testError2

	verify.That(t, errors.As(err, &err1)).IsTrue()
	verify.That(t, errors.As(err, &err2)).IsTrue()
}
