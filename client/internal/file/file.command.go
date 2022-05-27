package file

import (
	"fmt"
	"strings"

	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
	"github.com/xhsun/grpc-file-transfer/client/internal/cli/base"
	internalError "github.com/xhsun/grpc-file-transfer/client/internal/error"
	"github.com/xhsun/grpc-file-transfer/client/internal/file/service"
)

type FileCommand struct {
	base.BaseCommand
	fileTransferService service.IFileTransferService
}

// NewFileCommand method creates a new file command
func NewFileCommand(fileTransferService service.IFileTransferService) *FileCommand {
	return &FileCommand{
		fileTransferService: fileTransferService,
	}
}

// Upload - Upload the given file to server
func (fc *FileCommand) Upload(c *cli.Context) error {
	fileName := strings.TrimSpace(c.Args().First())
	if fileName == "" {
		fc.LogError(internalError.NewFileNameError())
	}

	if err := fc.fileTransferService.Upload(fileName); err != nil {
		fc.LogError(err)
		return err
	}

	pterm.Success.WithShowLineNumber(false).Printfln("Successfully uploaded %s to server", fileName)
	return nil
}

// Remove - Remove the given file from server
func (fc *FileCommand) Download(c *cli.Context) error {
	serverFileName := strings.TrimSpace(c.Args().First())
	localFileName := strings.TrimSpace(c.Args().Get(1))
	if serverFileName == "" || localFileName == "" {
		err := internalError.NewFileNameError()
		fc.LogError(err)
		return err
	}

	if err := fc.fileTransferService.Download(serverFileName, localFileName); err != nil {
		fc.LogError(err)
		return err
	}

	pterm.Success.WithShowLineNumber(false).Printfln("Successfully downloaded content of %s to %s", serverFileName, localFileName)
	return nil
}

// ListFiles -  Displays all files in the server
func (fc *FileCommand) ListFiles(c *cli.Context) error {
	files, err := fc.fileTransferService.ServerFileList()
	if err != nil {
		fc.LogError(err)
		return err
	}

	tableData := [][]string{{"File Name", "File Size"}}
	for name, size := range files {
		tableData = append(tableData, []string{name, fmt.Sprint(size)})
	}
	pterm.DefaultTable.WithHasHeader().WithData(tableData).Render()
	return nil
}

// Remove - Remove the given file from server
func (fc *FileCommand) Remove(c *cli.Context) error {
	fileName := strings.TrimSpace(c.Args().First())
	if fileName == "" {
		fc.LogError(internalError.NewFileNameError())
	}

	if err := fc.fileTransferService.DeleteFromServer(fileName); err != nil {
		fc.LogError(err)
		return err
	}

	pterm.Success.WithShowLineNumber(false).Printfln("Successfully removed %s from server", fileName)
	return nil
}
