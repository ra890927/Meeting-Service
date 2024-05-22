package test

import (
	"bytes"
	"encoding/json"
	"errors"
	"meeting-center/src/models"
	"meeting-center/src/presentations"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockMeetingService struct {
	mock.Mock
}

func (m *mockMeetingService) CreateMeeting(meeting *models.Meeting) (*models.Meeting, error) {
	args := m.Called(meeting)
	if args.Get(0) != nil {
		return args.Get(0).(*models.Meeting), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *mockMeetingService) UpdateMeeting(id string, meeting *models.Meeting) error {
	args := m.Called(id, meeting)
	return args.Error(0)
}

func (m *mockMeetingService) DeleteMeeting(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *mockMeetingService) GetMeeting(id string) (*models.Meeting, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*models.Meeting), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *mockMeetingService) GetAllMeetings() ([]*models.Meeting, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]*models.Meeting), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *mockMeetingService) GetMeetingsByRoomIdAndDate(roomID int, date time.Time) ([]*models.Meeting, error) {
	args := m.Called(roomID, date)
	if args.Get(0) != nil {
		return args.Get(0).([]*models.Meeting), args.Error(1)
	}
	return nil, args.Error(1)
}

// 新增測試NewMeetingPresentation的測試函數
func TestNewMeetingPresentation(t *testing.T) {
	mockMeetingService := new(mockMeetingService)
	mp := presentations.NewMeetingPresentation(mockMeetingService)
	assert.NotNil(t, mp)

	mp = presentations.NewMeetingPresentation()
	assert.NotNil(t, mp)
}

func TestCreateMeeting(t *testing.T) {
	meeting := models.Meeting{
		ID:          "1",
		Title:       "Board Meeting",
		Description: "Annual Board Meeting",
	}

	mockMeetingService := new(mockMeetingService)
	mockMeetingService.On("CreateMeeting", &meeting).Return(&meeting, nil).Once()
	mockMeetingService.On("CreateMeeting", &meeting).Return(nil, errors.New("internal server error")).Once()
	mp := presentations.NewMeetingPresentation(mockMeetingService)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/meeting", mp.CreateMeeting)

	// Test successful creation
	jsonData, _ := json.Marshal(meeting)
	req := httptest.NewRequest("POST", "/meeting", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	// Test invalid request
	req = httptest.NewRequest("POST", "/meeting", bytes.NewBuffer([]byte(`invalid json`)))
	req.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)

	// Test internal server error
	req = httptest.NewRequest("POST", "/meeting", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 500, w.Code)
}

func TestUpdateMeeting(t *testing.T) {
	meeting := models.Meeting{
		ID:          "1",
		Title:       "Updated Board Meeting",
		Description: "Updated Annual Board Meeting",
	}

	mockMeetingService := new(mockMeetingService)
	mockMeetingService.On("UpdateMeeting", "1", &meeting).Return(nil).Once()
	mockMeetingService.On("UpdateMeeting", "1", &meeting).Return(errors.New("internal server error")).Once()
	mp := presentations.NewMeetingPresentation(mockMeetingService)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.PUT("/meeting/:id", mp.UpdateMeeting)

	// Test successful update
	jsonData, _ := json.Marshal(meeting)
	req := httptest.NewRequest("PUT", "/meeting/1", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	// Test invalid request
	req = httptest.NewRequest("PUT", "/meeting/1", bytes.NewBuffer([]byte(`invalid json`)))
	req.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)

	// Test internal server error
	req = httptest.NewRequest("PUT", "/meeting/1", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 500, w.Code)
}

func TestDeleteMeeting(t *testing.T) {
	mockMeetingService := new(mockMeetingService)
	mockMeetingService.On("DeleteMeeting", "1").Return(nil).Once()
	mockMeetingService.On("DeleteMeeting", "1").Return(errors.New("internal server error")).Once()
	mp := presentations.NewMeetingPresentation(mockMeetingService)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.DELETE("/meeting/:id", mp.DeleteMeeting)

	// Test successful deletion
	req := httptest.NewRequest("DELETE", "/meeting/1", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	// Test internal server error
	req = httptest.NewRequest("DELETE", "/meeting/1", nil)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 500, w.Code)
}

func TestGetMeeting(t *testing.T) {
	meeting := models.Meeting{
		ID:          "1",
		Title:       "Board Meeting",
		Description: "Annual Board Meeting",
	}

	mockMeetingService := new(mockMeetingService)
	mockMeetingService.On("GetMeeting", "1").Return(&meeting, nil).Once()
	mockMeetingService.On("GetMeeting", "1").Return(nil, errors.New("internal server error")).Once()
	mp := presentations.NewMeetingPresentation(mockMeetingService)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/meeting/:id", mp.GetMeeting)

	// Test successful retrieval
	req := httptest.NewRequest("GET", "/meeting/1", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	// Test internal server error
	req = httptest.NewRequest("GET", "/meeting/1", nil)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 500, w.Code)
}

func TestGetAllMeetings(t *testing.T) {
	meetings := []*models.Meeting{{ID: "1", Title: "Board Meeting", Description: "Annual Board Meeting"}}
	mockMeetingService := new(mockMeetingService)
	mockMeetingService.On("GetAllMeetings").Return(meetings, nil).Once()
	mockMeetingService.On("GetAllMeetings").Return(nil, errors.New("internal server error")).Once()
	mp := presentations.NewMeetingPresentation(mockMeetingService)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/meetings", mp.GetAllMeetings)

	// Test successful retrieval
	req := httptest.NewRequest("GET", "/meetings", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	// Test internal server error
	req = httptest.NewRequest("GET", "/meetings", nil)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 500, w.Code)
}

func TestGetMeetingsByRoomIdAndDate(t *testing.T) {
	meetings := []*models.Meeting{{ID: "2", Title: "Team Meeting", Description: "Quarterly Planning"}}
	roomID := 101
	date := time.Now()
	mockMeetingService := new(mockMeetingService)
	mockMeetingService.On("GetMeetingsByRoomIdAndDate", roomID, mock.AnythingOfType("time.Time")).Return(meetings, nil).Once()
	mockMeetingService.On("GetMeetingsByRoomIdAndDate", roomID, mock.AnythingOfType("time.Time")).Return(nil, errors.New("internal server error")).Once()
	mp := presentations.NewMeetingPresentation(mockMeetingService)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/meetingsByRoomAndDate", mp.GetMeetingsByRoomIdAndDate)

	// Test successful retrieval
	req := httptest.NewRequest("GET", "/meetingsByRoomAndDate?room_id=101&date="+date.Format("2006-01-02"), nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	// Test missing parameters
	req = httptest.NewRequest("GET", "/meetingsByRoomAndDate?room_id=101", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)

	req = httptest.NewRequest("GET", "/meetingsByRoomAndDate?date="+date.Format("2006-01-02"), nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)

	// Test internal server error
	req = httptest.NewRequest("GET", "/meetingsByRoomAndDate?room_id=101&date="+date.Format("2006-01-02"), nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 500, w.Code)
}
