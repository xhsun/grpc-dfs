package error

import "fmt"

// UnexpectedError representing an unexpected error
type UnexpectedError struct {
	baseError error
}

// NewParseHttpResponseError - Create new UnexpectedError
func NewUnexpectedError(baseError error) *UnexpectedError {
	return &UnexpectedError{baseError: baseError}
}

// Error - detailed error message
func (e *UnexpectedError) Error() string {
	err := "Encountered an unexpected error"
	if e.baseError == nil {
		return err
	} else {
		return fmt.Sprintf("%s: %s", err, e.baseError.Error())
	}
}

// GetDisplayMessage - error message
func (e *UnexpectedError) GetDisplayMessage() string {
	return e.Error()
}
