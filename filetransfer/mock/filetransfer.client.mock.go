package mock

import (
	"context"

	"github.com/stretchr/testify/mock"
	pb "github.com/xhsun/grpc-file-transfer/filetransfer"
	"google.golang.org/grpc/metadata"
)

type FileTransferClientMock struct {
	mock.Mock
}

func (x *FileTransferClientMock) Send(m *pb.File) error {
	args := x.Called(m)
	return args.Error(0)
}

func (x *FileTransferClientMock) Recv() (*pb.FileContent, error) {
	args := x.Called()
	return args.Get(0).(*pb.FileContent), args.Error(1)
}

func (x *FileTransferClientMock) CloseAndRecv() (*pb.Empty, error) {
	args := x.Called()
	return args.Get(0).(*pb.Empty), args.Error(1)
}

func (x *FileTransferClientMock) CloseSend() error {
	args := x.Called()
	return args.Error(0)
}

func (x *FileTransferClientMock) Header() (metadata.MD, error) {
	args := x.Called()
	return args.Get(0).(metadata.MD), args.Error(1)
}
func (x *FileTransferClientMock) Trailer() metadata.MD {
	args := x.Called()
	return args.Get(0).(metadata.MD)
}
func (x *FileTransferClientMock) Context() context.Context {
	args := x.Called()
	return args.Get(0).(context.Context)
}
func (x *FileTransferClientMock) SendMsg(m interface{}) error {
	args := x.Called()
	return args.Error(0)
}
func (x *FileTransferClientMock) RecvMsg(m interface{}) error {
	args := x.Called()
	return args.Error(0)
}
