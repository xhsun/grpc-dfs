package service

type IFileService interface {
	//List - List all files iteratively in <FileStoragePath>
	List() (map[string]uint64, error)
	//Remove - Remove file at <FileStoragePath>/<FileName>
	Remove(fileName string) error
	//Write - Write data to file
	Write(data []byte) (int, error)
	//Read - Read data to file
	Read() ([]byte, error)
	//Sync - Sync currently opened file
	Sync()
	//Close - Close currently opened file
	Close()
}
