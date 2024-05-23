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

func (m *MockRoomRepo) CreateRoom(room *models.Room) error {
	args := m.Called(room)
	return args.Error(0)
}

func (m *MockRoomRepo) UpdateRoom(room *models.Room) error {
	args := m.Called(room)
	return args.Error(0)
}

func (m *MockRoomRepo) DeleteRoom(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockRoomRepo) GetRoomByID(id int) (*models.Room, error) {
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
		Capacity: 11,
	}

	mockRoomRepo.On("CreateRoom", room).Return(nil)
	err := rd.CreateRoom(room)

	assert.NoError(t, err)
}

func TestDomainUpdateRoom(t *testing.T) {
	mockRoomRepo := new(MockRoomRepo)
	rd := domains.NewRoomDomain(mockRoomRepo)
	room := &models.Room{
		RoomName: "Updated Conference Room",
		Type:     "Executive Meeting",
		Capacity: 12,
	}

	mockRoomRepo.On("UpdateRoom", room).Return(nil)
	err := rd.UpdateRoom(room)

	assert.NoError(t, err)
}

func TestDomainDeleteRoom(t *testing.T) {
	mockRoomRepo := new(MockRoomRepo)
	rd := domains.NewRoomDomain(mockRoomRepo)
	id := 1

	mockRoomRepo.On("DeleteRoom", id).Return(nil)
	err := rd.DeleteRoom(id)

	assert.NoError(t, err)
}

func TestDomainGetRoomByID(t *testing.T) {
	mockRoomRepo := new(MockRoomRepo)
	rd := domains.NewRoomDomain(mockRoomRepo)
	id := 1
	room := &models.Room{
		RoomName: "Conference Room A",
		Type:     "Board Meeting",
		Capacity: 10,
	}

	mockRoomRepo.On("GetRoomByID", id).Return(room, nil)
	fetchedRoom, err := rd.GetRoomByID(id)

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
