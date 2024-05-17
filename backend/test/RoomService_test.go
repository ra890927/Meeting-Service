package services

import (
	"meeting-center/src/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRoomDomain struct {
	mock.Mock
}

func (m *MockRoomDomain) CreateRoom(room *models.Room) (*models.Room, error) {
	args := m.Called(room)
	return args.Get(0).(*models.Room), args.Error(1)
}

func (m *MockRoomDomain) UpdateRoom(id string, room *models.Room) error {
	args := m.Called(id, room)
	return args.Error(0)
}

func (m *MockRoomDomain) DeleteRoom(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockRoomDomain) GetRoom(id string) (*models.Room, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Room), args.Error(1)
}

func TestCreateRoom(t *testing.T) {
	mockRoomDomain := new(MockRoomDomain)
	roomService := NewRoomService(mockRoomDomain)
	room := &models.Room{
		RoomName:  "Conference Room A",
		Type:      "Board Meeting",
		Rules:     []string{"No food allowed", "Respect the start time"},
		Capacity:  10,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	mockRoomDomain.On("CreateRoom", room).Return(room, nil)

	createdRoom, err := roomService.CreateRoom(room)

	assert.NoError(t, err)
	assert.Equal(t, room, createdRoom)
}

func TestUpdateRoom(t *testing.T) {
	mockRoomDomain := new(MockRoomDomain)
	roomService := NewRoomService(mockRoomDomain)
	id := "1"
	room := &models.Room{
		RoomName: "Updated Room A",
		Type:     "Executive Meeting",
		Capacity: 12,
	}
	mockRoomDomain.On("UpdateRoom", id, room).Return(nil)

	err := roomService.UpdateRoom(id, room)

	assert.NoError(t, err)
}

func TestDeleteRoom(t *testing.T) {
	mockRoomDomain := new(MockRoomDomain)
	roomService := NewRoomService(mockRoomDomain)
	id := "1"
	mockRoomDomain.On("DeleteRoom", id).Return(nil)

	err := roomService.DeleteRoom(id)

	assert.NoError(t, err)
}

func TestGetRoom(t *testing.T) {
	mockRoomDomain := new(MockRoomDomain)
	roomService := NewRoomService(mockRoomDomain)
	id := "1"
	room := &models.Room{
		RoomName: "Conference Room A",
		Type:     "Board Meeting",
		Capacity: 10,
	}
	mockRoomDomain.On("GetRoom", id).Return(room, nil)

	fetchedRoom, err := roomService.GetRoom(id)

	assert.NoError(t, err)
	assert.Equal(t, room, fetchedRoom)
}
