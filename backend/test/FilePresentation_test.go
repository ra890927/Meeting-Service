package test

import (
	"bytes"
	"context"
	"meeting-center/src/models"
	"meeting-center/src/presentations"
	"mime/multipart"
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockFileService struct {
	mock.Mock
}

type FilePresentationTestSuite struct {
	suite.Suite
	fp presentations.FilePresentation
	fs *MockFileService
}

func (m *MockFileService) UploadFile(ctx context.Context, operator models.User, file *models.File, object multipart.File) (string, error) {
	args := m.Called(ctx, operator, file, object)
	return args.String(0), args.Error(1)
}

func (m *MockFileService) GetAllSignedURLsByMeeting(ctx context.Context, operator models.User, meetingID string) ([]models.File, []string, error) {
	args := m.Called(ctx, operator, meetingID)
	return args.Get(0).([]models.File), args.Get(1).([]string), args.Error(2)
}

func (m *MockFileService) GetSignedURL(ctx context.Context, operater models.User, id string) (string, error) {
	args := m.Called(ctx, operater, id)
	return args.String(0), args.Error(1)
}

func (m *MockFileService) DeleteFile(ctx context.Context, operater models.User, id string) error {
	args := m.Called(ctx, operater, id)
	return args.Error(0)
}

func (m *MockFileService) NewFileService() *MockFileService {
	return &MockFileService{}
}

func (suite *FilePresentationTestSuite) SetupTest() {
	suite.fs = new(MockFileService)
	suite.fp = presentations.NewFilePresentation(suite.fs)
}

func (suite *FilePresentationTestSuite) TestNewFilePresentation() {
	// with 0 input
	// TODO: fix this
	// fp_0 := presentations.NewFilePresentation()
	// assert.NotNil(suite.T(), fp_0)

	// with 1 input
	fp_1 := presentations.NewFilePresentation(suite.fs)
	assert.NotNil(suite.T(), fp_1)

	// with 2 inputs
	defer func() {
		if r := recover(); r != nil {
			assert.NotNil(suite.T(), r)
		}
	}()
	_ = presentations.NewFilePresentation(suite.fs, suite.fs)
}

func addFakeUserMiddlewareUserFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := models.User{ID: 1}
		c.Set("validate_user", user)
		c.Next()
	}
}

func (suite *FilePresentationTestSuite) TestUploadFile() {
	// Arrange - Common
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/file", suite.fp.UploadFile)

	// Arrange&Act&Assert - Invalid request (form bind failed)
	reqForm := new(bytes.Buffer)
	req, _ := http.NewRequest(http.MethodPost, "/file", reqForm)
	req.Header.Set("Content-Type", "multipart/form-data")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)

	// Arrange&Act&Assert - Invalid request (form.File is nil)
	reqForm = new(bytes.Buffer)
	writer := multipart.NewWriter(reqForm)
	writer.Close()
	req, _ = http.NewRequest(http.MethodPost, "/file", reqForm)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)

	// Arrange&Act&Assert - Invalid request (form.File.Open failed)
	// TODO: add this test case

	// Arrange&Act&Assert - Error when calling UploadFile in FileService
	reqForm = new(bytes.Buffer)
	writer = multipart.NewWriter(reqForm)
	formFile, _ := writer.CreateFormFile("file", "test.txt")
	formFile.Write([]byte("test"))
	writer.Close()
	suite.fs.On("UploadFile", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return("", assert.AnError).Once()
	r.POST("/file_valid", addFakeUserMiddlewareUserFile(), suite.fp.UploadFile)
	req, _ = http.NewRequest(http.MethodPost, "/file_valid", reqForm)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)

	// Arrange&Act&Assert - Success
	reqForm = new(bytes.Buffer)
	writer = multipart.NewWriter(reqForm)
	formFile, _ = writer.CreateFormFile("file", "test.txt")
	formFile.Write([]byte("test"))
	writer.Close()
	suite.fs.On("UploadFile", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return("test", nil).Once()
	req, _ = http.NewRequest(http.MethodPost, "/file_valid", reqForm)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), http.StatusOK, w.Code)
}

func (suite *FilePresentationTestSuite) TestGetFileURLsByMeetingID() {
	// Arrange - Common
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/file/:meeting_id", suite.fp.GetFileURLsByMeetingID)

	// Arrange&Act&Assert - Invalid request (meeting_id is empty)
	r.GET("/file_empty_file/", suite.fp.GetFileURLsByMeetingID)
	req, _ := http.NewRequest(http.MethodGet, "/file_empty_file/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)

	// Arrange&Act&Assert - Error when calling GetAllSignedURLsByMeeting in FileService
	suite.fs.On("GetAllSignedURLsByMeeting", mock.Anything, mock.Anything, mock.Anything).Return([]models.File{}, []string{}, assert.AnError).Once()
	r.GET("/file_with_user/:meeting_id", addFakeUserMiddlewareUserFile(), suite.fp.GetFileURLsByMeetingID)
	req, _ = http.NewRequest(http.MethodGet, "/file_with_user/1", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)

	// Arrange&Act&Assert - Success
	suite.fs.On("GetAllSignedURLsByMeeting", mock.Anything, mock.Anything, mock.Anything).Return(
		[]models.File{{UploaderID: 1, FileName: "test"}},
		[]string{"test"},
		nil).Once()
	req, _ = http.NewRequest(http.MethodGet, "/file_with_user/1", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), http.StatusOK, w.Code)
}

func (suite *FilePresentationTestSuite) TestGetFile() {
	// Arrange - Common
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/file/:file_id", suite.fp.GetFile)

	// Arrange&Act&Assert - Invalid request (file_id is empty)
	r.GET("/file_empty_file/", suite.fp.GetFile)
	req, _ := http.NewRequest(http.MethodGet, "/file_empty_file/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)

	// Arrange&Act&Assert - Error when calling GetSignedURL in FileService
	suite.fs.On("GetSignedURL", mock.Anything, mock.Anything, mock.Anything).Return("", assert.AnError).Once()
	r.GET("/file_with_user/:id", addFakeUserMiddlewareUserFile(), suite.fp.GetFile)
	req, _ = http.NewRequest(http.MethodGet, "/file_with_user/1", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)

	// Arrange&Act&Assert - Success
	suite.fs.On("GetSignedURL", mock.Anything, mock.Anything, mock.Anything).Return("test", nil).Once()
	req, _ = http.NewRequest(http.MethodGet, "/file_with_user/1", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), http.StatusMovedPermanently, w.Code)
}

func (suite *FilePresentationTestSuite) TestDeleteFile() {
	// Arrange - Common
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.DELETE("/file/:id", suite.fp.DeleteFile)

	// Arrange&Act&Assert - Invalid request (file_id is empty)
	r.DELETE("/file_empty_file/", suite.fp.DeleteFile)
	req, _ := http.NewRequest(http.MethodDelete, "/file_empty_file/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)

	// Arrange&Act&Assert - Error when calling DeleteFile in FileService
	suite.fs.On("DeleteFile", mock.Anything, mock.Anything, mock.Anything).Return(assert.AnError).Once()
	r.DELETE("/file_with_user/:id", addFakeUserMiddlewareUserFile(), suite.fp.DeleteFile)
	req, _ = http.NewRequest(http.MethodDelete, "/file_with_user/1", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)

	// Arrange&Act&Assert - Success
	suite.fs.On("DeleteFile", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
	req, _ = http.NewRequest(http.MethodDelete, "/file_with_user/1", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), http.StatusOK, w.Code)
}

func TestFilePresentationTestSuite(t *testing.T) {
	suite.Run(t, new(FilePresentationTestSuite))
}
