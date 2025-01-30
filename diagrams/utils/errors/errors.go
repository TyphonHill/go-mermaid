package errors

import (
	"fmt"
)

// Common error types
var (
	ErrInvalidInput = fmt.Errorf("invalid input")
	ErrNotFound     = fmt.Errorf("not found")
)

// WrapError wraps an error with additional context
func WrapError(err error, message string) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", message, err)
}

// FileError wraps file operation errors
func FileError(op string, path string, err error) error {
	return fmt.Errorf("failed to %s file %s: %w", op, path, err)
}
