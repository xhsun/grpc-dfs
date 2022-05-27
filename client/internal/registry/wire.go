//go:build wireinject
// +build wireinject

package registry

import (
	"github.com/google/wire"
	"github.com/xhsun/grpc-file-transfer/client/internal/cli"
	"github.com/xhsun/grpc-file-transfer/client/internal/client"
	"github.com/xhsun/grpc-file-transfer/client/internal/config"
	"github.com/xhsun/grpc-file-transfer/client/internal/file"
	"github.com/xhsun/grpc-file-transfer/client/internal/file/service"
	"github.com/xhsun/grpc-file-transfer/fileservice"
)

var ServiceSet = wire.NewSet(
	service.NewFileRepository,
	wire.Bind(new(fileservice.IFileRepository), new(*service.FileRepository)),
	fileservice.NewFileServiceBuilder,
	wire.Bind(new(fileservice.IFileServiceBuilder), new(*fileservice.FileServiceBuilder)),
	service.NewFileTransferRepository,
	wire.Bind(new(service.IFileTransferRepository), new(*service.FileTransferRepository)),
	service.NewFileTransferServer,
	wire.Bind(new(service.IFileTransferService), new(*service.FileTransferServer)),
)

func InitializeCLI(config *config.Config) (*cli.CLI, error) {
	wire.Build(client.NewGRPCClient, ServiceSet, file.NewFileCommand, cli.NewCLI)
	return &cli.CLI{}, nil
}
