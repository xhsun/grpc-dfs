package error

import "fmt"

// IOError representing IO error
type IOError struct {
	baseError error
}

// NewIOError - Create new IOError
func NewIOError(baseError error) *IOError {
	return &IOError{baseError: baseError}
}

// Error - detailed error message
func (e *IOError) Error() string {
	return e.baseError.Error()
}

// GetDisplayMessage - error message
func (e *IOError) GetDisplayMessage() string {
	return fmt.Sprintf("Encountered unexpected issue with the file: %s", e.baseError.Error())
}
