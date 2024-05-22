package test

import (
	"errors"
	"meeting-center/src/models"
	"meeting-center/src/services"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockMeetingDomain struct {
	mock.Mock
}

func (m *MockMeetingDomain) CreateMeeting(meeting *models.Meeting) (*models.Meeting, error) {
	args := m.Called(meeting)
	if args.Get(0) != nil {
		return args.Get(0).(*models.Meeting), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockMeetingDomain) UpdateMeeting(id string, meeting *models.Meeting) error {
	args := m.Called(id, meeting)
	return args.Error(0)
}

func (m *MockMeetingDomain) DeleteMeeting(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockMeetingDomain) GetMeeting(id string) (*models.Meeting, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*models.Meeting), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockMeetingDomain) GetAllMeetings() ([]*models.Meeting, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]*models.Meeting), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockMeetingDomain) GetMeetingsByRoomIdAndDate(roomID int, date time.Time) ([]*models.Meeting, error) {
	args := m.Called(roomID, date)
	if args.Get(0) != nil {
		return args.Get(0).([]*models.Meeting), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestServiceCreateMeetingWithInvalidTime(t *testing.T) {
	mockMeetingDomain := new(MockMeetingDomain)
	ms := services.NewMeetingService(mockMeetingDomain)

	meeting := &models.Meeting{
		Title:       "Invalid Time Meeting",
		Description: "This meeting has invalid time",
		RoomID:      101,
		StartTime:   time.Now().Add(2 * time.Hour),
		EndTime:     time.Now(),
	}

	createdMeeting, err := ms.CreateMeeting(meeting)

	assert.Error(t, err)
	assert.Nil(t, createdMeeting)
	assert.Contains(t, err.Error(), "StartTime must be before EndTime")
}

func TestServiceCreateMeetingWithNoConflict(t *testing.T) {
	mockMeetingDomain := new(MockMeetingDomain)
	ms := services.NewMeetingService(mockMeetingDomain)

	meeting := &models.Meeting{
		ID:          "1",
		Title:       "Board Meeting",
		Description: "Annual Board Meeting",
		RoomID:      101,
		StartTime:   time.Now(),
		EndTime:     time.Now().Add(2 * time.Hour),
	}

	// Assume no existing meetings
	mockMeetingDomain.On("GetMeetingsByRoomIdAndDate", meeting.RoomID, meeting.StartTime).Return([]*models.Meeting{}, nil)
	mockMeetingDomain.On("CreateMeeting", meeting).Return(meeting, nil)

	createdMeeting, err := ms.CreateMeeting(meeting)

	assert.NoError(t, err)
	assert.Equal(t, meeting, createdMeeting)
}

func TestServiceCreateMeetingWithConflict(t *testing.T) {
	mockMeetingDomain := new(MockMeetingDomain)
	ms := services.NewMeetingService(mockMeetingDomain)

	meeting := &models.Meeting{
		ID:          "1",
		Title:       "Board Meeting",
		Description: "Annual Board Meeting",
		RoomID:      101,
		StartTime:   time.Now(),
		EndTime:     time.Now().Add(2 * time.Hour),
	}

	existingMeeting := &models.Meeting{
		ID:        "2",
		Title:     "Strategy Meeting",
		RoomID:    101,
		StartTime: time.Now().Add(1 * time.Hour),
		EndTime:   time.Now().Add(3 * time.Hour),
	}

	// Return existing meeting that conflicts
	mockMeetingDomain.On("GetMeetingsByRoomIdAndDate", meeting.RoomID, meeting.StartTime).Return([]*models.Meeting{existingMeeting}, nil)

	createdMeeting, err := ms.CreateMeeting(meeting)

	assert.Error(t, err)
	assert.Nil(t, createdMeeting)
	assert.Contains(t, err.Error(), "time slot is already booked")
}

func TestServiceCreateMeetingWithGetMeetingsError(t *testing.T) {
	mockMeetingDomain := new(MockMeetingDomain)
	ms := services.NewMeetingService(mockMeetingDomain)

	meeting := &models.Meeting{
		ID:          "1",
		Title:       "Board Meeting",
		Description: "Annual Board Meeting",
		RoomID:      101,
		StartTime:   time.Now(),
		EndTime:     time.Now().Add(2 * time.Hour),
	}

	mockMeetingDomain.On("GetMeetingsByRoomIdAndDate", meeting.RoomID, meeting.StartTime).Return(nil, errors.New("get meetings error"))

	createdMeeting, err := ms.CreateMeeting(meeting)

	assert.Error(t, err)
	assert.Nil(t, createdMeeting)
	assert.Contains(t, err.Error(), "get meetings error")
}

func TestServiceCreateMeetingWithCreateMeetingError(t *testing.T) {
	mockMeetingDomain := new(MockMeetingDomain)
	ms := services.NewMeetingService(mockMeetingDomain)

	meeting := &models.Meeting{
		ID:          "1",
		Title:       "Board Meeting",
		Description: "Annual Board Meeting",
		RoomID:      101,
		StartTime:   time.Now(),
		EndTime:     time.Now().Add(2 * time.Hour),
	}

	mockMeetingDomain.On("GetMeetingsByRoomIdAndDate", meeting.RoomID, meeting.StartTime).Return([]*models.Meeting{}, nil)
	mockMeetingDomain.On("CreateMeeting", meeting).Return(nil, errors.New("create meeting error"))

	createdMeeting, err := ms.CreateMeeting(meeting)

	assert.Error(t, err)
	assert.Nil(t, createdMeeting)
	assert.Contains(t, err.Error(), "create meeting error")
}

func TestServiceUpdateMeeting(t *testing.T) {
	mockMeetingDomain := new(MockMeetingDomain)
	ms := services.NewMeetingService(mockMeetingDomain)

	id := "1"
	meeting := &models.Meeting{
		Title:       "Updated Board Meeting",
		Description: "Updated Annual Board Meeting",
	}

	mockMeetingDomain.On("UpdateMeeting", id, meeting).Return(nil)

	err := ms.UpdateMeeting(id, meeting)

	assert.NoError(t, err)
}

func TestServiceUpdateMeetingWithError(t *testing.T) {
	mockMeetingDomain := new(MockMeetingDomain)
	ms := services.NewMeetingService(mockMeetingDomain)

	id := "1"
	meeting := &models.Meeting{
		Title:       "Updated Board Meeting",
		Description: "Updated Annual Board Meeting",
	}

	mockMeetingDomain.On("UpdateMeeting", id, meeting).Return(errors.New("update error"))

	err := ms.UpdateMeeting(id, meeting)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "update error")
}

func TestServiceDeleteMeeting(t *testing.T) {
	mockMeetingDomain := new(MockMeetingDomain)
	ms := services.NewMeetingService(mockMeetingDomain)

	id := "1"

	mockMeetingDomain.On("DeleteMeeting", id).Return(nil)

	err := ms.DeleteMeeting(id)

	assert.NoError(t, err)
}

func TestServiceDeleteMeetingWithError(t *testing.T) {
	mockMeetingDomain := new(MockMeetingDomain)
	ms := services.NewMeetingService(mockMeetingDomain)

	id := "1"

	mockMeetingDomain.On("DeleteMeeting", id).Return(errors.New("delete error"))

	err := ms.DeleteMeeting(id)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "delete error")
}

func TestServiceGetMeeting(t *testing.T) {
	mockMeetingDomain := new(MockMeetingDomain)
	ms := services.NewMeetingService(mockMeetingDomain)

	id := "1"
	meeting := &models.Meeting{
		Title:       "Board Meeting",
		Description: "Annual Board Meeting",
	}

	mockMeetingDomain.On("GetMeeting", id).Return(meeting, nil)

	fetchedMeeting, err := ms.GetMeeting(id)

	assert.NoError(t, err)
	assert.Equal(t, meeting, fetchedMeeting)
}

func TestServiceGetMeetingWithError(t *testing.T) {
	mockMeetingDomain := new(MockMeetingDomain)
	ms := services.NewMeetingService(mockMeetingDomain)

	id := "1"

	mockMeetingDomain.On("GetMeeting", id).Return(nil, errors.New("get error"))

	fetchedMeeting, err := ms.GetMeeting(id)

	assert.Error(t, err)
	assert.Nil(t, fetchedMeeting)
	assert.Contains(t, err.Error(), "get error")
}

func TestServiceGetAllMeetings(t *testing.T) {
	mockMeetingDomain := new(MockMeetingDomain)
	ms := services.NewMeetingService(mockMeetingDomain)
	expectedMeetings := []*models.Meeting{{ID: "1", Title: "Board Meeting", Description: "Annual Board Meeting"}}

	mockMeetingDomain.On("GetAllMeetings").Return(expectedMeetings, nil)

	meetings, err := ms.GetAllMeetings()

	assert.NoError(t, err)
	assert.Equal(t, expectedMeetings, meetings)
}

func TestServiceGetAllMeetingsWithError(t *testing.T) {
	mockMeetingDomain := new(MockMeetingDomain)
	ms := services.NewMeetingService(mockMeetingDomain)

	mockMeetingDomain.On("GetAllMeetings").Return(nil, errors.New("get all error"))

	meetings, err := ms.GetAllMeetings()

	assert.Error(t, err)
	assert.Nil(t, meetings)
	assert.Contains(t, err.Error(), "get all error")
}

func TestServiceGetMeetingsByRoomIdAndDate(t *testing.T) {
	mockMeetingDomain := new(MockMeetingDomain)
	ms := services.NewMeetingService(mockMeetingDomain)
	roomID := 10
	date := time.Now()
	expectedMeetings := []*models.Meeting{{ID: "2", Title: "Strategy Meeting", Description: "Strategy planning session"}}

	mockMeetingDomain.On("GetMeetingsByRoomIdAndDate", roomID, date).Return(expectedMeetings, nil)

	meetings, err := ms.GetMeetingsByRoomIdAndDate(roomID, date)

	assert.NoError(t, err)
	assert.Equal(t, expectedMeetings, meetings)
}

func TestServiceGetMeetingsByRoomIdAndDateWithError(t *testing.T) {
	mockMeetingDomain := new(MockMeetingDomain)
	ms := services.NewMeetingService(mockMeetingDomain)
	roomID := 10
	date := time.Now()

	mockMeetingDomain.On("GetMeetingsByRoomIdAndDate", roomID, date).Return(nil, errors.New("get by room and date error"))

	meetings, err := ms.GetMeetingsByRoomIdAndDate(roomID, date)

	assert.Error(t, err)
	assert.Nil(t, meetings)
	assert.Contains(t, err.Error(), "get by room and date error")
}
