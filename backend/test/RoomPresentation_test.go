package test

import (
	"bytes"
	"encoding/json"
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
	return args.Get(0).(*models.Room), args.Error(1)
}

func (m *mockRoomService) GetAllRooms() ([]*models.Room, error) {
	args := m.Called()
	return args.Get(0).([]*models.Room), args.Error(1)
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
	mockService.On("CreateRoom", room).Return(nil)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/room", rp.CreateRoom)

	jsonData, _ := json.Marshal(room)
	req := httptest.NewRequest("POST", "/room", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
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
	mockService.On("GetAllRooms").Return(rooms, nil)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/rooms", rp.GetAllRooms)

	req := httptest.NewRequest("GET", "/rooms", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, 2, len(rooms))
}

func TestUpdateRoom(t *testing.T) {
	mockService := new(mockRoomService)
	rp := presentations.NewRoomPresentation(mockService)
	id := "1"
	room := &models.Room{
		RoomName: "Updated Room A",
		Type:     "Team Meeting",
		Capacity: 15,
	}
	mockService.On("UpdateRoom", room).Return(nil)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.PUT("/room/:id", rp.UpdateRoom)

	jsonData, _ := json.Marshal(room)
	req := httptest.NewRequest("PUT", "/room/"+id, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestDeleteRoom(t *testing.T) {
	mockService := new(mockRoomService)
	rp := presentations.NewRoomPresentation(mockService)
	id := 1
	mockService.On("DeleteRoom", id).Return(nil)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.DELETE("/room/:id", rp.DeleteRoom)

	req := httptest.NewRequest("DELETE", "/room/"+fmt.Sprint(id), nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestGetRoomByID(t *testing.T) {
	mockService := new(mockRoomService)
	rp := presentations.NewRoomPresentation(mockService)
	id := 1
	room := &models.Room{
		RoomName: "Conference Room A",
		Type:     "Board Meeting",
		Capacity: 10,
	}
	mockService.On("GetRoomByID", id).Return(room, nil)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/room/:id", rp.GetRoomByID)

	req := httptest.NewRequest("GET", "/room/"+fmt.Sprint(id), nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
