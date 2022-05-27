package service

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/xhsun/grpc-file-transfer/client/internal/config"
)

type FileRepository struct {
	fileChunkSize int
}

// NewFileRepository method creates a new file repository
func NewFileRepository(config *config.Config) *FileRepository {
	return &FileRepository{
		fileChunkSize: config.FileChunkSize,
	}
}

//FullFilePath - Generate full relative file path for the given file
func (fr *FileRepository) FullFilePath(filename string) (string, error) {
	return filename, nil
}

//Open - Open file at given path with given flags
func (fr *FileRepository) Open(path string, flag int) (*os.File, error) {
	if path == "" {
		return nil, errors.New("path cannot be empty")
	}
	file, err := os.OpenFile(path, flag, 0644)
	if err != nil {
		return nil, err
	}
	return file, nil
}

//List - List all file iteratively in the file storage folder
func (fr *FileRepository) List(path string) (map[string]uint64, error) {
	files := make(map[string]uint64)
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files[path] = uint64(info.Size())
		}
		return nil
	})
	return files, err
}

//Read - Read chunk from file
func (fr *FileRepository) Read(file *os.File) ([]byte, error) {
	buffer := make([]byte, fr.fileChunkSize)
	n, err := file.Read(buffer[:cap(buffer)])
	if err != nil {
		return buffer, err
	}
	buffer = buffer[:n]
	return buffer, nil
}

//Write - Write chunk to file
func (fr *FileRepository) Write(file *os.File, data []byte) (int, error) {
	return file.Write(data)
}

//Remove - Remove file at given path
func (fr *FileRepository) Remove(path string) error {
	return os.Remove(path)
}

//Sync - Sync current writes
func (fr *FileRepository) Sync(file *os.File) {
	if file != nil {
		file.Sync()
	}
}

//Close - Close file
func (fr *FileRepository) Close(file *os.File) {
	if file != nil {
		file.Close()
	}
}