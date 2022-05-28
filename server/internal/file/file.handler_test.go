package file_test

import (
	"context"
	"errors"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	mocks "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	pb "github.com/xhsun/grpc-file-transfer/dfs"
	pbMock "github.com/xhsun/grpc-file-transfer/dfs/mock"
	fileMock "github.com/xhsun/grpc-file-transfer/fileservice/mock"
	"github.com/xhsun/grpc-file-transfer/server/internal/config"
	"github.com/xhsun/grpc-file-transfer/server/internal/file"
	"syreclabs.com/go/faker"
)

type FileHandlerSuite struct {
	suite.Suite
	target                 *file.FileHandler
	fileServiceBuilderMock *fileMock.FileServiceBuilderMock
	fileServiceMock        *fileMock.FileServiceMock
	config                 *config.Config
	streamMock             *pbMock.FileTransferServerMock
}

// The SetupTest method will be run before every test in the suite
func (suite *FileHandlerSuite) SetupTest() {
	suite.config = &config.Config{FileChunkSize: faker.RandomInt(5, 10), FileStoragePath: faker.RandomString(5)}
	suite.fileServiceBuilderMock = new(fileMock.FileServiceBuilderMock)
	suite.fileServiceMock = new(fileMock.FileServiceMock)
	suite.streamMock = new(pbMock.FileTransferServerMock)
	suite.target = file.NewFileHandler(suite.config, suite.fileServiceBuilderMock)
}

func (suite *FileHandlerSuite) TestStore() {
	suite.streamMock.On("Recv").Return(nil, io.EOF)
	suite.streamMock.On("SendAndClose", mocks.Anything).Return(nil)

	err := suite.target.Store(suite.streamMock)
	assert.NoError(suite.T(), err)
}

func (suite *FileHandlerSuite) TestStoreFileAccessIssue() {
	inputFile := suite.FileFaker()

	suite.streamMock.On("Recv").Return(inputFile, nil)
	suite.fileServiceBuilderMock.On("WithFile", inputFile.Name, mocks.Anything).Return(suite.fileServiceBuilderMock)
	suite.fileServiceBuilderMock.On("Build").Return(nil, errors.New(""))

	err := suite.target.Store(suite.streamMock)
	assert.Error(suite.T(), err)
}

func (suite *FileHandlerSuite) TestStoreWriteIssue() {
	inputFile := suite.FileFaker()

	suite.streamMock.On("Recv").Return(inputFile, nil)
	suite.fileServiceBuilderMock.On("WithFile", inputFile.Name, mocks.Anything).Return(suite.fileServiceBuilderMock)
	suite.fileServiceBuilderMock.On("Build").Return(suite.fileServiceMock, nil)

	suite.fileServiceMock.On("Write", inputFile.Content).Return(0, errors.New(""))
	suite.fileServiceMock.On("Close").Return()

	err := suite.target.Store(suite.streamMock)
	assert.Error(suite.T(), err)
}

func (suite *FileHandlerSuite) TestStoreFetchIssue() {
	suite.streamMock.On("Recv").Return(nil, errors.New(""))

	err := suite.target.Store(suite.streamMock)
	assert.Error(suite.T(), err)
}

func (suite *FileHandlerSuite) TestFetch() {
	inputFileName := suite.FileNameFaker()

	suite.fileServiceBuilderMock.On("WithFile", inputFileName.Name, mocks.Anything).Return(suite.fileServiceBuilderMock)
	suite.fileServiceBuilderMock.On("Build").Return(suite.fileServiceMock, nil)
	suite.fileServiceMock.On("FileSize").Return(uint64(faker.RandomInt64(5, 10)), nil)

	suite.fileServiceMock.On("Read").Return([]byte{}, io.EOF)
	suite.fileServiceMock.On("Close").Return()

	err := suite.target.Fetch(inputFileName, suite.streamMock)
	assert.NoError(suite.T(), err)
}

func (suite *FileHandlerSuite) TestFetchNilFileName() {
	err := suite.target.Fetch(nil, suite.streamMock)
	assert.Error(suite.T(), err)
}

func (suite *FileHandlerSuite) TestFetchNoFileName() {
	err := suite.target.Fetch(&pb.FileName{}, suite.streamMock)
	assert.Error(suite.T(), err)
}

func (suite *FileHandlerSuite) TestFetchFileAccessIssue() {
	inputFileName := suite.FileNameFaker()

	suite.fileServiceBuilderMock.On("WithFile", inputFileName.Name, mocks.Anything).Return(suite.fileServiceBuilderMock)
	suite.fileServiceBuilderMock.On("Build").Return(nil, errors.New(""))

	err := suite.target.Fetch(inputFileName, suite.streamMock)
	assert.Error(suite.T(), err)
}

func (suite *FileHandlerSuite) TestFetchFileReadIssue() {
	inputFileName := suite.FileNameFaker()

	suite.fileServiceBuilderMock.On("WithFile", inputFileName.Name, mocks.Anything).Return(suite.fileServiceBuilderMock)
	suite.fileServiceBuilderMock.On("Build").Return(suite.fileServiceMock, nil)
	suite.fileServiceMock.On("FileSize").Return(uint64(faker.RandomInt64(5, 10)), nil)

	suite.fileServiceMock.On("Read").Return([]byte{}, errors.New(""))
	suite.fileServiceMock.On("Close").Return()

	err := suite.target.Fetch(inputFileName, suite.streamMock)
	assert.Error(suite.T(), err)
}

func (suite *FileHandlerSuite) TestFetchFileSendIssue() {
	inputFileName := suite.FileNameFaker()
	inputData := []byte(faker.Lorem().Word())

	suite.fileServiceBuilderMock.On("WithFile", inputFileName.Name, mocks.Anything).Return(suite.fileServiceBuilderMock)
	suite.fileServiceBuilderMock.On("Build").Return(suite.fileServiceMock, nil)
	suite.fileServiceMock.On("FileSize").Return(uint64(faker.RandomInt64(5, 10)), nil)

	suite.fileServiceMock.On("Read").Return(inputData, nil)
	suite.streamMock.On("Send", mocks.Anything).Return(errors.New(""))
	suite.fileServiceMock.On("Close").Return()

	err := suite.target.Fetch(inputFileName, suite.streamMock)
	assert.Error(suite.T(), err)
}

func (suite *FileHandlerSuite) TestDelete() {
	inputFileName := suite.FileNameFaker()

	suite.fileServiceBuilderMock.On("WithFileName", inputFileName.Name).Return(suite.fileServiceBuilderMock)
	suite.fileServiceBuilderMock.On("Build").Return(suite.fileServiceMock, nil)

	suite.fileServiceMock.On("Remove").Return(nil)

	_, err := suite.target.Delete(context.TODO(), inputFileName)
	assert.NoError(suite.T(), err)
}

func (suite *FileHandlerSuite) TestDeleteNilFileName() {
	_, err := suite.target.Delete(context.TODO(), nil)
	assert.Error(suite.T(), err)
}

func (suite *FileHandlerSuite) TestDeleteNoFileName() {
	_, err := suite.target.Delete(context.TODO(), &pb.FileName{})
	assert.Error(suite.T(), err)
}

func (suite *FileHandlerSuite) TestDeleteError() {
	inputFileName := suite.FileNameFaker()

	suite.fileServiceBuilderMock.On("WithFileName", inputFileName.Name).Return(suite.fileServiceBuilderMock)
	suite.fileServiceBuilderMock.On("Build").Return(suite.fileServiceMock, nil)

	suite.fileServiceMock.On("Remove").Return(errors.New(""))

	_, err := suite.target.Delete(context.TODO(), inputFileName)
	assert.Error(suite.T(), err)
}

func (suite *FileHandlerSuite) TestListAll() {
	expected := map[string]uint64{faker.RandomString(5): uint64(faker.RandomInt(5, 10))}

	suite.fileServiceBuilderMock.On("Build").Return(suite.fileServiceMock, nil)

	suite.fileServiceMock.On("List").Return(expected, nil)

	actual, err := suite.target.ListAll(context.TODO(), &pb.Empty{})
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expected, actual.Files)
}

func (suite *FileHandlerSuite) TestListAllError() {
	suite.fileServiceBuilderMock.On("Build").Return(suite.fileServiceMock, nil)

	suite.fileServiceMock.On("List").Return(map[string]uint64{}, errors.New(""))

	actual, err := suite.target.ListAll(context.TODO(), &pb.Empty{})
	assert.Error(suite.T(), err)
	assert.Empty(suite.T(), actual)
}

func TestFileHandlerSuite(t *testing.T) {
	suite.Run(t, new(FileHandlerSuite))
}

func (suite *FileHandlerSuite) FileFaker() *pb.File {
	return &pb.File{
		Name:    faker.RandomString(5),
		Content: []byte(faker.RandomString(5)),
	}
}

func (suite *FileHandlerSuite) FileNameFaker() *pb.FileName {
	return &pb.FileName{
		Name: faker.RandomString(5),
	}
}
