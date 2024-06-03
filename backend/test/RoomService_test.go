package test

import (
	"errors"
	"meeting-center/src/models"
	"meeting-center/src/services"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
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
	if args.Get(0) != nil {
		return args.Get(0).(*models.Room), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockRoomDomain) GetAllRooms() ([]*models.Room, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]*models.Room), args.Error(1)
	}
	return nil, args.Error(1)
}

type RoomServiceTestSuite struct {
	suite.Suite
	rs services.RoomService
	rd *MockRoomDomain
}

func (suite *RoomServiceTestSuite) SetupTest() {
	suite.rd = new(MockRoomDomain)
	suite.rs = services.NewRoomService(suite.rd)
}

func TestRoomServiceTestSuite(t *testing.T) {
	suite.Run(t, new(RoomServiceTestSuite))
}

func (suite *RoomServiceTestSuite) TestNewRoomService() {
	// Arrange
	mockDomain := new(MockRoomDomain)

	// Act and Assert

	// Test case with one argument
	rs := services.NewRoomService(mockDomain)
	assert.NotNil(suite.T(), rs)

	// Test case with no arguments
	// rs = services.NewRoomService()
	// assert.NotNil(suite.T(), rs)

	// Test case with multiple arguments should panic
	assert.Panics(suite.T(), func() {
		services.NewRoomService(mockDomain, mockDomain)
	})
}

func (suite *RoomServiceTestSuite) TestCreateRoom() {
	room := &models.Room{
		RoomName: "Conference Room A",
		Type:     "Board Meeting",
		Capacity: 11,
	}
	suite.rd.On("CreateRoom", room).Return(nil).Once()

	err := suite.rs.CreateRoom(room)

	assert.NoError(suite.T(), err)

	// Test error case
	suite.rd.On("CreateRoom", room).Return(errors.New("some error")).Once()

	err = suite.rs.CreateRoom(room)

	assert.Error(suite.T(), err)
}

func (suite *RoomServiceTestSuite) TestUpdateRoom() {
	room := &models.Room{
		ID:       1,
		RoomName: "Updated Conference Room",
		Type:     "Executive Meeting",
		Capacity: 12,
	}
	suite.rd.On("UpdateRoom", room).Return(nil).Once()

	err := suite.rs.UpdateRoom(room)

	assert.NoError(suite.T(), err)

	// Test error case
	suite.rd.On("UpdateRoom", room).Return(errors.New("some error")).Once()

	err = suite.rs.UpdateRoom(room)

	assert.Error(suite.T(), err)
}

func (suite *RoomServiceTestSuite) TestDeleteRoom() {
	id := 1
	suite.rd.On("DeleteRoom", id).Return(nil).Once()

	err := suite.rs.DeleteRoom(id)

	assert.NoError(suite.T(), err)

	// Test error case
	suite.rd.On("DeleteRoom", id).Return(errors.New("some error")).Once()

	err = suite.rs.DeleteRoom(id)

	assert.Error(suite.T(), err)
}

func (suite *RoomServiceTestSuite) TestGetRoomByID() {
	id := 1
	room := &models.Room{
		RoomName: "Conference Room A",
		Type:     "Board Meeting",
		Capacity: 10,
	}

	suite.rd.On("GetRoomByID", id).Return(room, nil).Once()

	fetchedRoom, err := suite.rs.GetRoomByID(id)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), room, fetchedRoom)

	// Test error case
	suite.rd.On("GetRoomByID", id).Return(nil, errors.New("some error")).Once()

	fetchedRoom, err = suite.rs.GetRoomByID(id)

	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), fetchedRoom)
}

func (suite *RoomServiceTestSuite) TestGetAllRooms() {
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
	suite.rd.On("GetAllRooms").Return(rooms, nil).Once()

	allRooms, err := suite.rs.GetAllRooms()

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), rooms, allRooms)

	// Test error case
	suite.rd.On("GetAllRooms").Return(nil, errors.New("some error")).Once()

	allRooms, err = suite.rs.GetAllRooms()

	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), allRooms)
}
