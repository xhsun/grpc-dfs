package service

import (
	"context"

	pb "github.com/xhsun/grpc-file-transfer/filetransfer"
)

type IFileTransferRepository interface {
	//UploadStream - Get upload stream to store file to server
	UploadStream(ctx context.Context) (pb.FileTransfer_StoreClient, error)
	//Upload - Upload file chunk to to server
	Upload(stream pb.FileTransfer_StoreClient, fileName string, fileChunk []byte) error
	//UploadClose - Notify server upload completed
	UploadClose(stream pb.FileTransfer_StoreClient) error
	//DownloadStream - Get download stream to fetch file from server
	DownloadStream(ctx context.Context, fileName string) (pb.FileTransfer_FetchClient, error)
	//Download - Download file chunk from server
	Download(stream pb.FileTransfer_FetchClient) ([]byte, error)
	//ServerFileList - Get list of files from server
	ServerFileList(ctx context.Context) (map[string]uint64, error)
	//Delete - Delete file from server
	Delete(ctx context.Context, fileName string) error
}
