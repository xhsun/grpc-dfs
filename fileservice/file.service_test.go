package fileservice_test

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	mocks "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/xhsun/grpc-file-transfer/fileservice"
	"github.com/xhsun/grpc-file-transfer/fileservice/mock"
	"syreclabs.com/go/faker"
)

type FileServiceSuite struct {
	suite.Suite
	target             *fileservice.FileService
	fileRepositoryMock *mock.FileRepositoryMock
	fileName           string
}

// The SetupTest method will be run before every test in the suite
func (suite *FileServiceSuite) SetupTest() {
	suite.fileName = faker.RandomString(5)
	suite.fileRepositoryMock = new(mock.FileRepositoryMock)
	suite.target = fileservice.NewFileService(suite.fileRepositoryMock, suite.fileName, new(os.File))
}

func (suite *FileServiceSuite) TestFileName() {
	actual := suite.target.FileName()
	assert.Equal(suite.T(), suite.fileName, actual)
}

func (suite *FileServiceSuite) TestFileSize() {
	inputFullPath := faker.Lorem().Word()
	expected := uint64(faker.RandomInt64(3, 10))

	suite.fileRepositoryMock.On("FullFilePath", suite.fileName).Return(inputFullPath, nil)
	suite.fileRepositoryMock.On("FileSize", inputFullPath).Return(expected, nil)

	actual, err := suite.target.FileSize()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expected, actual)
}

func (suite *FileServiceSuite) TestFileSizePathIssue() {
	suite.fileRepositoryMock.On("FullFilePath", suite.fileName).Return("", errors.New(""))

	actual, err := suite.target.FileSize()
	assert.Error(suite.T(), err)
	assert.Empty(suite.T(), actual)
}

func (suite *FileServiceSuite) TestFileSizeFileIssue() {
	inputFullPath := faker.Lorem().Word()

	suite.fileRepositoryMock.On("FullFilePath", suite.fileName).Return(inputFullPath, nil)
	suite.fileRepositoryMock.On("FileSize", inputFullPath).Return(uint64(0), errors.New(""))

	actual, err := suite.target.FileSize()
	assert.Error(suite.T(), err)
	assert.Empty(suite.T(), actual)
}

func (suite *FileServiceSuite) TestList() {
	inputFullPath := faker.Lorem().Word()
	expected := map[string]uint64{faker.RandomString(5): uint64(faker.RandomInt(1, 40))}

	suite.fileRepositoryMock.On("FullFilePath", "").Return(inputFullPath, nil)
	suite.fileRepositoryMock.On("List", inputFullPath).Return(expected, nil)

	actual, err := suite.target.List()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expected, actual)
}

func (suite *FileServiceSuite) TestListPathIssue() {
	suite.fileRepositoryMock.On("FullFilePath", "").Return("", errors.New(""))

	actual, err := suite.target.List()
	assert.Error(suite.T(), err)
	assert.Empty(suite.T(), actual)
}

func (suite *FileServiceSuite) TestListListIssue() {
	inputFullPath := faker.Lorem().Word()

	suite.fileRepositoryMock.On("FullFilePath", "").Return(inputFullPath, nil)
	suite.fileRepositoryMock.On("List", inputFullPath).Return(map[string]uint64{}, errors.New(""))

	actual, err := suite.target.List()
	assert.Error(suite.T(), err)
	assert.Empty(suite.T(), actual)
}

func (suite *FileServiceSuite) TestRemove() {
	inputFullPath := faker.Lorem().Word()

	suite.fileRepositoryMock.On("FullFilePath", suite.fileName).Return(inputFullPath, nil)
	suite.fileRepositoryMock.On("Remove", inputFullPath).Return(nil)

	err := suite.target.Remove()
	assert.NoError(suite.T(), err)
}

func (suite *FileServiceSuite) TestRemovePathIssue() {
	suite.fileRepositoryMock.On("FullFilePath", suite.fileName).Return("", errors.New(""))

	err := suite.target.Remove()
	assert.Error(suite.T(), err)
}

func (suite *FileServiceSuite) TestRemoveDeleteIssue() {
	inputFullPath := faker.Lorem().Word()

	suite.fileRepositoryMock.On("FullFilePath", suite.fileName).Return(inputFullPath, nil)
	suite.fileRepositoryMock.On("Remove", inputFullPath).Return(errors.New(""))

	err := suite.target.Remove()
	assert.Error(suite.T(), err)
}

func (suite *FileServiceSuite) TestWrite() {
	input := []byte(faker.RandomString(5))
	expected := faker.RandomInt(5, 10)

	suite.fileRepositoryMock.On("Write", mocks.Anything, input).Return(expected, nil)

	actual, err := suite.target.Write(input)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expected, actual)
}

func (suite *FileServiceSuite) TestRead() {
	expected := []byte(faker.RandomString(5))

	suite.fileRepositoryMock.On("Read", mocks.Anything).Return(expected, nil)

	actual, err := suite.target.Read()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expected, actual)
}

func TestFileServiceSuite(t *testing.T) {
	suite.Run(t, new(FileServiceSuite))
}
