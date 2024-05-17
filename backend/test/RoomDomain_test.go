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

func TestDomainCreateRoom(t *testing.T) {
	room := &models.Room{
		RoomName: "Conference Room A",
		Type:     "Board Meeting",
		Capacity: 10,
	}
	mockRoomRepo := new(MockRoomRepo)
	mockRoomRepo.On("CreateRoom", room).Return(room, nil)
	rd := domains.NewRoomDomain(mockRoomRepo)
	createdRoom, err := rd.CreateRoom(room)

	assert.NoError(t, err)
	assert.Equal(t, room, createdRoom)
}

func TestDomainUpdateRoom(t *testing.T) {
	id := "1"
	room := &models.Room{
		RoomName: "Updated Conference Room",
		Type:     "Executive Meeting",
		Capacity: 12,
	}
	mockRoomRepo := new(MockRoomRepo)
	mockRoomRepo.On("UpdateRoom", id, room).Return(nil)
	rd := domains.NewRoomDomain(mockRoomRepo)
	err := rd.UpdateRoom(id, room)

	assert.NoError(t, err)
}

func TestDomainDeleteRoom(t *testing.T) {
	id := "1"
	mockRoomRepo := new(MockRoomRepo)
	mockRoomRepo.On("DeleteRoom", id).Return(nil)
	rd := domains.NewRoomDomain(mockRoomRepo)
	err := rd.DeleteRoom(id)

	assert.NoError(t, err)
}

func TestDomainGetRoom(t *testing.T) {
	id := "1"
	room := &models.Room{
		RoomName: "Conference Room A",
		Type:     "Board Meeting",
		Capacity: 10,
	}
	mockRoomRepo := new(MockRoomRepo)
	mockRoomRepo.On("GetRoom", id).Return(room, nil)
	rd := domains.NewRoomDomain(mockRoomRepo)
	fetchedRoom, err := rd.GetRoom(id)

	assert.NoError(t, err)
	assert.Equal(t, room, fetchedRoom)
}
