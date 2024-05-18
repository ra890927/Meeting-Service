package test

import (
	"meeting-center/src/domains"
	"meeting-center/src/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRoomRepo struct {
	mock.Mock
}

func (m *MockRoomRepo) CreateRoom(room *models.Room) (*models.Room, error) {
	args := m.Called(room)
	return args.Get(0).(*models.Room), args.Error(1)
}

func (m *MockRoomRepo) UpdateRoom(id string, room *models.Room) error {
	args := m.Called(id, room)
	return args.Error(0)
}

func (m *MockRoomRepo) DeleteRoom(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockRoomRepo) GetRoom(id string) (*models.Room, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Room), args.Error(1)
}

func (m *MockRoomRepo) GetAllRooms() ([]*models.Room, error) {
	args := m.Called()
	return args.Get(0).([]*models.Room), args.Error(1)
}

func TestDomainCreateRoom(t *testing.T) {
	mockRoomRepo := new(MockRoomRepo)
	rd := domains.NewRoomDomain(mockRoomRepo)
	room := &models.Room{
		RoomName: "Conference Room A",
		Type:     "Board Meeting",
		Capacity: 10,
	}

	mockRoomRepo.On("CreateRoom", room).Return(room, nil)
	createdRoom, err := rd.CreateRoom(room)

	assert.NoError(t, err)
	assert.Equal(t, room, createdRoom)
}

func TestDomainUpdateRoom(t *testing.T) {
	mockRoomRepo := new(MockRoomRepo)
	rd := domains.NewRoomDomain(mockRoomRepo)
	id := "1"
	room := &models.Room{
		RoomName: "Updated Conference Room",
		Type:     "Executive Meeting",
		Capacity: 12,
	}

	mockRoomRepo.On("UpdateRoom", id, room).Return(nil)
	err := rd.UpdateRoom(id, room)

	assert.NoError(t, err)
}

func TestDomainDeleteRoom(t *testing.T) {
	mockRoomRepo := new(MockRoomRepo)
	rd := domains.NewRoomDomain(mockRoomRepo)
	id := "1"

	mockRoomRepo.On("DeleteRoom", id).Return(nil)
	err := rd.DeleteRoom(id)

	assert.NoError(t, err)
}

func TestDomainGetRoom(t *testing.T) {
	mockRoomRepo := new(MockRoomRepo)
	rd := domains.NewRoomDomain(mockRoomRepo)
	id := "1"
	room := &models.Room{
		RoomName: "Conference Room A",
		Type:     "Board Meeting",
		Capacity: 10,
	}

	mockRoomRepo.On("GetRoom", id).Return(room, nil)
	fetchedRoom, err := rd.GetRoom(id)

	assert.NoError(t, err)
	assert.Equal(t, room, fetchedRoom)
}

func TestDomainGetAllRooms(t *testing.T) {
	mockRoomRepo := new(MockRoomRepo)
	rd := domains.NewRoomDomain(mockRoomRepo)
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

	mockRoomRepo.On("GetAllRooms").Return(rooms, nil)
	allRooms, err := rd.GetAllRooms()

	assert.NoError(t, err)
	assert.Equal(t, 2, len(allRooms))
	assert.Equal(t, rooms, allRooms)
}
