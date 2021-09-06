package errors_test

import (
	"testing"

	"github.com/maargenton/go-testpredicate/pkg/verify"

	"github.com/maargenton/go-errors"
)

func TestSentinel(t *testing.T) {
	const MyError errors.Sentinel = "MyError"
	verify.That(t, MyError.Error()).Eq("MyError")
}

func TestSentinelValueIdentity(t *testing.T) {
	const MyError errors.Sentinel = "MyError"
	const MyErrorCpy errors.Sentinel = "MyError"

	verify.That(t, MyError).IsError(MyError)
	verify.That(t, MyErrorCpy).IsError(MyError)
}

func TestSentinelNew(t *testing.T) {
	const MyError errors.Sentinel = "MyError"
	var err = MyError.New("error message")

	verify.That(t, err).IsError(MyError)
	verify.That(t, err).ToString().Contains("error message")
}

func TestSentinelWrap(t *testing.T) {
	const MyError errors.Sentinel = "MyError"
	var nestedError = errors.New("nested error")
	var err = MyError.Wrap(nestedError)

	verify.That(t, err).IsError(MyError)
	verify.That(t, err).IsError(nestedError)
}

func TestSentinelErrorf(t *testing.T) {
	verify.That(t, true).IsTrue()

	const MyError errors.Sentinel = "MyError"
	var nestedError = errors.New("nested error")
	var err = MyError.Errorf("error message: %w", nestedError)

	verify.That(t, err).IsError(MyError)
	verify.That(t, err).ToString().Contains("error message")
	verify.That(t, err).IsError(nestedError)
}
