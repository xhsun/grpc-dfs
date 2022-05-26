// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package registry

import (
	"github.com/google/wire"
	"github.com/xhsun/grpc-file-transfer/fileservice"
	"github.com/xhsun/grpc-file-transfer/server/internal/config"
	"github.com/xhsun/grpc-file-transfer/server/internal/file"
	"github.com/xhsun/grpc-file-transfer/server/internal/file/service"
	"github.com/xhsun/grpc-file-transfer/server/internal/server"
)

// Injectors from wire.go:

func InitializeServer(config2 *config.Config) (*server.GRPCServer, error) {
	fileRepository := service.NewFileRepository(config2)
	fileServiceBuilder := fileservice.NewFileServiceBuilder(fileRepository)
	fileHandler := file.NewFileHandler(config2, fileServiceBuilder)
	grpcServer := server.NewGRPCServer(config2, fileHandler)
	return grpcServer, nil
}

// wire.go:

var ServiceBuilderSet = wire.NewSet(service.NewFileRepository, wire.Bind(new(fileservice.IFileRepository), new(*service.FileRepository)), fileservice.NewFileServiceBuilder, wire.Bind(new(fileservice.IFileServiceBuilder), new(*fileservice.FileServiceBuilder)))
