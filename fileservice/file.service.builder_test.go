package fileservice_test

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/xhsun/grpc-file-transfer/fileservice"
	"github.com/xhsun/grpc-file-transfer/fileservice/mock"
	"syreclabs.com/go/faker"
)

type FileServiceBuilderSuite struct {
	suite.Suite
	target             *fileservice.FileServiceBuilder
	fileRepositoryMock *mock.FileRepositoryMock
}

// The SetupTest method will be run before every test in the suite
func (suite *FileServiceBuilderSuite) SetupTest() {
	suite.fileRepositoryMock = new(mock.FileRepositoryMock)
	suite.target = fileservice.NewFileServiceBuilder(suite.fileRepositoryMock)
}

func (suite *FileServiceBuilderSuite) TestWithFileName() {
	input := faker.Lorem().Word()

	actual, err := suite.target.WithFileName(input).Build()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), input, actual.FileName())
}

func (suite *FileServiceBuilderSuite) TestWithFileNameEmptyName() {
	actual, err := suite.target.WithFileName("").Build()
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), actual)
}

func (suite *FileServiceBuilderSuite) TestWithFile() {
	input := faker.Lorem().Word()
	inputFlag := os.O_RDONLY
	inputFullPath := faker.Lorem().Word()

	suite.fileRepositoryMock.On("FullFilePath", input).Return(inputFullPath, nil)
	suite.fileRepositoryMock.On("Open", inputFullPath, inputFlag).Return(new(os.File), nil)

	actual, err := suite.target.WithFile(input, inputFlag).Build()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), input, actual.FileName())
}

func (suite *FileServiceBuilderSuite) TestWithFileEmptyName() {
	actual, err := suite.target.WithFile("", faker.RandomInt(1, 5)).Build()
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), actual)
}

func (suite *FileServiceBuilderSuite) TestWithFileFullPathGenFail() {
	input := faker.Lorem().Word()
	inputFlag := os.O_RDONLY

	suite.fileRepositoryMock.On("FullFilePath", input).Return("", errors.New(""))

	actual, err := suite.target.WithFile(input, inputFlag).Build()
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), actual)
}

func (suite *FileServiceBuilderSuite) TestWithFileFileError() {
	input := faker.Lorem().Word()
	inputFlag := os.O_RDONLY
	inputFullPath := faker.Lorem().Word()

	suite.fileRepositoryMock.On("FullFilePath", input).Return(inputFullPath, nil)
	suite.fileRepositoryMock.On("Open", inputFullPath, inputFlag).Return(nil, errors.New(""))

	actual, err := suite.target.WithFile(input, inputFlag).Build()
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), actual)
}

func TestFileServiceBuilderSuite(t *testing.T) {
	suite.Run(t, new(FileServiceBuilderSuite))
}
