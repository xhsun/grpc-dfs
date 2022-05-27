package error

// IError - interface for base errors
type IError interface {
	error
	// GetDisplayMessage - User friendly error message
	GetDisplayMessage() string
}
