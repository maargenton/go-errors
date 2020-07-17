package errors_test

import (
	"testing"

	"github.com/maargenton/go-errors"

	"github.com/maargenton/go-testpredicate/pkg/asserter"
	"github.com/maargenton/go-testpredicate/pkg/p"
)

func TestSentinel(t *testing.T) {
	assert := asserter.New(t)

	const MyError errors.Sentinel = "MyError"

	assert.That(MyError.Error(), p.Eq("MyError"))
}

func TestSentinelValueIdentity(t *testing.T) {
	assert := asserter.New(t)

	const MyError errors.Sentinel = "MyError"
	const MyErrorCpy errors.Sentinel = "MyError"

	assert.That(MyError, p.IsError(MyError))
	assert.That(MyErrorCpy, p.IsError(MyError))
}
