package service

import (
	"os"

	log "github.com/sirupsen/logrus"
)

type FileService struct {
	fileName       string
	fileRepository IFileRepository
	file           *os.File
}

// NewFileService to construct the file service
func NewFileService(fileRepository IFileRepository, fileName string, file *os.File) *FileService {
	return &FileService{
		fileRepository: fileRepository,
		fileName:       fileName,
		file:           file,
	}
}

//FileName - return current file name
func (fs *FileService) FileName() string {
	return fs.fileName
}

//List - List all files iteratively in <FileStoragePath>
func (fs *FileService) List() (map[string]uint64, error) {
	path, err := fs.fileRepository.FullFilePath("")
	if err != nil {
		log.WithError(err).Error("Encountered unexpected error while attempt to retrieving file storage path")
		return nil, err
	}

	fileList, err := fs.fileRepository.List(path)
	if err != nil {
		log.WithError(err).Error("Encountered unexpected error while attempt to retrieving file list")
		return nil, err
	}

	return fileList, nil
}

//Remove - Remove file at <FileStoragePath>/<FileName>
func (fs *FileService) Remove() error {
	path, err := fs.fileRepository.FullFilePath(fs.fileName)
	if err != nil {
		return err
	}

	err = fs.fileRepository.Remove(path)
	if err != nil {
		log.WithField("FileName", fs.fileName).WithError(err).Error("Encountered unexpected error while attempt to delete file")
		return err
	}
	return nil
}

//Write - Write data to file
func (fs *FileService) Write(data []byte) (int, error) {
	return fs.fileRepository.Write(fs.file, data)
}

//Read - Read data to file
func (fs *FileService) Read() ([]byte, error) {
	return fs.fileRepository.Read(fs.file)
}

//Sync - Sync currently opened file
func (fs *FileService) Sync() {
	fs.fileRepository.Sync(fs.file)
}

//Close - Close currently opened file
func (fs *FileService) Close() {
	fs.fileRepository.Close(fs.file)
}
