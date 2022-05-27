package fileservice

import "os"

//go:generate mockery --srcpkg=github.com/xhsun/grpc-file-transfer/fileservice --name=IFileRepository --output=./mock --outpkg=mock --structname=FileRepositoryMock --filename=file.repository.mock.go
type IFileRepository interface {
	//FullFilePath - Generate full relative file path for the given file
	FullFilePath(filename string) (string, error)
	//Open - Open file at given path with given flags
	Open(path string, flag int) (*os.File, error)
	//List - List all file iteratively in the file storage folder
	List(path string) (map[string]uint64, error)
	//Read - Read chunk from file
	Read(file *os.File) ([]byte, error)
	//Write - Write chunk to file
	Write(file *os.File, data []byte) (int, error)
	//Remove - Remove file at given path
	Remove(path string) error
	//Sync - Sync current writes
	Sync(file *os.File)
	//Close - Close file
	Close(file *os.File)
}
