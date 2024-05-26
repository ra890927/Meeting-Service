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

func (m *MockRoomDomain) CreateRoom(room *models.Room) error {
	args := m.Called(room)
	return args.Error(0)
}

func (m *MockRoomDomain) UpdateRoom(room *models.Room) error {
	args := m.Called(room)
	return args.Error(0)
}

func (m *MockRoomDomain) DeleteRoom(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockRoomDomain) GetRoomByID(id int) (*models.Room, error) {
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

	mockRoomDomain.On("CreateRoom", room).Return(nil)

	err := rs.CreateRoom(room)

	assert.NoError(t, err)
}

func TestServiceUpdateRoom(t *testing.T) {
	mockRoomDomain := new(MockRoomDomain)
	rs := services.NewRoomService(mockRoomDomain)
	room := &models.Room{
		ID:       1,
		RoomName: "Updated Conference Room",
		Type:     "Executive Meeting",
		Capacity: 12,
	}

	mockRoomDomain.On("UpdateRoom", room).Return(nil)

	err := rs.UpdateRoom(room)

	assert.NoError(t, err)
}

func TestServiceDeleteRoom(t *testing.T) {
	mockRoomDomain := new(MockRoomDomain)
	rs := services.NewRoomService(mockRoomDomain)

	id := 1

	mockRoomDomain.On("DeleteRoom", id).Return(nil)

	err := rs.DeleteRoom(id)

	assert.NoError(t, err)
}

func TestServiceGetRoomByID(t *testing.T) {
	mockRoomDomain := new(MockRoomDomain)
	rs := services.NewRoomService(mockRoomDomain)

	id := 1
	room := &models.Room{
		RoomName: "Conference Room A",
		Type:     "Board Meeting",
		Capacity: 10,
	}

	mockRoomDomain.On("GetRoomByID", id).Return(room, nil)

	fetchedRoom, err := rs.GetRoomByID(id)

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
