package file_test

import (
	"errors"
	"flag"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/urfave/cli/v2"
	"github.com/xhsun/grpc-file-transfer/client/internal/file"
	"github.com/xhsun/grpc-file-transfer/client/internal/mock"
	"syreclabs.com/go/faker"
)

type FileCommandSuite struct {
	suite.Suite
	target                  *file.FileCommand
	fileTransferServiceMock *mock.FileTransferServiceMock
}

// The SetupTest method will be run before every test in the suite
func (suite *FileCommandSuite) SetupTest() {
	suite.fileTransferServiceMock = new(mock.FileTransferServiceMock)
	suite.target = file.NewFileCommand(suite.fileTransferServiceMock)
}

func (suite *FileCommandSuite) TestUpload() {
	inputFileName := faker.RandomString(5)
	inputContext := suite.CliContextFaker([]string{inputFileName})

	suite.fileTransferServiceMock.On("Upload", inputFileName).Return(nil)

	err := suite.target.Upload(inputContext)
	assert.NoError(suite.T(), err)
}

func (suite *FileCommandSuite) TestUploadError() {
	inputFileName := faker.RandomString(5)
	inputContext := suite.CliContextFaker([]string{inputFileName})

	suite.fileTransferServiceMock.On("Upload", inputFileName).Return(errors.New(""))

	err := suite.target.Upload(inputContext)
	assert.Error(suite.T(), err)
}

func (suite *FileCommandSuite) TestDownload() {
	inputFileName := faker.RandomString(5)
	inputLocalFileName := faker.RandomString(5)
	inputContext := suite.CliContextFaker([]string{inputFileName, inputLocalFileName})

	suite.fileTransferServiceMock.On("Download", inputFileName, inputLocalFileName).Return(nil)

	err := suite.target.Download(inputContext)
	assert.NoError(suite.T(), err)
}

func (suite *FileCommandSuite) TestDownloadError() {
	inputFileName := faker.RandomString(5)
	inputLocalFileName := faker.RandomString(5)
	inputContext := suite.CliContextFaker([]string{inputFileName, inputLocalFileName})

	suite.fileTransferServiceMock.On("Download", inputFileName, inputLocalFileName).Return(errors.New(""))

	err := suite.target.Download(inputContext)
	assert.Error(suite.T(), err)
}

func (suite *FileCommandSuite) TestListFiles() {
	input := map[string]uint64{faker.RandomString(5): uint64(faker.RandomInt(5, 10))}

	suite.fileTransferServiceMock.On("ServerFileList").Return(input, nil)

	err := suite.target.ListFiles(nil)
	assert.NoError(suite.T(), err)
}

func (suite *FileCommandSuite) TestListFilesEmptyList() {
	suite.fileTransferServiceMock.On("ServerFileList").Return(map[string]uint64{}, nil)

	err := suite.target.ListFiles(nil)
	assert.NoError(suite.T(), err)
}

func (suite *FileCommandSuite) TestListFilesError() {
	suite.fileTransferServiceMock.On("ServerFileList").Return(map[string]uint64{}, errors.New(""))

	err := suite.target.ListFiles(nil)
	assert.Error(suite.T(), err)
}

// CliContextFaker for faking *cli.Context
func (suite *FileCommandSuite) CliContextFaker(args []string) *cli.Context {
	set := flag.NewFlagSet("", flag.ContinueOnError)
	_ = set.Parse(args)
	return cli.NewContext(nil, set, nil)
}

func (suite *FileCommandSuite) TestRemove() {
	inputFileName := faker.RandomString(5)
	inputContext := suite.CliContextFaker([]string{inputFileName})

	suite.fileTransferServiceMock.On("DeleteFromServer", inputFileName).Return(nil)

	err := suite.target.Remove(inputContext)
	assert.NoError(suite.T(), err)
}

func (suite *FileCommandSuite) TestRemoveError() {
	inputFileName := faker.RandomString(5)
	inputContext := suite.CliContextFaker([]string{inputFileName})

	suite.fileTransferServiceMock.On("DeleteFromServer", inputFileName).Return(errors.New(""))

	err := suite.target.Remove(inputContext)
	assert.Error(suite.T(), err)
}

func TestFileCommandSuite(t *testing.T) {
	suite.Run(t, new(FileCommandSuite))
}
