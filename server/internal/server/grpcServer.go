package server

import (
	"fmt"
	"net"

	log "github.com/sirupsen/logrus"
	pb "github.com/xhsun/grpc-file-transfer/filetransfer"
	"github.com/xhsun/grpc-file-transfer/server/internal/config"
	"github.com/xhsun/grpc-file-transfer/server/internal/file"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// GRPCServer - The gRPC server
type GRPCServer struct {
	config      *config.Config
	fileHandler *file.FileHandler
}

// NewGRPCServer method creates a new gRPC server
func NewGRPCServer(config *config.Config, fileHandler *file.FileHandler) *GRPCServer {
	return &GRPCServer{
		config:      config,
		fileHandler: fileHandler,
	}
}

// Start - Start the gRPC server
func (s *GRPCServer) Start() error {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", s.config.Port))
	if err != nil {
		log.WithError(err).Fatalf("failed to listen to port %d", s.config.Port)
	}

	creds, err := credentials.NewServerTLSFromFile(s.config.CertFilePath, s.config.KeyFilePath)
	if err != nil {
		log.WithError(err).Fatal("Failed to generate credentials")
	}
	grpcServer := grpc.NewServer([]grpc.ServerOption{grpc.Creds(creds)}...)
	pb.RegisterFileTransferServer(grpcServer, s.fileHandler)
	log.Info("Started file transfer server")
	return grpcServer.Serve(lis)
}
