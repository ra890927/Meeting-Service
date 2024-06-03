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

func addFakeUserMiddlewareMeeting() gin.HandlerFunc {
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

func TestNewMeetingPresentation(t *testing.T) {
	// t.Run("NewMeetingPresentation with 0 input", func(t *testing.T) {
	// 	mp := presentations.NewMeetingPresentation()
	// 	assert.NotNil(t, mp)
	// })

	t.Run("NewMeetingPresentation with 1 input", func(t *testing.T) {
		mockMeetingService := new(mockMeetingService)
		mp := presentations.NewMeetingPresentation(mockMeetingService)
		assert.NotNil(t, mp)
	})

	t.Run("NewMeetingPresentation with more than 1 input", func(t *testing.T) {
		mockMeetingService := new(mockMeetingService)

		defer func() {
			if r := recover(); r != nil {
				assert.NotNil(t, r)
			}
		}()
		presentations.NewMeetingPresentation(mockMeetingService, mockMeetingService)
	})
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

	r.POST("/meeting", addFakeUserMiddlewareMeeting(), mp.CreateMeeting)

	jsonData, _ := json.Marshal(meeting)
	req := httptest.NewRequest("POST", "/meeting", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestCreateMeeting_InvalidRequest(t *testing.T) {
	meeting := models.Meeting{
		Title:       "Board Meeting",
		Description: "Annual Board Meeting",
	}
	mockMeetingService := new(mockMeetingService)
	mockMeetingService.On("CreateMeeting", models.User{ID: 1}, &meeting).Return(nil)
	mp := presentations.NewMeetingPresentation(mockMeetingService)

	gin.SetMode(gin.TestMode)
	r := gin.Default()

	r.POST("/meeting", addFakeUserMiddlewareMeeting(), mp.CreateMeeting)

	jsonData := []byte(`{"title": 123}`)
	req := httptest.NewRequest("POST", "/meeting", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func TestCreateMeeting_ErrorInService(t *testing.T) {
	meeting := models.Meeting{
		Title:       "Board Meeting",
		Description: "Annual Board Meeting",
	}
	mockMeetingService := new(mockMeetingService)
	mockMeetingService.On("CreateMeeting", models.User{ID: 1}, &meeting).Return(assert.AnError)
	mp := presentations.NewMeetingPresentation(mockMeetingService)

	gin.SetMode(gin.TestMode)
	r := gin.Default()

	r.POST("/meeting", addFakeUserMiddlewareMeeting(), mp.CreateMeeting)

	jsonData, _ := json.Marshal(meeting)
	req := httptest.NewRequest("POST", "/meeting", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
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
	mockMeetingService.On("UpdateMeeting", models.User{ID: 1}, &meeting).Return(nil)
	mp := presentations.NewMeetingPresentation(mockMeetingService)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.PUT("/meeting", addFakeUserMiddlewareMeeting(), mp.UpdateMeeting)

	jsonData, _ := json.Marshal(meeting)
	req := httptest.NewRequest("PUT", "/meeting", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestUpdateMeeting_InvalidRequest(t *testing.T) {
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
	r.PUT("/meeting", addFakeUserMiddlewareMeeting(), mp.UpdateMeeting)

	jsonData := []byte(`{"title": 123}`)
	req := httptest.NewRequest("PUT", "/meeting", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func TestUpdateMeeting_ErrorInService(t *testing.T) {
	meeting := models.Meeting{
		ID:          "1",
		Title:       "Updated Board Meeting",
		Description: "Updated Annual Board Meeting",
	}

	mockMeetingService := new(mockMeetingService)
	mockMeetingService.On("UpdateMeeting", models.User{ID: 1}, &meeting).Return(assert.AnError)
	mp := presentations.NewMeetingPresentation(mockMeetingService)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.PUT("/meeting", addFakeUserMiddlewareMeeting(), mp.UpdateMeeting)

	jsonData, _ := json.Marshal(meeting)
	req := httptest.NewRequest("PUT", "/meeting", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Code)
}

func TestDeleteMeeting(t *testing.T) {
	mockMeetingService := new(mockMeetingService)
	mockMeetingService.On("DeleteMeeting", models.User{ID: 1}, "1").Return(nil).Once()
	mp := presentations.NewMeetingPresentation(mockMeetingService)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.DELETE("/meeting/:id", addFakeUserMiddlewareMeeting(), mp.DeleteMeeting)

	t.Run("DeleteMeeting Normal", func(t *testing.T) {
		req := httptest.NewRequest("DELETE", "/meeting/1", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})

	t.Run("DeleteMeeting Invalid", func(t *testing.T) {
		r.DELETE("/meeting/", addFakeUserMiddlewareMeeting(), mp.DeleteMeeting) // temperary route to test invalid request
		req := httptest.NewRequest("DELETE", "/meeting/", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)
	})

	t.Run("DeleteMeeting ErrorInService", func(t *testing.T) {
		mockMeetingService.On("DeleteMeeting", models.User{ID: 1}, "1").Return(assert.AnError).Once()
		req := httptest.NewRequest("DELETE", "/meeting/1", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, 500, w.Code)
	})
}

func TestGetMeeting(t *testing.T) {
	meeting := models.Meeting{
		ID:          "1",
		Title:       "Board Meeting",
		Description: "Annual Board Meeting",
	}

	mockMeetingService := new(mockMeetingService)
	mockMeetingService.On("GetMeeting", "1").Return(meeting, nil).Once()
	mp := presentations.NewMeetingPresentation(mockMeetingService)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/meeting/:id", mp.GetMeeting)

	t.Run("GetMeeting Normal", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/meeting/1", nil)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})

	t.Run("GetMeeting Invalid", func(t *testing.T) {
		r.GET("/meeting/", mp.GetMeeting) // temperary route to test invalid request
		req := httptest.NewRequest("GET", "/meeting/", nil)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)
	})

	t.Run("GetMeeting ErrorInService", func(t *testing.T) {
		mockMeetingService.On("GetMeeting", "1").Return(models.Meeting{}, assert.AnError).Once()
		req := httptest.NewRequest("GET", "/meeting/1", nil)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, 500, w.Code)
	})
}

func TestGetAllMeetings(t *testing.T) {
	meetings := []models.Meeting{{ID: "1", Title: "Board Meeting", Description: "Annual Board Meeting"}}
	mockMeetingService := new(mockMeetingService)
	mockMeetingService.On("GetAllMeetings").Return(meetings, nil).Once()
	mp := presentations.NewMeetingPresentation(mockMeetingService)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/meetings", mp.GetAllMeetings)

	t.Run("GetAllMeetings Normal", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/meetings", nil)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})

	t.Run("GetAllMeetings ErrorInService", func(t *testing.T) {
		mockMeetingService.On("GetAllMeetings").Return([]models.Meeting{}, assert.AnError).Once()
		req := httptest.NewRequest("GET", "/meetings", nil)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, 500, w.Code)
	})
}

func TestGetMeetingsByRoomIdAndDate(t *testing.T) {
	meetings := []models.Meeting{{ID: "2", Title: "Team Meeting", Description: "Quarterly Planning"}}
	roomID := 101
	mockMeetingService := new(mockMeetingService)
	mockMeetingService.On("GetMeetingsByRoomIdAndDatePeriod", roomID, mock.AnythingOfType("time.Time"), mock.AnythingOfType("time.Time")).Return(meetings, nil).Once()
	// mockMeetingService.On("GetMeetingsByRoomIdAndDate", roomID, date).Return(meetings, nil)
	mp := presentations.NewMeetingPresentation(mockMeetingService)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/meetingsByRoomAndDate", mp.GetMeetingsByRoomIdAndDatePeriod)

	t.Run("GetMeetingsByRoomIdAndDate Normal", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/meetingsByRoomAndDate?room_id=101&date_from=2021-01-01&date_to=2021-01-01", nil)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})

	t.Run("GetMeetingsByRoomIdAndDate Invalid", func(t *testing.T) {
		r.GET("/meetingsByRoomAndDate/", mp.GetMeetingsByRoomIdAndDatePeriod) // temperary route to test invalid request
		req := httptest.NewRequest("GET", "/meetingsByRoomAndDate/", nil)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)
	})

	t.Run("GetMeetingsByRoomIdAndDate InvalidDate", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/meetingsByRoomAndDate?room_id=a&date_from=2021-01-01&date_to=a", nil)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)
	})

	t.Run("GetMeetingsByRoomIdAndDate NoDateTo", func(t *testing.T) {
		mockMeetingService.On("GetMeetingsByRoomIdAndDatePeriod", roomID, mock.AnythingOfType("time.Time"), mock.AnythingOfType("time.Time")).Return([]models.Meeting{}, nil).Once()
		req := httptest.NewRequest("GET", "/meetingsByRoomAndDate?room_id=101&date_from=2021-01-01", nil)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})

	t.Run("GetMeetingsByRoomIdAndDate empty response", func(t *testing.T) {
		mockMeetingService.On("GetMeetingsByRoomIdAndDatePeriod", roomID, mock.AnythingOfType("time.Time"), mock.AnythingOfType("time.Time")).Return([]models.Meeting{}, nil).Once()
		req := httptest.NewRequest("GET", "/meetingsByRoomAndDate?room_id=101&date_from=2021-01-01&date_to=2021-01-01", nil)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})

	t.Run("GetMeetingsByRoomIdAndDate ErrorInService", func(t *testing.T) {
		mockMeetingService.On("GetMeetingsByRoomIdAndDatePeriod", roomID, mock.AnythingOfType("time.Time"), mock.AnythingOfType("time.Time")).Return([]models.Meeting{}, assert.AnError).Once()
		req := httptest.NewRequest("GET", "/meetingsByRoomAndDate?room_id=101&date_from=2021-01-01&date_to=2021-01-01", nil)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, 500, w.Code)
	})

}

func TestGetMeetingsByParticipantId(t *testing.T) {
	meetings := []models.Meeting{{ID: "2", Title: "Team Meeting", Description: "Quarterly Planning"}}
	mockMeetingService := new(mockMeetingService)
	mockMeetingService.On("GetMeetingsByParticipantId", uint(1)).Return(meetings, nil).Once()
	mp := presentations.NewMeetingPresentation(mockMeetingService)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/meetingsByParticipantId", mp.GetMeetingsByParticipantId)

	t.Run("GetMeetingsByParticipantId Normal", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/meetingsByParticipantId?id=1", nil)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})

	t.Run("GetMeetingsByParticipantId Invalid", func(t *testing.T) {
		r.GET("/meetingsByParticipantId/", mp.GetMeetingsByParticipantId) // temperary route to test invalid request
		req := httptest.NewRequest("GET", "/meetingsByParticipantId/", nil)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)
	})

	t.Run("GetMeetingsByParticipantId Invalid (participantID Atoi failed)", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/meetingsByParticipantId?id=a", nil)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)
	})

	t.Run("GetMeetingsByParticipantId ErrorInService", func(t *testing.T) {
		mockMeetingService.On("GetMeetingsByParticipantId", uint(1)).Return([]models.Meeting{}, assert.AnError).Once()
		req := httptest.NewRequest("GET", "/meetingsByParticipantId?id=1", nil)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, 500, w.Code)
	})

	t.Run("GetMeetingsByParticipantId empty response", func(t *testing.T) {
		mockMeetingService.On("GetMeetingsByParticipantId", uint(1)).Return([]models.Meeting{}, nil).Once()
		req := httptest.NewRequest("GET", "/meetingsByParticipantId?id=1", nil)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})
}
