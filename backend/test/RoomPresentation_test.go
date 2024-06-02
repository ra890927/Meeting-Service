package test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"meeting-center/src/models"
	"meeting-center/src/presentations"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockRoomService struct {
	mock.Mock
}

func (m *mockRoomService) CreateRoom(room *models.Room) error {
	args := m.Called(room)
	return args.Error(0)
}

func (m *mockRoomService) UpdateRoom(room *models.Room) error {
	args := m.Called(room)
	return args.Error(0)
}

func (m *mockRoomService) DeleteRoom(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *mockRoomService) GetRoomByID(id int) (*models.Room, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*models.Room), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *mockRoomService) GetAllRooms() ([]*models.Room, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]*models.Room), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestNewRoomPresentation(t *testing.T) {
	mockService := new(mockRoomService)

	assert.NotPanics(t, func() {
		rp := presentations.NewRoomPresentation(mockService)
		assert.NotNil(t, rp)
	})

	assert.NotPanics(t, func() {
		rp := presentations.NewRoomPresentation()
		assert.NotNil(t, rp)
	})

	assert.Panics(t, func() {
		presentations.NewRoomPresentation(mockService, mockService)
	})
}

func TestCreateRoom(t *testing.T) {
	mockService := new(mockRoomService)
	rp := presentations.NewRoomPresentation(mockService)
	room := &models.Room{
		RoomName: "Conference Room A",
		Type:     "Board Meeting",
		Rules:    []int{1, 2, 3},
		Capacity: 10,
	}
	mockService.On("CreateRoom", room).Return(nil).Once()

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/room", rp.CreateRoom)

	// Test successful creation
	jsonData, _ := json.Marshal(room)
	req := httptest.NewRequest("POST", "/room", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	// Test invalid JSON
	invalidJSON := `{"room_name": "Conference Room A", "type": "Board Meeting", "rules": [1, 2, 3], "capacity": "invalid"}`
	req = httptest.NewRequest("POST", "/room", bytes.NewBufferString(invalidJSON))
	req.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)

	// Test service error
	mockService.On("CreateRoom", room).Return(errors.New("some error")).Once()
	req = httptest.NewRequest("POST", "/room", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Code)
}

func TestGetAllRooms(t *testing.T) {
	mockService := new(mockRoomService)
	rp := presentations.NewRoomPresentation(mockService)
	rooms := []*models.Room{
		{
			RoomName: "Conference Room A",
			Type:     "Board Meeting",
			Capacity: 10,
		},
		{
			RoomName: "Conference Room B",
			Type:     "Seminar",
			Capacity: 15,
		},
	}
	mockService.On("GetAllRooms").Return(rooms, nil).Once()

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/rooms", rp.GetAllRooms)

	req := httptest.NewRequest("GET", "/rooms", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, 2, len(rooms))

	// Test service error
	mockService.On("GetAllRooms").Return(nil, errors.New("db error")).Once()
	req = httptest.NewRequest("GET", "/rooms", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, 500, w.Code)
}

func TestUpdateRoom(t *testing.T) {
	mockService := new(mockRoomService)
	rp := presentations.NewRoomPresentation(mockService)
	room := &models.Room{
		ID:       1,
		RoomName: "Updated Room A",
		Type:     "Team Meeting",
		Rules:    []int{1, 2, 3},
		Capacity: 15,
	}
	mockService.On("UpdateRoom", room).Return(nil).Once()

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.PUT("/room/:id", rp.UpdateRoom)

	// Test successful update
	jsonData, _ := json.Marshal(room)
	req := httptest.NewRequest("PUT", "/room/1", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	// Test invalid JSON
	invalidJSON := `{"id": 1, "room_name": "Updated Room A", "type": "Team Meeting", "rules": [1, 2, 3], "capacity": "invalid"}`
	req = httptest.NewRequest("PUT", "/room/1", bytes.NewBufferString(invalidJSON))
	req.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)

	// Test service error
	mockService.On("UpdateRoom", room).Return(errors.New("some error")).Once()
	req = httptest.NewRequest("PUT", "/room/1", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Code)
}

func TestDeleteRoom(t *testing.T) {
	mockService := new(mockRoomService)
	rp := presentations.NewRoomPresentation(mockService)
	id := 1
	mockService.On("DeleteRoom", id).Return(nil).Once()

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.DELETE("/room/:id", rp.DeleteRoom)

	// Test successful deletion
	req := httptest.NewRequest("DELETE", "/room/"+fmt.Sprint(id), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	// Test invalid ID
	req = httptest.NewRequest("DELETE", "/room/invalid", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)

	// Test service error
	mockService.On("DeleteRoom", id).Return(errors.New("some error")).Once()
	req = httptest.NewRequest("DELETE", "/room/"+fmt.Sprint(id), nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, 500, w.Code)
}

func TestGetRoomByID(t *testing.T) {
	mockService := new(mockRoomService)
	rp := presentations.NewRoomPresentation(mockService)
	id := 1
	room := &models.Room{
		ID:       id,
		RoomName: "Conference Room A",
		Type:     "Board Meeting",
		Capacity: 10,
	}
	mockService.On("GetRoomByID", id).Return(room, nil).Once()

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/room/:id", rp.GetRoomByID)

	// Test successful get by ID
	req := httptest.NewRequest("GET", "/room/"+fmt.Sprint(id), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	// Test invalid ID
	req = httptest.NewRequest("GET", "/room/invalid", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)

	// Test service error
	mockService.On("GetRoomByID", id).Return(nil, errors.New("some error")).Once()
	req = httptest.NewRequest("GET", "/room/"+fmt.Sprint(id), nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, 500, w.Code)
}
