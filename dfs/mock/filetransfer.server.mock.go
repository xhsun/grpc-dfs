package mock

import (
	"context"

	"github.com/stretchr/testify/mock"
	pb "github.com/xhsun/grpc-file-transfer/dfs"
	"google.golang.org/grpc/metadata"
)

type FileTransferServerMock struct {
	mock.Mock
}

func (x *FileTransferServerMock) Send(m *pb.FileContent) error {
	args := x.Called(m)
	return args.Error(0)
}

func (x *FileTransferServerMock) SendAndClose(m *pb.Empty) error {
	args := x.Called(m)
	return args.Error(0)
}

func (x *FileTransferServerMock) Recv() (*pb.File, error) {
	args := x.Called()
	arg := args.Get(0)
	if arg != nil {
		return args.Get(0).(*pb.File), args.Error(1)
	}
	return nil, args.Error(1)
}
func (x *FileTransferServerMock) SetHeader(m metadata.MD) error {
	args := x.Called(m)
	return args.Error(0)
}
func (x *FileTransferServerMock) SendHeader(m metadata.MD) error {
	args := x.Called(m)
	return args.Error(0)
}
func (x *FileTransferServerMock) SetTrailer(m metadata.MD) {
	_ = x.Called(m)
}
func (x *FileTransferServerMock) Context() context.Context {
	args := x.Called()
	return args.Get(0).(context.Context)
}
func (x *FileTransferServerMock) SendMsg(m interface{}) error {
	args := x.Called(m)
	return args.Error(0)
}
func (x *FileTransferServerMock) RecvMsg(m interface{}) error {
	args := x.Called(m)
	return args.Error(0)
}
