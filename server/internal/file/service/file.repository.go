package service

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/xhsun/grpc-file-transfer/server/internal/config"
)

type FileRepository struct {
	fileStoragePath string
	fileChunkSize   int
}

// NewFileRepository method creates a new file repository
func NewFileRepository(config *config.Config) *FileRepository {
	return &FileRepository{
		fileStoragePath: config.FileStoragePath,
		fileChunkSize:   config.FileChunkSize,
	}
}

//FullFilePath - Generate full relative file path for the given file
func (fr *FileRepository) FullFilePath(filename string) (string, error) {
	cleaned := filepath.Join("/", filename)
	path := filepath.Join(fr.fileStoragePath, cleaned)
	dirName := filepath.Dir(path)
	if _, err := os.Stat(dirName); err != nil {
		err := os.MkdirAll(dirName, os.ModePerm)
		if err != nil {
			log.WithField("Path", path).Error("Unable to create directories")
			return "", err
		}
	}
	return path, nil
}

//FileSize - Return file size of the given file name
func (fr *FileRepository) FileSize(path string) (uint64, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	return uint64(fi.Size()), nil
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
		if err != nil {
			log.WithError(err).Debug("Encountered unexpected error while attempt to retrieving file info")
			return nil
		}
		if !info.IsDir() {
			parts := strings.SplitN(path, "/", 2)
			files[filepath.Join(parts[1:]...)] = uint64(info.Size())
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
