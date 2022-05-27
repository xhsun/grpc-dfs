package error

// FileNameError representing file name not provided
type FileNameError struct {
}

// NewFileNameError - Create new FileNameError
func NewFileNameError() *FileNameError {
	return &FileNameError{}
}

// Error - detailed error message
func (e *FileNameError) Error() string {
	return "File name cannot be empty"
}

// GetDisplayMessage - error message
func (e *FileNameError) GetDisplayMessage() string {
	return e.Error()
}
