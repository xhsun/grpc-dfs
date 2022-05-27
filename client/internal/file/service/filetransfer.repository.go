package service

import (
	"context"

	"github.com/xhsun/grpc-file-transfer/client/internal/client"
	pb "github.com/xhsun/grpc-file-transfer/filetransfer"
)

type FileTransferRepository struct {
	client pb.FileTransferClient
}

// NewFileTransferRepository method creates a new file transfer repository
func NewFileTransferRepository(client client.GRPCClient) *FileTransferRepository {
	return &FileTransferRepository{
		client: client.Client,
	}
}

//UploadStream - Get upload stream to store file to server
func (ftr *FileTransferRepository) UploadStream(ctx context.Context) (pb.FileTransfer_StoreClient, error) {
	return ftr.client.Store(ctx)
}

//Upload - Upload file chunk to server
func (ftr *FileTransferRepository) Upload(stream pb.FileTransfer_StoreClient, fileName string, fileChunk []byte) error {
	return stream.Send(&pb.File{
		Name:    fileName,
		Content: fileChunk,
	})
}

//UploadClose - Notify server upload completed
func (ftr *FileTransferRepository) UploadClose(stream pb.FileTransfer_StoreClient) error {
	_, err := stream.CloseAndRecv()
	return err
}

//DownloadStream - Get download stream to fetch file from server
func (ftr *FileTransferRepository) DownloadStream(ctx context.Context, fileName string) (pb.FileTransfer_FetchClient, error) {
	return ftr.client.Fetch(ctx, &pb.FileName{Name: fileName})
}

//Download - Download file chunk from server
func (ftr *FileTransferRepository) Download(stream pb.FileTransfer_FetchClient) ([]byte, error) {
	content, err := stream.Recv()
	if err != nil {
		return nil, err
	}
	if content != nil {
		return content.Data, nil
	}
	return []byte{}, nil
}

//ServerFileList - Get list of files from server
func (ftr *FileTransferRepository) ServerFileList(ctx context.Context) (map[string]uint64, error) {
	files, err := ftr.client.ListAll(ctx, &pb.Empty{})
	if err != nil {
		return nil, err
	}
	if files != nil {
		return files.Files, nil
	}
	return make(map[string]uint64), nil
}

//Delete - Delete file from server
func (ftr *FileTransferRepository) Delete(ctx context.Context, fileName string) error {
	_, err := ftr.client.Delete(ctx, &pb.FileName{Name: fileName})
	return err
}
