package base

import (
	"github.com/pterm/pterm"
	internalError "github.com/xhsun/grpc-file-transfer/client/internal/error"
)

// BaseCommand contains utility methods for CLI commands
type BaseCommand struct {
}

// NewBaseCommand method creates a new base command
func NewBaseCommand() *BaseCommand {
	return &BaseCommand{}
}

// LogError method displays an appropriate message based on the provided error
func (bc *BaseCommand) LogError(err error) {
	var internalErr internalError.IError
	switch err := err.(type) {
	case internalError.IError:
		internalErr = err
	default:
		internalErr = internalError.NewUnexpectedError(err)
		pterm.Debug.Println(internalErr.Error())
	}
	pterm.Error.WithShowLineNumber(false).Printf(internalErr.GetDisplayMessage())
}
