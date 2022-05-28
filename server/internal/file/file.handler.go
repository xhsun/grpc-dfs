package file

import (
	"context"
	"errors"
	"io"
	"os"

	log "github.com/sirupsen/logrus"
	fileService "github.com/xhsun/grpc-file-transfer/fileservice"
	pb "github.com/xhsun/grpc-file-transfer/filetransfer"
	"github.com/xhsun/grpc-file-transfer/server/internal/config"
)

type FileHandler struct {
	pb.UnimplementedFileTransferServer
	fileServiceBuilder fileService.IFileServiceBuilder
	fileStoragePath    string
	fileChunkSize      int
}

// FileHandler method creates a new file handler
func NewFileHandler(config *config.Config, fileServiceBuilder fileService.IFileServiceBuilder) *FileHandler {
	return &FileHandler{
		fileStoragePath:    config.FileStoragePath,
		fileChunkSize:      config.FileChunkSize,
		fileServiceBuilder: fileServiceBuilder,
	}
}

// Store files on the server
func (fh *FileHandler) Store(stream pb.FileTransfer_StoreServer) error {
	var fileService fileService.IFileService
	logger := log.WithField("Function", "Store")
	for {
		segment, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				if fileService != nil {
					fileService.Sync()
					logger.WithField("FileName", fileService.FileName()).Info("Successfully stored given file")
				}
				return stream.SendAndClose(&pb.Empty{})
			}
			logger.WithError(err).Error("Encountered unexpected error while reading file segments")
			return err
		}
		if fileService == nil {
			fileService, err = fh.fileServiceBuilder.WithFile(segment.Name, os.O_CREATE|os.O_WRONLY).Build()
			if err != nil {
				logger.WithError(err).Error("Encountered unexpected error while generating file service")
				return err
			}
			logger = logger.WithField("FileName", segment.Name)
			defer fileService.Close()
		}

		saved, err := fileService.Write(segment.Content)
		if err != nil {
			logger.WithError(err).Error("Encountered unexpected error while attempt to write to file")
			return err
		}
		logger.Debugf("Wrote %d bytes", saved)
	}
}

func (fh *FileHandler) Fetch(fileName *pb.FileName, stream pb.FileTransfer_FetchServer) error {
	logger := log.WithField("Function", "Fetch")
	if fileName == nil || fileName.Name == "" {
		logger.Error("File name cannot be empty")
		return errors.New("file name cannot be empty")
	}
	logger = logger.WithField("FileName", fileName.Name)
	fileService, err := fh.fileServiceBuilder.WithFile(fileName.Name, os.O_RDONLY).Build()
	if err != nil {
		logger.WithError(err).Error("Encountered unexpected error while generating file service")
		return err
	}
	defer fileService.Close()

	total, err := fileService.FileSize()
	if err != nil {
		return err
	}

	for {
		buffer, err := fileService.Read()
		if err != nil {
			if err == io.EOF {
				logger.Info("Successfully send file content to client")
				return nil
			}
			logger.WithError(err).Error("Encountered unexpected error while attempt to read file")
			return err
		}
		err = stream.Send(&pb.FileContent{Data: buffer, Total: total})
		if err != nil {
			logger.WithError(err).Error("Encountered unexpected error while attempt to send file chunk")
			return err
		}
	}
}

func (fh *FileHandler) Delete(ctx context.Context, fileName *pb.FileName) (*pb.Empty, error) {
	logger := log.WithField("Function", "Delete")
	if fileName == nil || fileName.Name == "" {
		logger.Error("File name cannot be empty")
		return nil, errors.New("file name cannot be empty")
	}

	fileService, _ := fh.fileServiceBuilder.WithFileName(fileName.Name).Build()
	err := fileService.Remove()
	if err != nil {
		logger.WithError(err).Error("Encountered unexpected error while removing file")
		return nil, err
	}

	logger.WithField("FileName", fileName.Name).Info("Successfully removed file")
	return &pb.Empty{}, nil
}

func (fh *FileHandler) ListAll(ctx context.Context, empty *pb.Empty) (*pb.FileList, error) {
	logger := log.WithField("Function", "ListAll")
	fileService, _ := fh.fileServiceBuilder.Build()
	fileList, err := fileService.List()
	if err != nil {
		logger.WithError(err).Error("Encountered unexpected error while retrieving file list")
		return nil, err
	}

	logger.WithField("data", fileList).Info("Successfully retrieved file list")
	return &pb.FileList{Files: fileList}, nil
}
