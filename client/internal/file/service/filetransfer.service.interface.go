package service

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
