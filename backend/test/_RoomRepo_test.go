package test

import (
	"meeting-center/src/models"
	"meeting-center/src/repos"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockGormDB simulates the GormDB for testing
type MockGormDB struct {
	mock.Mock
}

func (db *MockGormDB) Create(value interface{}) *MockGormDB {
	args := db.Called(value)
	return args.Get(0).(*MockGormDB)
}

func (db *MockGormDB) Update(id string, value interface{}) *MockGormDB {
	args := db.Called(id, value)
	return args.Get(0).(*MockGormDB)
}

func (db *MockGormDB) Delete(value interface{}, id string) *MockGormDB {
	args := db.Called(value, id)
	return args.Get(0).(*MockGormDB)
}

func (db *MockGormDB) First(out interface{}, where ...interface{}) *MockGormDB {
	args := db.Called(out, where)
	return args.Get(0).(*MockGormDB)
}

func TestRoomRepo_CreateRoom(t *testing.T) {
	mockDB := new(MockGormDB)
	roomRepo := repos.RoomRepo{DB: mockDB}

	room := &models.Room{
		RoomName: "Test Room",
		Type:     "Conference",
		Capacity: 10,
	}

	mockDB.On("Create", room).Return(mockDB)
	mockDB.On("Error").Return(nil)

	createdRoom, err := roomRepo.CreateRoom(room)

	assert.NoError(t, err)
	assert.Equal(t, room.RoomName, createdRoom.RoomName)
}

func TestRoomRepo_UpdateRoom(t *testing.T) {
	mockDB := new(MockGormDB)
	roomRepo := repos.RoomRepo{DB: mockDB}

	room := &models.Room{
		RoomName: "Updated Room",
		Type:     "Seminar",
		Capacity: 20,
	}
	id := "1"

	mockDB.On("Model", &models.Room{}).Return(mockDB)
	mockDB.On("Where", "id = ?", id).Return(mockDB)
	mockDB.On("Updates", room).Return(mockDB)
	mockDB.On("Error").Return(nil)

	err := roomRepo.UpdateRoom(id, room)

	assert.NoError(t, err)
}

func TestRoomRepo_DeleteRoom(t *testing.T) {
	mockDB := new(MockGormDB)
	roomRepo := repos.RoomRepo{DB: mockDB}
	id := "1"

	mockDB.On("Delete", &models.Room{}, id).Return(mockDB)
	mockDB.On("Error").Return(nil)

	err := roomRepo.DeleteRoom(id)

	assert.NoError(t, err)
}

func TestRoomRepo_GetRoom(t *testing.T) {
	mockDB := new(MockGormDB)
	roomRepo := repos.RoomRepo{DB: mockDB}
	id := "1"
	room := &models.Room{}

	mockDB.On("First", room, "id = ?", id).Return(mockDB)
	mockDB.On("Error").Return(nil)

	fetchedRoom, err := roomRepo.GetRoom(id)

	assert.NoError(t, err)
	assert.Equal(t, room, fetchedRoom)
}
