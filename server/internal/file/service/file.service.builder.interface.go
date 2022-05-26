package service

type IFileServiceBuilder interface {
	//WithFile - Open a file for read/write
	WithFile(filename string, flag int) IFileServiceBuilder
	// Build a new file service
	Build() (IFileService, error)
}
