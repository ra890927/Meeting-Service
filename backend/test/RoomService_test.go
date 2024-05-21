package test

import (
	"meeting-center/src/models"
	"meeting-center/src/services"
	"testing"

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

func (m *MockRoomDomain) GetAllRooms() ([]*models.Room, error) {
	args := m.Called()
	return args.Get(0).([]*models.Room), args.Error(1)
}

func TestServiceCreateRoom(t *testing.T) {
	mockRoomDomain := new(MockRoomDomain)
	rs := services.NewRoomService(mockRoomDomain)

	room := &models.Room{
		RoomName: "Conference Room A",
		Type:     "Board Meeting",
		Capacity: 11,
	}

	mockRoomDomain.On("CreateRoom", room).Return(room, nil)

	createdRoom, err := rs.CreateRoom(room)

	assert.NoError(t, err)
	assert.Equal(t, room, createdRoom)
}

func TestServiceUpdateRoom(t *testing.T) {
	mockRoomDomain := new(MockRoomDomain)
	rs := services.NewRoomService(mockRoomDomain)

	id := "1"
	room := &models.Room{
		RoomName: "Updated Conference Room",
		Type:     "Executive Meeting",
		Capacity: 12,
	}

	mockRoomDomain.On("UpdateRoom", id, room).Return(nil)

	err := rs.UpdateRoom(id, room)

	assert.NoError(t, err)
}

func TestServiceDeleteRoom(t *testing.T) {
	mockRoomDomain := new(MockRoomDomain)
	rs := services.NewRoomService(mockRoomDomain)

	id := "1"

	mockRoomDomain.On("DeleteRoom", id).Return(nil)

	err := rs.DeleteRoom(id)

	assert.NoError(t, err)
}

func TestServiceGetRoom(t *testing.T) {
	mockRoomDomain := new(MockRoomDomain)
	rs := services.NewRoomService(mockRoomDomain)

	id := "1"
	room := &models.Room{
		RoomName: "Conference Room A",
		Type:     "Board Meeting",
		Capacity: 10,
	}

	mockRoomDomain.On("GetRoom", id).Return(room, nil)

	fetchedRoom, err := rs.GetRoom(id)

	assert.NoError(t, err)
	assert.Equal(t, room, fetchedRoom)
}

func TestServiceGetAllRooms(t *testing.T) {
	mockRoomDomain := new(MockRoomDomain)
	rs := services.NewRoomService(mockRoomDomain)

	rooms := []*models.Room{
		{
			RoomName: "Conference Room A",
			Type:     "Board Meeting",
			Capacity: 10,
		},
		{
			RoomName: "Conference Room B",
			Type:     "Seminar",
			Capacity: 20,
		},
	}

	mockRoomDomain.On("GetAllRooms").Return(rooms, nil)

	allRooms, err := rs.GetAllRooms()

	assert.NoError(t, err)
	assert.Equal(t, 2, len(allRooms))
	assert.Equal(t, rooms, allRooms)
}
