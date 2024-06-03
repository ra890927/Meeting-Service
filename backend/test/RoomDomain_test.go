package test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"meeting-center/src/domains"
	"meeting-center/src/models"
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
	if args.Get(0) != nil {
		return args.Get(0).(*models.Room), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockRoomRepo) GetAllRooms() ([]*models.Room, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]*models.Room), args.Error(1)
	}
	return nil, args.Error(1)
}

type RoomDomainTestSuite struct {
	suite.Suite
	rd domains.RoomDomain
	rr *MockRoomRepo
}

func (suite *RoomDomainTestSuite) SetupTest() {
	suite.rr = new(MockRoomRepo)
	suite.rd = domains.NewRoomDomain(suite.rr)
}

func TestRoomDomainTestSuite(t *testing.T) {
	suite.Run(t, new(RoomDomainTestSuite))
}

func (suite *RoomDomainTestSuite) TestNewRoomDomain() {
	// Test case with no arguments
	// rd := domains.NewRoomDomain()
	// assert.NotNil(suite.T(), rd)

	// Test case with one argument
	mockRepo := new(MockRoomRepo)
	rd := domains.NewRoomDomain(mockRepo)
	assert.NotNil(suite.T(), rd)

	// Test case with too many arguments should panic
	assert.Panics(suite.T(), func() {
		domains.NewRoomDomain(mockRepo, mockRepo)
	})
}

func (suite *RoomDomainTestSuite) TestCreateRoom() {
	room := &models.Room{
		RoomName: "Conference Room A",
		Type:     "Board Meeting",
		Capacity: 11,
	}

	// Test successful creation
	suite.rr.On("CreateRoom", room).Return(nil).Once()
	err := suite.rd.CreateRoom(room)
	assert.NoError(suite.T(), err)
	suite.rr.AssertCalled(suite.T(), "CreateRoom", room)

	// Test creation error
	suite.rr.On("CreateRoom", room).Return(errors.New("create error")).Once()
	err = suite.rd.CreateRoom(room)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "create error", err.Error())
	suite.rr.AssertCalled(suite.T(), "CreateRoom", room)
}

func (suite *RoomDomainTestSuite) TestUpdateRoom() {
	room := &models.Room{
		RoomName: "Updated Conference Room",
		Type:     "Executive Meeting",
		Capacity: 12,
	}

	// Test successful update
	suite.rr.On("UpdateRoom", room).Return(nil).Once()
	err := suite.rd.UpdateRoom(room)
	assert.NoError(suite.T(), err)
	suite.rr.AssertCalled(suite.T(), "UpdateRoom", room)

	// Test update error
	suite.rr.On("UpdateRoom", room).Return(errors.New("update error")).Once()
	err = suite.rd.UpdateRoom(room)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "update error", err.Error())
	suite.rr.AssertCalled(suite.T(), "UpdateRoom", room)
}

func (suite *RoomDomainTestSuite) TestDeleteRoom() {
	id := 1

	// Test successful deletion
	suite.rr.On("DeleteRoom", id).Return(nil).Once()
	err := suite.rd.DeleteRoom(id)
	assert.NoError(suite.T(), err)
	suite.rr.AssertCalled(suite.T(), "DeleteRoom", id)

	// Test deletion error
	suite.rr.On("DeleteRoom", id).Return(errors.New("delete error")).Once()
	err = suite.rd.DeleteRoom(id)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "delete error", err.Error())
	suite.rr.AssertCalled(suite.T(), "DeleteRoom", id)
}

func (suite *RoomDomainTestSuite) TestGetRoomByID() {
	id := 1
	room := &models.Room{
		RoomName: "Conference Room A",
		Type:     "Board Meeting",
		Capacity: 10,
	}

	// Test successful retrieval
	suite.rr.On("GetRoomByID", id).Return(room, nil).Once()
	fetchedRoom, err := suite.rd.GetRoomByID(id)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), room, fetchedRoom)
	suite.rr.AssertCalled(suite.T(), "GetRoomByID", id)

	// Test retrieval error
	suite.rr.On("GetRoomByID", id).Return(nil, errors.New("not found")).Once()
	fetchedRoom, err = suite.rd.GetRoomByID(id)
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), fetchedRoom)
	assert.Equal(suite.T(), "not found", err.Error())
	suite.rr.AssertCalled(suite.T(), "GetRoomByID", id)
}

func (suite *RoomDomainTestSuite) TestGetAllRooms() {
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

	// Test successful retrieval
	suite.rr.On("GetAllRooms").Return(rooms, nil).Once()
	allRooms, err := suite.rd.GetAllRooms()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), rooms, allRooms)
	suite.rr.AssertCalled(suite.T(), "GetAllRooms")

	// Test retrieval error
	suite.rr.On("GetAllRooms").Return(nil, errors.New("db error")).Once()
	allRooms, err = suite.rd.GetAllRooms()
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), allRooms)
	assert.Equal(suite.T(), "db error", err.Error())
	suite.rr.AssertCalled(suite.T(), "GetAllRooms")

	// Test no rooms
	suite.rr.On("GetAllRooms").Return(nil, nil).Once()
	allRooms, err = suite.rd.GetAllRooms()
	assert.NoError(suite.T(), err)
	assert.Empty(suite.T(), allRooms)
	suite.rr.AssertCalled(suite.T(), "GetAllRooms")
}
