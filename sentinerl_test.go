package errors_test

import (
	"testing"

	"github.com/maargenton/go-testpredicate/pkg/asserter"
	"github.com/maargenton/go-testpredicate/pkg/p"

	"github.com/maargenton/go-errors"
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

func TestSentinelNew(t *testing.T) {
	assert := asserter.New(t)

	const MyError errors.Sentinel = "MyError"
	var err = MyError.New("error message")

	assert.That(err, p.IsError(MyError))
	assert.That(err, p.ToString(p.Contains("error message")))
}

func TestSentinelWrap(t *testing.T) {
	assert := asserter.New(t)

	const MyError errors.Sentinel = "MyError"
	var nestedError = errors.New("nested error")
	var err = MyError.Wrap(nestedError)

	assert.That(err, p.IsError(MyError))
	assert.That(err, p.IsError(nestedError))
}

func TestSentinelErrorf(t *testing.T) {
	assert := asserter.New(t)
	assert.That(true, p.IsTrue())

	const MyError errors.Sentinel = "MyError"
	var nestedError = errors.New("nested error")
	var err = MyError.Errorf("error message: %w", nestedError)

	assert.That(err, p.IsError(MyError))
	assert.That(err, p.ToString(p.Contains("error message")))
	assert.That(err, p.IsError(nestedError))
}
