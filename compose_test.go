package errors_test

import (
	"testing"

	"github.com/maargenton/go-errors"

	"github.com/maargenton/go-testpredicate/pkg/asserter"
	"github.com/maargenton/go-testpredicate/pkg/p"
)

var err1 = errors.New("err1")
var err2 = errors.New("err2")
var err3 = errors.New("err3")

func TestCompose(t *testing.T) {
	assert := asserter.New(t)
	var err = errors.Compose(err1, err2)

	assert.That(err.Error(), p.Eq("err1: err2"))
	assert.That(errors.Unwrap(err), p.Eq(err2))
	assert.That(err, p.IsError(err1))
	assert.That(err, p.IsError(err2))
}

func TestComposeMultiple(t *testing.T) {
	assert := asserter.New(t)
	var err = errors.Compose(err1, err2, err3)

	assert.That(err.Error(), p.Eq("err1: err2: err3"))
	assert.That(err, p.IsError(err1))
	assert.That(err, p.IsError(err2))
	assert.That(err, p.IsError(err3))
}

type testError struct{}
type testError2 struct{}

func (*testError) Error() string  { return "testError" }
func (*testError2) Error() string { return "testError2" }

func TestComposeAs(t *testing.T) {
	assert := asserter.New(t)

	var err = errors.Compose(&testError{}, &testError2{})
	var err1 *testError
	var err2 *testError2

	assert.That(errors.As(err, &err1), p.IsTrue())
	assert.That(errors.As(err, &err2), p.IsTrue())
}
