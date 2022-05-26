package file

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	log "github.com/sirupsen/logrus"
	pb "github.com/xhsun/grpc-file-transfer/filetransfer"
	"github.com/xhsun/grpc-file-transfer/server/internal/config"
)

type FileHandler struct {
	pb.UnimplementedFileTransferServer
	fileStoragePath string
	fileChunkSize   int
}

// FileHandler method creates a new file handler
func NewFileHandler(config *config.Config) *FileHandler {
	return &FileHandler{
		fileStoragePath: config.FileStoragePath,
		fileChunkSize:   config.FileChunkSize,
	}
}

// Store files on the server
func (fh *FileHandler) Store(stream pb.FileTransfer_StoreServer) error {
	var file *os.File
	for {
		segment, err := stream.Recv()
		logger := log.WithField("FileName", segment.Name)
		if err == io.EOF {
			if file != nil {
				file.Sync()
			}
			logger.Info("Successfully stored given file")
			return stream.SendAndClose(&pb.Empty{})
		}
		if err != nil {
			log.WithError(err).Error("Encountered unexpected error while reading file segments")
			return err
		}
		if file == nil {
			if segment.Name == "" {
				log.Error("File name cannot be empty")
				return err
			}
			file, err = os.OpenFile(fmt.Sprintf(fh.fileStoragePath, segment.Name), os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				logger.WithError(err).Error("Encountered unexpected error while attempt to open or create file")
				return err
			}
			defer file.Close()
		}

		saved, err := file.Write(segment.Content)
		logger.Debugf("Wrote %d bytes", saved)
		if err != nil {
			logger.WithError(err).Error("Encountered unexpected error while attempt to write to file")
			return err
		}
	}
}

func (fh *FileHandler) Fetch(fileName *pb.FileName, stream pb.FileTransfer_FetchServer) error {
	if fileName == nil || fileName.Name == "" {
		log.Error("File name cannot be empty")
		return errors.New("File name cannot be empty")
	}
	logger := log.WithField("FileName", fileName.Name)

	file, err := os.Open(fmt.Sprintf(fh.fileStoragePath, fileName.Name))
	if err != nil {
		logger.WithError(err).Error("Encountered unexpected error while attempt to open file")
	}
	defer file.Close()

	buffer := make([]byte, fh.fileChunkSize)
	for {
		n, err := file.Read(buffer[:cap(buffer)])
		if err != nil {
			if err == io.EOF {
				logger.Info("Successfully send file content to client")
				return nil
			}
			logger.WithError(err).Error("Encountered unexpected error while attempt to read file")
			return err
		}
		buffer = buffer[:n]
		err = stream.Send(&pb.FileContent{Data: buffer})
		if err != nil {
			logger.WithError(err).Error("Encountered unexpected error while attempt to send file chunk")
			return err
		}
	}
}

func (fh *FileHandler) Delete(ctx context.Context, fileName *pb.FileName) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}

func (fh *FileHandler) ListAll(ctx context.Context, empty *pb.Empty) (*pb.FileList, error) {
	return &pb.FileList{}, nil
}
