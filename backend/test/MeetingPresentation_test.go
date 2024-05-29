package test

import (
	"bytes"
	"encoding/json"
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

func addFakeUserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := models.User{ID: 1}
		c.Set("validate_user", user)
		c.Next()
	}
}

func (m *mockMeetingService) CreateMeeting(operator models.User, meeting *models.Meeting) error {
	args := m.Called(operator, meeting)
	return args.Error(0)
}

func (m *mockMeetingService) UpdateMeeting(operator models.User, meeting *models.Meeting) error {
	args := m.Called(operator, meeting)
	return args.Error(0)
}

func (m *mockMeetingService) DeleteMeeting(operator models.User, id string) error {
	args := m.Called(operator, id)
	return args.Error(0)
}

func (m *mockMeetingService) GetMeeting(id string) (models.Meeting, error) {
	args := m.Called(id)
	return args.Get(0).(models.Meeting), args.Error(1)
}

func (m *mockMeetingService) GetAllMeetings() ([]models.Meeting, error) {
	args := m.Called()
	return args.Get(0).([]models.Meeting), args.Error(1)
}

func (m *mockMeetingService) GetMeetingsByRoomIdAndDatePeriod(roomID int, dateFrom time.Time, dateTo time.Time) ([]models.Meeting, error) {
	args := m.Called(roomID, dateFrom, dateTo)
	return args.Get(0).([]models.Meeting), args.Error(1)
}

func (m *mockMeetingService) GetMeetingsByParticipantId(participantID uint) ([]models.Meeting, error) {
	args := m.Called(participantID)
	return args.Get(0).([]models.Meeting), args.Error(1)
}

func TestCreateMeeting(t *testing.T) {
	meeting := models.Meeting{
		Title:       "Board Meeting",
		Description: "Annual Board Meeting",
	}
	mockMeetingService := new(mockMeetingService)
	mockMeetingService.On("CreateMeeting", models.User{ID: 1}, &meeting).Return(nil)
	mp := presentations.NewMeetingPresentation(mockMeetingService)

	gin.SetMode(gin.TestMode)
	r := gin.Default()

	r.POST("/meeting", addFakeUserMiddleware(), mp.CreateMeeting)

	jsonData, _ := json.Marshal(meeting)
	req := httptest.NewRequest("POST", "/meeting", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestUpdateMeeting(t *testing.T) {
	meeting := models.Meeting{
		ID:          "1",
		Title:       "Updated Board Meeting",
		Description: "Updated Annual Board Meeting",
	}

	mockMeetingService := new(mockMeetingService)
	mockMeetingService.On("UpdateMeeting", models.User{ID: 1}, &meeting).Return(nil)
	mp := presentations.NewMeetingPresentation(mockMeetingService)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.PUT("/meeting", addFakeUserMiddleware(), mp.UpdateMeeting)

	jsonData, _ := json.Marshal(meeting)
	req := httptest.NewRequest("PUT", "/meeting", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestDeleteMeeting(t *testing.T) {
	mockMeetingService := new(mockMeetingService)
	mockMeetingService.On("DeleteMeeting", models.User{ID: 1}, "1").Return(nil)
	mp := presentations.NewMeetingPresentation(mockMeetingService)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.DELETE("/meeting/:id", addFakeUserMiddleware(), mp.DeleteMeeting)

	req := httptest.NewRequest("DELETE", "/meeting/1", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestGetMeeting(t *testing.T) {
	meeting := models.Meeting{
		ID:          "1",
		Title:       "Board Meeting",
		Description: "Annual Board Meeting",
	}

	mockMeetingService := new(mockMeetingService)
	mockMeetingService.On("GetMeeting", "1").Return(meeting, nil)
	mp := presentations.NewMeetingPresentation(mockMeetingService)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/meeting/:id", mp.GetMeeting)

	req := httptest.NewRequest("GET", "/meeting/1", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestGetAllMeetings(t *testing.T) {
	meetings := []models.Meeting{{ID: "1", Title: "Board Meeting", Description: "Annual Board Meeting"}}
	mockMeetingService := new(mockMeetingService)
	mockMeetingService.On("GetAllMeetings").Return(meetings, nil)
	mp := presentations.NewMeetingPresentation(mockMeetingService)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/meetings", mp.GetAllMeetings)

	req := httptest.NewRequest("GET", "/meetings", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestGetMeetingsByRoomIdAndDate(t *testing.T) {
	meetings := []models.Meeting{{ID: "2", Title: "Team Meeting", Description: "Quarterly Planning"}}
	roomID := 101
	mockMeetingService := new(mockMeetingService)
	mockMeetingService.On("GetMeetingsByRoomIdAndDatePeriod", roomID, mock.AnythingOfType("time.Time"), mock.AnythingOfType("time.Time")).Return(meetings, nil)
	// mockMeetingService.On("GetMeetingsByRoomIdAndDate", roomID, date).Return(meetings, nil)
	mp := presentations.NewMeetingPresentation(mockMeetingService)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/meetingsByRoomAndDate", mp.GetMeetingsByRoomIdAndDatePeriod)

	req := httptest.NewRequest("GET", "/meetingsByRoomAndDate?room_id=101&date_from=2021-01-01&date_to=2021-01-01", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestGetMeetingsByParticipantId(t *testing.T) {
	meetings := []models.Meeting{{ID: "2", Title: "Team Meeting", Description: "Quarterly Planning"}}
	mockMeetingService := new(mockMeetingService)
	mockMeetingService.On("GetMeetingsByParticipantId", uint(1)).Return(meetings, nil)
	mp := presentations.NewMeetingPresentation(mockMeetingService)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/meetingsByParticipantId", mp.GetMeetingsByParticipantId)

	req := httptest.NewRequest("GET", "/meetingsByParticipantId?id=1", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
