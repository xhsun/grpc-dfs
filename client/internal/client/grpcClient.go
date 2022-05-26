package client

import (
	log "github.com/sirupsen/logrus"

	"github.com/xhsun/grpc-file-transfer/client/internal/config"
	pb "github.com/xhsun/grpc-file-transfer/filetransfer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type GRPCClient struct {
	Client pb.FileTransferClient
}

// NewGRPCClient method creates a new gRPC client
func NewGRPCClient(config *config.Config) *GRPCClient {
	creds, err := credentials.NewClientTLSFromFile(config.CertFilePath, config.ServerHostOverride)
	if err != nil {
		log.WithError(err).Fatal("Failed to create TLS credentials")
	}

	conn, err := grpc.Dial(config.ServerAddress, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.WithError(err).Fatal("fail to dial")
	}
	client := pb.NewFileTransferClient(conn)
	return &GRPCClient{client}
}
