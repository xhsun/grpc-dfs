package service

import (
	"errors"
	"os"

	log "github.com/sirupsen/logrus"
)

type FileServiceBuilder struct {
	fileRepository IFileRepository
	err            error
	file           *os.File
	fileName       string
}

// NewFileServiceBuilder to construct the file service builder
func NewFileServiceBuilder(fileRepository IFileRepository) *FileServiceBuilder {
	return &FileServiceBuilder{
		fileRepository: fileRepository,
	}
}

//WithFileName - Add file name
func (fsb *FileServiceBuilder) WithFileName(filename string) IFileServiceBuilder {
	if filename == "" {
		err := errors.New("file name cannot be empty")
		log.WithError(err).Error("File name cannot be empty")
		return &FileServiceBuilder{
			fileRepository: fsb.fileRepository,
			err:            err,
		}
	}
	return &FileServiceBuilder{
		fileRepository: fsb.fileRepository,
		fileName:       filename,
		file:           fsb.file,
	}
}

//WithFile - Open a file for read/write
func (fsb *FileServiceBuilder) WithFile(filename string, flag int) IFileServiceBuilder {
	if filename == "" {
		err := errors.New("file name cannot be empty")
		log.WithError(err).Error("File name cannot be empty")
		return &FileServiceBuilder{
			fileRepository: fsb.fileRepository,
			err:            err,
		}
	}
	path, err := fsb.fileRepository.FullFilePath(filename)
	if err != nil {
		log.WithError(err).WithField("FileName", filename).Error("Encountered an error while generating full path")
		return &FileServiceBuilder{
			fileRepository: fsb.fileRepository,
			err:            err,
		}
	}
	file, err := fsb.fileRepository.Open(path, flag)
	if err != nil {
		log.WithError(err).WithField("Path", path).Error("Encountered an error while opening or creating the file")
		return &FileServiceBuilder{
			fileRepository: fsb.fileRepository,
			err:            err,
		}
	}
	return &FileServiceBuilder{
		fileRepository: fsb.fileRepository,
		fileName:       filename,
		file:           file,
	}
}

// Build a new file service
func (fsb *FileServiceBuilder) Build() (IFileService, error) {
	if fsb.err != nil {
		return nil, fsb.err
	}
	return NewFileService(fsb.fileRepository, fsb.fileName, fsb.file), nil
}
