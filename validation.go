package rqp

import (
	"fmt"
	"slices"
)

// ValidationFunc represents validator for Filters
type ValidationFunc func(value any) error

// Validations type replacement for map.
// Used in NewParse(), NewQV(), SetValidations()
type Validations map[string]ValidationFunc

// Multi multiple validation func
// usage: Multi(Min(10), Max(100))
func Multi(values ...ValidationFunc) ValidationFunc {
	return func(value any) error {
		for _, v := range values {
			if err := v(value); err != nil {
				return err
			}
		}
		return nil
	}
}

// In validation if values contatin value
func In(values ...any) ValidationFunc {
	return func(value any) error {
		if !slices.Contains(values, value) {
			return notInScope(value)
		}
		return nil
	}
}

// Min validation if value greater or equal then min
func Min(min int) ValidationFunc {
	return func(value any) error {
		if limit, ok := value.(int); ok {
			if limit >= min {
				return nil
			}
		}
		return notInScope(value)
	}
}

// Max validation if value lower or equal then max
func Max(max int) ValidationFunc {
	return func(value any) error {
		if limit, ok := value.(int); ok {
			if limit <= max {
				return nil
			}
		}
		return notInScope(value)
	}
}

// MinMax validation if value between or equal min and max
func MinMax(min, max int) ValidationFunc {
	return func(value any) error {
		if limit, ok := value.(int); ok {
			if min <= limit && limit <= max {
				return nil
			}
		}
		return notInScope(value)
	}
}

// NotEmpty validation if string value length more then 0
func NotEmpty() ValidationFunc {
	return func(value any) error {
		if s, ok := value.(string); ok {
			if len(s) > 0 {
				return nil
			}
		}
		return notInScope(value)
	}
}

// notInScope wraps ErrNotInScope with the rejected value.
func notInScope(value any) error {
	return fmt.Errorf("%v: %w", value, ErrNotInScope)
}
