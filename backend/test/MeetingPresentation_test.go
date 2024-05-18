package test

import (
	"bytes"
	"encoding/json"
	"meeting-center/src/models"
	"meeting-center/src/presentations"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockMeetingService struct {
	mock.Mock
}

func (m *mockMeetingService) CreateMeeting(meeting *models.Meeting) (*models.Meeting, error) {
	args := m.Called(meeting)
	return args.Get(0).(*models.Meeting), args.Error(1)
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
	return args.Get(0).(*models.Meeting), args.Error(1)
}

func TestCreateMeeting(t *testing.T) {
	meeting := models.Meeting{
		ID:          "1",
		Title:       "Board Meeting",
		Description: "Annual Board Meeting",
	}

	mockMeetingService := new(mockMeetingService)
	mockMeetingService.On("CreateMeeting", &meeting).Return(&meeting, nil)
	mp := presentations.NewMeetingPresentation(mockMeetingService)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/meeting", mp.CreateMeeting)

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
	mockMeetingService.On("UpdateMeeting", "1", &meeting).Return(nil)
	mp := presentations.NewMeetingPresentation(mockMeetingService)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.PUT("/meeting/:id", mp.UpdateMeeting)

	jsonData, _ := json.Marshal(meeting)
	req := httptest.NewRequest("PUT", "/meeting/1", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestDeleteMeeting(t *testing.T) {
	mockMeetingService := new(mockMeetingService)
	mockMeetingService.On("DeleteMeeting", "1").Return(nil)
	mp := presentations.NewMeetingPresentation(mockMeetingService)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.DELETE("/meeting/:id", mp.DeleteMeeting)

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
	mockMeetingService.On("GetMeeting", "1").Return(&meeting, nil)
	mp := presentations.NewMeetingPresentation(mockMeetingService)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/meeting/:id", mp.GetMeeting)

	req := httptest.NewRequest("GET", "/meeting/1", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
