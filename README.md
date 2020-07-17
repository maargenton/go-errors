# go-errors

Drop-in replacement of native `errors` package that adds value type sentinel errors and error composition.

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
As ogf Go 1.13, this include:

- `errors.New()`
- `errors.Unwrap()`
- `errors.Is()`
- `errors.As()`


The package defines a `Sentinel` error type based on string that can be
defined as const. They are immutable and proper value type.

```go
const ErrMyError errors.Sentinel = "ErrMyError: something happened"
```

It also defines a way to compose multiple errors together, typically a
sentinel error and an detailed error message potentially with underlying
error. Unwrap() returns the second error, but Is() and As() follow both path.

```go
err = errors.Compose(ErrMyError, underlyingError)

errors.Is(err, ErrMyError)      // => true
errors.Is(err, underlyingError) // => true
```
