# go-errors

Drop-in replacement of the native Go `errors` package that adds value type
sentinel errors and error composition.

[![GoDoc](https://godoc.org/github.com/maargenton/go-errors?status.svg)](https://godoc.org/github.com/maargenton/go-errors)
[![Version](https://img.shields.io/github/tag/maargenton/go-errors.svg)](https://github.com/maargenton/go-errors)
[![Build Status](https://travis-ci.org/maargenton/go-errors.svg?branch=master)](https://travis-ci.org/maargenton/go-errors)
[![codecov](https://codecov.io/gh/maargenton/go-errors/branch/master/graph/badge.svg)](https://codecov.io/gh/maargenton/go-errors)
[![Go Report Card](https://goreportcard.com/badge/github.com/maargenton/go-errors)](https://goreportcard.com/report/github.com/maargenton/go-errors)


## Installation

```shell
go get -u github.com/maargenton/go-errors
```

```go
import "github.com/maargenton/go-errors"
```

## Usage

Package errors is a drop-in replacement for the built-in errors package. It
forwards all call to the errors API directly to the built-in errors package.
As of Go 1.13, this include:

- `errors.New()`
- `errors.Unwrap()`
- `errors.Is()`
- `errors.As()`

### Sentinel error

Sentinel errors can be defined as const values of type errors.Sentinel. Such sentinel errors are immutable and proper value type. They can be used directly in foreign packages to test against, but they can also be redefined locally to reduce package dependency and still be tested against.

```go
const ErrMyError errors.Sentinel = "ErrMyError"

if errors.Is(err, ErrMyError) {
    // Specific error ErrMyError can be handled here
}

if errors.Is(err, errors.Sentinel("ErrMyError") {
    // This also detects the same error
}

```

### Error composition

With sentinel errors, instead of wrapping one error into another, you might want
to compose two or more errors together with `errors.Compose()`. Typical use-case
would compose a sentinel error with either an underlying error or a more
detailed error message.

The error returned by `errors.Compose()` unwraps into the second error argument, but when using `errors.Is()` or `errors.As()`, all error branches are evaluated from left to right.

```go
err = errors.Compose(ErrMyError, underlyingError)

errors.Is(err, ErrMyError)      // => true
errors.Is(err, underlyingError) // => true
```
