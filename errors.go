package rqp

import "fmt"

// Error special rqp.Error type
type Error struct {
	s string
}

func (e *Error) Error() string {
	return e.s
}

// NewError constructor for internal errors
func NewError(msg string) *Error {
	return &Error{msg}
}

// wrap prefixes err with msg while keeping it matchable via errors.Is.
func wrap(err error, msg string) error {
	return fmt.Errorf("%s: %w", msg, err)
}

// Errors list:
var (
	ErrRequired           = NewError("required")
	ErrBadFormat          = NewError("bad format")
	ErrEmptyValue         = NewError("empty value")
	ErrUnknownMethod      = NewError("unknown method")
	ErrNotInScope         = NewError("not in scope")
	ErrSimilarNames       = NewError("similar names of keys are not allowed")
	ErrMethodNotAllowed   = NewError("method are not allowed")
	ErrFilterNotAllowed   = NewError("filter are not allowed")
	ErrFilterNotFound     = NewError("filter not found")
	ErrValidationNotFound = NewError("validation not found")
)
