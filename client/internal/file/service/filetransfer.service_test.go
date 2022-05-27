package service_test

import (
	"errors"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	mocks "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/xhsun/grpc-file-transfer/client/internal/config"
	"github.com/xhsun/grpc-file-transfer/client/internal/file/service"
	"github.com/xhsun/grpc-file-transfer/client/internal/mock"
	fileMock "github.com/xhsun/grpc-file-transfer/fileservice/mock"
	pbMock "github.com/xhsun/grpc-file-transfer/filetransfer/mock"
	"syreclabs.com/go/faker"
)

type FileTransferServerSuite struct {
	suite.Suite
	target                     *service.FileTransferServer
	fileServiceBuilderMock     *fileMock.FileServiceBuilderMock
	fileServiceMock            *fileMock.FileServiceMock
	fileTransferRepositoryMock *mock.FileTransferRepositoryMock
	config                     *config.Config
}

// The SetupTest method will be run before every test in the suite
func (suite *FileTransferServerSuite) SetupTest() {
	suite.config = &config.Config{FileChunkSize: faker.RandomInt(5, 10)}
	suite.fileServiceBuilderMock = new(fileMock.FileServiceBuilderMock)
	suite.fileServiceMock = new(fileMock.FileServiceMock)
	suite.fileTransferRepositoryMock = new(mock.FileTransferRepositoryMock)
	suite.target = service.NewFileTransferServer(suite.config, suite.fileServiceBuilderMock, suite.fileTransferRepositoryMock)
}

func (suite *FileTransferServerSuite) TestUpload() {
	inputFileName := faker.Lorem().Word()

	suite.fileServiceBuilderMock.On("WithFile", inputFileName, mocks.Anything).Return(suite.fileServiceBuilderMock)
	suite.fileServiceBuilderMock.On("Build").Return(suite.fileServiceMock, nil)

	suite.fileTransferRepositoryMock.On("UploadStream", mocks.Anything).Return(new(pbMock.FileTransferClientMock), nil)
	suite.fileServiceMock.On("Read").Return([]byte{}, io.EOF)
	suite.fileTransferRepositoryMock.On("UploadClose", mocks.Anything).Return(nil)
	suite.fileServiceMock.On("Close").Return()

	err := suite.target.Upload(inputFileName)
	assert.NoError(suite.T(), err)
}

func (suite *FileTransferServerSuite) TestUploadNoFileName() {
	err := suite.target.Upload("")
	assert.Error(suite.T(), err)
}

func (suite *FileTransferServerSuite) TestUploadFileAccessIssue() {
	inputFileName := faker.Lorem().Word()

	suite.fileServiceBuilderMock.On("WithFile", inputFileName, mocks.Anything).Return(suite.fileServiceBuilderMock)
	suite.fileServiceBuilderMock.On("Build").Return(nil, errors.New(""))

	err := suite.target.Upload(inputFileName)
	assert.Error(suite.T(), err)
}

func (suite *FileTransferServerSuite) TestUploadStreamAccessIssue() {
	inputFileName := faker.Lorem().Word()

	suite.fileServiceBuilderMock.On("WithFile", inputFileName, mocks.Anything).Return(suite.fileServiceBuilderMock)
	suite.fileServiceBuilderMock.On("Build").Return(suite.fileServiceMock, nil)

	suite.fileTransferRepositoryMock.On("UploadStream", mocks.Anything).Return(nil, errors.New(""))
	suite.fileServiceMock.On("Close").Return()

	err := suite.target.Upload(inputFileName)
	assert.Error(suite.T(), err)
}

func (suite *FileTransferServerSuite) TestUploadFileUploadIssue() {
	inputFileName := faker.Lorem().Word()
	inputData := []byte(faker.Lorem().Word())

	suite.fileServiceBuilderMock.On("WithFile", inputFileName, mocks.Anything).Return(suite.fileServiceBuilderMock)
	suite.fileServiceBuilderMock.On("Build").Return(suite.fileServiceMock, nil)

	suite.fileTransferRepositoryMock.On("UploadStream", mocks.Anything).Return(new(pbMock.FileTransferClientMock), nil)
	suite.fileServiceMock.On("Read").Return(inputData, nil)
	suite.fileTransferRepositoryMock.On("Upload", mocks.Anything, inputFileName, mocks.Anything).Return(errors.New(""))
	suite.fileServiceMock.On("Close").Return()

	err := suite.target.Upload(inputFileName)
	assert.Error(suite.T(), err)
}

func (suite *FileTransferServerSuite) TestUploadFileReadIssue() {
	inputFileName := faker.Lorem().Word()

	suite.fileServiceBuilderMock.On("WithFile", inputFileName, mocks.Anything).Return(suite.fileServiceBuilderMock)
	suite.fileServiceBuilderMock.On("Build").Return(suite.fileServiceMock, nil)

	suite.fileTransferRepositoryMock.On("UploadStream", mocks.Anything).Return(new(pbMock.FileTransferClientMock), nil)
	suite.fileServiceMock.On("Read").Return([]byte{}, errors.New(""))
	suite.fileServiceMock.On("Close").Return()

	err := suite.target.Upload(inputFileName)
	assert.Error(suite.T(), err)
}

func (suite *FileTransferServerSuite) TestUploadStreamCloseIssue() {
	inputFileName := faker.Lorem().Word()

	suite.fileServiceBuilderMock.On("WithFile", inputFileName, mocks.Anything).Return(suite.fileServiceBuilderMock)
	suite.fileServiceBuilderMock.On("Build").Return(suite.fileServiceMock, nil)

	suite.fileTransferRepositoryMock.On("UploadStream", mocks.Anything).Return(new(pbMock.FileTransferClientMock), nil)
	suite.fileServiceMock.On("Read").Return([]byte{}, io.EOF)
	suite.fileTransferRepositoryMock.On("UploadClose", mocks.Anything).Return(errors.New(""))
	suite.fileServiceMock.On("Close").Return()

	err := suite.target.Upload(inputFileName)
	assert.Error(suite.T(), err)
}

func (suite *FileTransferServerSuite) TestDownload() {
	inputFileName := faker.Lorem().Word()
	inputLocalFileName := faker.Lorem().Word()

	suite.fileServiceBuilderMock.On("WithFile", inputLocalFileName, mocks.Anything).Return(suite.fileServiceBuilderMock)
	suite.fileServiceBuilderMock.On("Build").Return(suite.fileServiceMock, nil)

	suite.fileTransferRepositoryMock.On("DownloadStream", mocks.Anything, inputFileName).Return(new(pbMock.FileTransferClientMock), nil)
	suite.fileTransferRepositoryMock.On("Download", mocks.Anything).Return([]byte{}, io.EOF)
	suite.fileServiceMock.On("Sync").Return()
	suite.fileServiceMock.On("Close").Return()

	err := suite.target.Download(inputFileName, inputLocalFileName)
	assert.NoError(suite.T(), err)
}

func (suite *FileTransferServerSuite) TestDownloadNoServerFileName() {
	err := suite.target.Download("", faker.Lorem().Word())
	assert.Error(suite.T(), err)
}

func (suite *FileTransferServerSuite) TestDownloadNoLocalFileName() {
	err := suite.target.Download(faker.Lorem().Word(), "")
	assert.Error(suite.T(), err)
}

func (suite *FileTransferServerSuite) TestDownloadFileAccessIssue() {
	inputFileName := faker.Lorem().Word()
	inputLocalFileName := faker.Lorem().Word()

	suite.fileServiceBuilderMock.On("WithFile", inputLocalFileName, mocks.Anything).Return(suite.fileServiceBuilderMock)
	suite.fileServiceBuilderMock.On("Build").Return(nil, errors.New(""))

	err := suite.target.Download(inputFileName, inputLocalFileName)
	assert.Error(suite.T(), err)
}

func (suite *FileTransferServerSuite) TestDownloadStreamIssue() {
	inputFileName := faker.Lorem().Word()
	inputLocalFileName := faker.Lorem().Word()

	suite.fileServiceBuilderMock.On("WithFile", inputLocalFileName, mocks.Anything).Return(suite.fileServiceBuilderMock)
	suite.fileServiceBuilderMock.On("Build").Return(suite.fileServiceMock, nil)

	suite.fileTransferRepositoryMock.On("DownloadStream", mocks.Anything, inputFileName).Return(nil, errors.New(""))
	suite.fileServiceMock.On("Close").Return()

	err := suite.target.Download(inputFileName, inputLocalFileName)
	assert.Error(suite.T(), err)
}

func (suite *FileTransferServerSuite) TestDownloadWriteIssue() {
	inputFileName := faker.Lorem().Word()
	inputLocalFileName := faker.Lorem().Word()
	inputData := []byte(faker.Lorem().Word())

	suite.fileServiceBuilderMock.On("WithFile", inputLocalFileName, mocks.Anything).Return(suite.fileServiceBuilderMock)
	suite.fileServiceBuilderMock.On("Build").Return(suite.fileServiceMock, nil)

	suite.fileTransferRepositoryMock.On("DownloadStream", mocks.Anything, inputFileName).Return(new(pbMock.FileTransferClientMock), nil)
	suite.fileTransferRepositoryMock.On("Download", mocks.Anything).Return(inputData, nil)
	suite.fileServiceMock.On("Write", inputData).Return(0, errors.New(""))
	suite.fileServiceMock.On("Close").Return()

	err := suite.target.Download(inputFileName, inputLocalFileName)
	assert.Error(suite.T(), err)
}

func (suite *FileTransferServerSuite) TestDownloadFetchIssue() {
	inputFileName := faker.Lorem().Word()
	inputLocalFileName := faker.Lorem().Word()

	suite.fileServiceBuilderMock.On("WithFile", inputLocalFileName, mocks.Anything).Return(suite.fileServiceBuilderMock)
	suite.fileServiceBuilderMock.On("Build").Return(suite.fileServiceMock, nil)

	suite.fileTransferRepositoryMock.On("DownloadStream", mocks.Anything, inputFileName).Return(new(pbMock.FileTransferClientMock), nil)
	suite.fileTransferRepositoryMock.On("Download", mocks.Anything).Return([]byte{}, errors.New(""))
	suite.fileServiceMock.On("Close").Return()

	err := suite.target.Download(inputFileName, inputLocalFileName)
	assert.Error(suite.T(), err)
}

func (suite *FileTransferServerSuite) TestServerFileList() {
	expected := map[string]uint64{faker.RandomString(5): uint64(faker.RandomInt(5, 10))}

	suite.fileTransferRepositoryMock.On("ServerFileList", mocks.Anything).Return(expected, nil)

	actual, err := suite.target.ServerFileList()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expected, actual)
}

func (suite *FileTransferServerSuite) TestServerFileListServerIssue() {

	suite.fileTransferRepositoryMock.On("ServerFileList", mocks.Anything).Return(map[string]uint64{}, errors.New(""))

	actual, err := suite.target.ServerFileList()
	assert.Error(suite.T(), err)
	assert.Empty(suite.T(), actual)
}

func (suite *FileTransferServerSuite) TestDeleteFromServer() {
	input := faker.RandomString(5)

	suite.fileTransferRepositoryMock.On("Delete", mocks.Anything, input).Return(nil)

	err := suite.target.DeleteFromServer(input)
	assert.NoError(suite.T(), err)
}

func (suite *FileTransferServerSuite) TestDeleteFromServerServerIssue() {
	input := faker.RandomString(5)

	suite.fileTransferRepositoryMock.On("Delete", mocks.Anything, input).Return(errors.New(""))

	err := suite.target.DeleteFromServer(input)
	assert.Error(suite.T(), err)
}

func TestFileTransferServerSuite(t *testing.T) {
	suite.Run(t, new(FileTransferServerSuite))
}
