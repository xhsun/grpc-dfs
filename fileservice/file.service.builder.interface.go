package fileservice

//go:generate mockery --srcpkg=github.com/xhsun/grpc-file-transfer/fileservice --name=IFileServiceBuilder --output=./mock --outpkg=mock --structname=FileServiceBuilderMock --filename=file.service.builder.mock.go
type IFileServiceBuilder interface {
	//WithFile - Open a file for read/write
	WithFile(filename string, flag int) IFileServiceBuilder
	//WithFileName - Add file name
	WithFileName(filename string) IFileServiceBuilder
	// Build a new file service
	Build() (IFileService, error)
}
