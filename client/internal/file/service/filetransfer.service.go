package service

import (
	"context"
	"io"
	"os"
	"time"

	"github.com/pterm/pterm"
	"github.com/xhsun/grpc-file-transfer/client/internal/config"
	internalError "github.com/xhsun/grpc-file-transfer/client/internal/error"
	"github.com/xhsun/grpc-file-transfer/fileservice"
)

type FileTransferServer struct {
	fileChunkSize          int
	fileServiceBuilder     fileservice.IFileServiceBuilder
	fileTransferRepository IFileTransferRepository
}

// NewFileTransferServer method creates a new file transfer service
func NewFileTransferServer(config *config.Config, fileServiceBuilder fileservice.IFileServiceBuilder, fileTransferRepository IFileTransferRepository) *FileTransferServer {
	return &FileTransferServer{
		fileChunkSize:          config.FileChunkSize,
		fileServiceBuilder:     fileServiceBuilder,
		fileTransferRepository: fileTransferRepository,
	}
}

//Upload - Upload the content of given file to server
func (fts *FileTransferServer) Upload(fileName string) error {
	if fileName == "" {
		return internalError.NewFileNameError()
	}
	fileService, err := fts.fileServiceBuilder.WithFile(fileName, os.O_RDONLY).Build()
	if err != nil {
		return internalError.NewIOError(err)
	}
	defer fileService.Close()

	pterm.Debug.Printfln("Start to upload %s to server", fileName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := fts.fileTransferRepository.UploadStream(ctx)
	if err != nil {
		return err
	}
	for {
		buffer, err := fileService.Read()
		if err != nil {
			if err == io.EOF {
				pterm.Debug.Printfln("Upload %s complete", fileName)
				err := fts.fileTransferRepository.UploadClose(stream)
				if err != nil {
					return err
				}
				return nil
			}
			return internalError.NewIOError(err)
		}
		err = fts.fileTransferRepository.Upload(stream, fileName, buffer)
		if err != nil {
			return err
		}
	}
}

//Download - Download the content of given service file from server to the given local location
func (fts *FileTransferServer) Download(serverFileName string, localFileName string) error {
	if serverFileName == "" || localFileName == "" {
		return internalError.NewFileNameError()
	}
	fileService, err := fts.fileServiceBuilder.WithFile(localFileName, os.O_CREATE|os.O_WRONLY).Build()
	if err != nil {
		return internalError.NewIOError(err)
	}
	defer fileService.Close()

	pterm.Debug.Printfln("Start to download %s from server to %s", serverFileName, localFileName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := fts.fileTransferRepository.DownloadStream(ctx, serverFileName)
	if err != nil {
		return err
	}
	for {
		segment, err := fts.fileTransferRepository.Download(stream)
		if err != nil {
			if err == io.EOF {
				fileService.Sync()
				pterm.Debug.Printfln("Download %s to %s complete", serverFileName, localFileName)
				return nil
			}
			return internalError.NewIOError(err)
		}

		saved, err := fileService.Write(segment)
		if err != nil {
			return internalError.NewIOError(err)
		}
		pterm.Debug.Printfln("Downloaded %d bytes", saved)
	}
}

//ServerFileList - Get list of files from server
func (fts *FileTransferServer) ServerFileList() (map[string]uint64, error) {
	pterm.Debug.Println("Start to get list of files from server")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	files, err := fts.fileTransferRepository.ServerFileList(ctx)
	if err != nil {
		return nil, err
	}

	pterm.Debug.Printfln("Successfully retrieved list of files: %v", files)
	return files, nil
}

//DeleteFromServer - Delete given file from server
func (fts *FileTransferServer) DeleteFromServer(fileName string) error {
	if fileName == "" {
		return internalError.NewFileNameError()
	}

	pterm.Debug.Printfln("Start to delete %s to server", fileName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := fts.fileTransferRepository.Delete(ctx, fileName)
	if err != nil {
		return err
	}
	pterm.Debug.Printfln("Delete %s complete", fileName)
	return nil
}
