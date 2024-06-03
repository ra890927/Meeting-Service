package test

import (
	"meeting-center/src/domains"
	"meeting-center/src/models"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockFileRepo struct {
	mock.Mock
}

type FileDomainTestSuite struct {
	suite.Suite
	fd domains.FileDomain
	fr *MockFileRepo
}

func (m *MockFileRepo) UploadFile(file *models.File) error {
	args := m.Called(file)
	return args.Error(0)
}

func (m *MockFileRepo) DeleteFile(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockFileRepo) GetFile(id string) (models.File, error) {
	args := m.Called(id)
	return args.Get(0).(models.File), args.Error(1)
}

func (m *MockFileRepo) GetFilesByMeetingID(meetingID string) ([]models.File, error) {
	args := m.Called(meetingID)
	return args.Get(0).([]models.File), args.Error(1)
}

func (m *MockFileRepo) NewFileRepo() *MockFileRepo {
	return &MockFileRepo{}
}

func (suite *FileDomainTestSuite) SetupTest() {
	suite.fr = suite.fr.NewFileRepo()
	suite.fd = domains.NewFileDomain(suite.fr)
}

func (suite *FileDomainTestSuite) TestNewFileDomain() {
	// 0 arguments
	// TODO: find a way to test when 0 arguments

	// 1 argument
	suite.fr.On("NewFileDomain", suite.fr).Return(nil)
	fd := domains.NewFileDomain(suite.fr)
	assert.NotNil(suite.T(), fd)

	// more than 1 argument
	assert.Panics(suite.T(), func() { domains.NewFileDomain(suite.fr, suite.fr) })
}

func (suite *FileDomainTestSuite) TestUploadFile() {
	file := models.File{}
	suite.fr.On("UploadFile", &file).Return(nil)
	err := suite.fd.UploadFile(&file)
	assert.Nil(suite.T(), err)
}

func (suite *FileDomainTestSuite) TestDeleteFile() {
	id := "test"
	suite.fr.On("DeleteFile", id).Return(nil)
	err := suite.fd.DeleteFile(id)
	assert.Nil(suite.T(), err)
}

func (suite *FileDomainTestSuite) TestGetFile() {
	id := "test"
	file := models.File{}
	suite.fr.On("GetFile", id).Return(file, nil)
	result, err := suite.fd.GetFile(id)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), file, result)
}

func (suite *FileDomainTestSuite) TestGetAllFilesByMeetingID() {
	meetingID := "test"
	files := []models.File{}
	suite.fr.On("GetFilesByMeetingID", meetingID).Return(files, nil)
	result, err := suite.fd.GetAllFilesByMeetingID(meetingID)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), files, result)
}

func TestFileDomainTestSuite(t *testing.T) {
	suite.Run(t, new(FileDomainTestSuite))
}
