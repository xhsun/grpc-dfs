package service

//go:generate mockery --srcpkg=github.com/xhsun/grpc-file-transfer/client/internal/file/service --name=IFileTransferService --output=../../mock --outpkg=mock --structname=FileTransferServiceMock --filename=filetransfer.service.mock.go
type IFileTransferService interface {
	//Upload - Upload the content of given file to server
	Upload(fileName string) error
	//Download - Download the content of given service file from server to the given local location
	Download(serverFileName string, localFileName string) error
	//ServerFileList - Get list of files from server
	ServerFileList() (map[string]uint64, error)
	//DeleteFromServer - Delete given file from server
	DeleteFromServer(fileName string) error
}
