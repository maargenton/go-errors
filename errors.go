// Package errors is a drop-in replacement for the built-in errors package. It
// forwards all call to the errors API directly to the built-in errors package.
// As ogf Go 1.13, this include:
//
//   	errors.New()
//   	errors.Unwrap()
//   	errors.Is()
//   	errors.As()
//
// The package defines a `Sentinel` error type based on string that can be
// defined as const. They are immutable and proper value type.
//
//      const ErrMyError errors.Sentinel = "ErrMyError: something happened"
//
// It also defines a way to compose multiple errors together, typically a
// sentinel error and an detailed error message potentially with underlying
// error. Unwrap() returns the second error, but Is() and As() follow both path.
//
//      err = errors.Compose(ErrMyError, underlyingError)
//
package errors

import "errors"

// New is a passthrough reference to errors.New
var New = errors.New

// Unwrap is a passthrough reference to errors.Unwrap
var Unwrap = errors.Unwrap

// Is is a passthrough reference to errors.Is
var Is = errors.Is

// As is a passthrough reference to errors.As
var As = errors.As
