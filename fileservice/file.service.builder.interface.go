package fileservice

type IFileServiceBuilder interface {
	//WithFile - Open a file for read/write
	WithFile(filename string, flag int) IFileServiceBuilder
	//WithFileName - Add file name
	WithFileName(filename string) IFileServiceBuilder
	// Build a new file service
	Build() (IFileService, error)
}
