//go:build wireinject
// +build wireinject

package registry

import (
	"github.com/google/wire"
	"github.com/xhsun/grpc-file-transfer/server/internal/config"
	"github.com/xhsun/grpc-file-transfer/server/internal/file"
	"github.com/xhsun/grpc-file-transfer/server/internal/file/service"
	"github.com/xhsun/grpc-file-transfer/server/internal/server"
)

var ServiceBuilderSet = wire.NewSet(
	service.NewFileRepository,
	wire.Bind(new(service.IFileRepository), new(*service.FileRepository)),
	service.NewFileServiceBuilder,
	wire.Bind(new(service.IFileServiceBuilder), new(*service.FileServiceBuilder)),
)

func InitializeServer(config *config.Config) (*server.GRPCServer, error) {
	wire.Build(ServiceBuilderSet, file.NewFileHandler, server.NewGRPCServer)
	return &server.GRPCServer{}, nil
}
