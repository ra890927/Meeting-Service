package test

import (
	"context"
	"meeting-center/src/models"
	"meeting-center/src/services"
	"mime/multipart"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockFileDomain struct {
	mock.Mock
}

type MockGcsDomain struct {
	mock.Mock
}

type FileServiceTestSuite struct {
	suite.Suite
	fs services.FileService
	fd *MockFileDomain
	gd *MockGcsDomain
	md *MockMeetingDomain
}

func (m *MockFileDomain) UploadFile(file *models.File) error {
	args := m.Called(file)
	return args.Error(0)
}

func (m *MockFileDomain) DeleteFile(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockFileDomain) GetFile(id string) (models.File, error) {
	args := m.Called(id)
	return args.Get(0).(models.File), args.Error(1)
}

func (m *MockFileDomain) GetAllFilesByMeetingID(meetingID string) ([]models.File, error) {
	args := m.Called(meetingID)
	return args.Get(0).([]models.File), args.Error(1)
}

func (m *MockFileDomain) NewFileDomain() *MockFileDomain {
	return &MockFileDomain{}
}

func (m *MockGcsDomain) UploadFile(ctx context.Context, file multipart.File, objectID string, objExt string) error {
	args := m.Called(ctx, file, objectID, objExt)
	return args.Error(0)
}

func (m *MockGcsDomain) GetSignedURL(ctx context.Context, objectID string, objExt string) (string, error) {
	args := m.Called(ctx, objectID, objExt)
	return args.String(0), args.Error(1)
}

func (m *MockGcsDomain) DeleteFile(ctx context.Context, objectID string, objExt string) error {
	args := m.Called(ctx, objectID, objExt)
	return args.Error(0)
}

func (suite *FileServiceTestSuite) SetupTest() {
	suite.fd = new(MockFileDomain)
	suite.gd = new(MockGcsDomain)
	suite.md = new(MockMeetingDomain)
	suite.fs = services.NewFileService(
		services.FileServiceArg{
			FileDomain:    suite.fd,
			GcsDomain:     suite.gd,
			MeetingDomain: suite.md,
		},
	)
}

func (suite *FileServiceTestSuite) TestNewFileService() {
	// 0 input
	// TODO: find a way to test when no input is given

	// 1 input
	fsArg := services.FileServiceArg{
		FileDomain:    suite.fd,
		GcsDomain:     suite.gd,
		MeetingDomain: suite.md,
	}
	fs := services.NewFileService(fsArg)
	assert.NotNil(suite.T(), fs)

	// more than 1 input
	assert.Panics(suite.T(), func() {
		services.NewFileService(fsArg, fsArg)
	})
}

func (suite *FileServiceTestSuite) TestUploadFile() {
	// error when getPermissionByFile
	file := models.File{
		ID:        "fileID",
		MeetingID: "meetingID",
		FileName:  "fileName.txt",
		FileExt:   ".txt",
	}
	operator := models.User{}
	object := new(multipart.File)
	suite.fd.On("UploadFile", &file).Return(nil).Once()
	suite.gd.On("UploadFile", context.Background(), *object, file.ID, file.FileExt).Return(nil).Once()
	suite.gd.On("GetSignedURL", context.Background(), file.ID, file.FileExt).Return("url", nil).Once()
	suite.md.On("GetMeeting", file.MeetingID).Return(models.Meeting{}, nil).Once()
	_, err := suite.fs.UploadFile(context.Background(), operator, &file, *object)
	assert.Nil(suite.T(), err)

	// error when UploadFile in GcsDomain
	suite.fd.On("UploadFile", &file).Return(nil).Once()
	suite.gd.On("UploadFile", context.Background(), *object, file.ID, file.FileExt).Return(assert.AnError).Once()
	suite.gd.On("GetSignedURL", context.Background(), file.ID, file.FileExt).Return("url", nil).Once()
	suite.md.On("GetMeeting", file.MeetingID).Return(models.Meeting{}, nil).Once()
	_, err = suite.fs.UploadFile(context.Background(), operator, &file, *object)
	assert.NotNil(suite.T(), err)

}

func (suite *FileServiceTestSuite) TestGetAllSignedURLsByMeeting() {
	meetingID := "meetingID"
	operator := models.User{}
	files := []models.File{
		{ID: "fileID", FileName: "fileName.txt", FileExt: ".txt"},
	}
	urls := []string{"url"}

	// normal case
	suite.md.On("GetMeeting", meetingID).Return(models.Meeting{}, nil).Once()
	suite.fd.On("GetAllFilesByMeetingID", meetingID).Return(files, nil).Once()
	suite.gd.On("GetSignedURL", context.Background(), files[0].ID, files[0].FileExt).Return(urls[0], nil).Once()
	_, _, err := suite.fs.GetAllSignedURLsByMeeting(context.Background(), operator, meetingID)
	assert.Nil(suite.T(), err)

	// error when GetMeeting
	suite.md.On("GetMeeting", meetingID).Return(models.Meeting{}, assert.AnError).Once()
	_, _, err = suite.fs.GetAllSignedURLsByMeeting(context.Background(), operator, meetingID)
	assert.NotNil(suite.T(), err)

	// error when GetAllFilesByMeetingID
	suite.md.On("GetMeeting", meetingID).Return(models.Meeting{}, nil).Once()
	suite.fd.On("GetAllFilesByMeetingID", meetingID).Return([]models.File{}, assert.AnError).Once()
	_, _, err = suite.fs.GetAllSignedURLsByMeeting(context.Background(), operator, meetingID)
	assert.NotNil(suite.T(), err)

	// error when fs.gcsDomain.GetSignedURL
	suite.md.On("GetMeeting", meetingID).Return(models.Meeting{}, nil).Once()
	suite.fd.On("GetAllFilesByMeetingID", meetingID).Return(files, nil).Once()
	suite.gd.On("GetSignedURL", context.Background(), files[0].ID, files[0].FileExt).Return("", assert.AnError).Once()
	_, _, err = suite.fs.GetAllSignedURLsByMeeting(context.Background(), operator, meetingID)
	assert.NotNil(suite.T(), err)
}

func (suite *FileServiceTestSuite) TestGetSignedURL() {
	operator := models.User{}
	files := []models.File{
		{ID: "fileID", FileName: "fileName.txt", FileExt: ".txt"},
	}
	id := "fileID"

	// normal case
	suite.fd.On("GetFile", id).Return(files[0], nil).Once()
	suite.md.On("GetMeeting", files[0].MeetingID).Return(models.Meeting{}, nil).Once()
	suite.gd.On("GetSignedURL", context.Background(), files[0].ID, files[0].FileExt).Return("url", nil).Once()
	_, err := suite.fs.GetSignedURL(context.Background(), operator, id)
	assert.Nil(suite.T(), err)

	// error when GetFile
	suite.fd.On("GetFile", id).Return(models.File{}, assert.AnError).Once()
	_, err = suite.fs.GetSignedURL(context.Background(), operator, id)
	assert.NotNil(suite.T(), err)
}

func (suite *FileServiceTestSuite) TestDeleteFile() {
	operator := models.User{}
	file := models.File{
		ID:        "fileID",
		MeetingID: "meetingID",
		FileName:  "fileName.txt",
		FileExt:   ".txt",
	}

	// normal case
	suite.fd.On("GetFile", file.ID).Return(file, nil).Once()
	suite.md.On("GetMeeting", file.MeetingID).Return(models.Meeting{}, nil).Once()
	suite.fd.On("DeleteFile", file.ID).Return(nil).Once()
	suite.gd.On("DeleteFile", context.Background(), file.ID, file.FileExt).Return(nil).Once()
	err := suite.fs.DeleteFile(context.Background(), operator, file.ID)
	assert.Nil(suite.T(), err)

	// error when GetFile
	suite.fd.On("GetFile", file.ID).Return(models.File{}, assert.AnError).Once()
	err = suite.fs.DeleteFile(context.Background(), operator, file.ID)
	assert.NotNil(suite.T(), err)

	// gcsDomain.DeleteFile failed
	suite.fd.On("GetFile", file.ID).Return(file, nil).Once()
	suite.md.On("GetMeeting", file.MeetingID).Return(models.Meeting{}, nil).Once()
	suite.gd.On("DeleteFile", context.Background(), file.ID, file.FileExt).Return(assert.AnError).Once()
	err = suite.fs.DeleteFile(context.Background(), operator, file.ID)
	assert.NotNil(suite.T(), err)

	// fileDomain.DeleteFile failed
	suite.fd.On("GetFile", file.ID).Return(file, nil).Once()
	suite.md.On("GetMeeting", file.MeetingID).Return(models.Meeting{}, nil).Once()
	suite.gd.On("DeleteFile", context.Background(), file.ID, file.FileExt).Return(nil).Once()
	suite.fd.On("DeleteFile", file.ID).Return(assert.AnError).Once()
	err = suite.fs.DeleteFile(context.Background(), operator, file.ID)
	assert.NotNil(suite.T(), err)
}

func TestFileServiceTestSuite(t *testing.T) {
	suite.Run(t, new(FileServiceTestSuite))
}
