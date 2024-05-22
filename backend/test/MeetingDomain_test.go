package test

import (
	"errors"
	"meeting-center/src/domains"
	"meeting-center/src/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockMeetingRepo struct {
	mock.Mock
}

func (m *MockMeetingRepo) CreateMeeting(meeting *models.Meeting) (*models.Meeting, error) {
	args := m.Called(meeting)
	if args.Get(0) != nil {
		return args.Get(0).(*models.Meeting), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockMeetingRepo) UpdateMeeting(id string, meeting *models.Meeting) error {
	args := m.Called(id, meeting)
	return args.Error(0)
}

func (m *MockMeetingRepo) DeleteMeeting(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockMeetingRepo) GetMeeting(id string) (*models.Meeting, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*models.Meeting), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockMeetingRepo) GetAllMeetings() ([]*models.Meeting, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]*models.Meeting), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockMeetingRepo) GetMeetingsByRoomIdAndDate(roomID int, date time.Time) ([]*models.Meeting, error) {
	args := m.Called(roomID, date)
	if args.Get(0) != nil {
		return args.Get(0).([]*models.Meeting), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestDomainCreateMeeting(t *testing.T) {
	meeting := &models.Meeting{
		Title:       "Board Meeting",
		Description: "Annual Board Meeting",
	}

	mockMeetingRepo := new(MockMeetingRepo)
	mockMeetingRepo.On("CreateMeeting", meeting).Return(meeting, nil).Once()
	md := domains.NewMeetingDomain(mockMeetingRepo)
	createdMeeting, err := md.CreateMeeting(meeting)

	assert.NoError(t, err)
	assert.Equal(t, meeting.Title, createdMeeting.Title)
	assert.Equal(t, meeting.Description, createdMeeting.Description)

	// Test error case
	mockMeetingRepo.On("CreateMeeting", meeting).Return(nil, errors.New("create error")).Once()
	createdMeeting, err = md.CreateMeeting(meeting)
	assert.Error(t, err)
	assert.Nil(t, createdMeeting)
	assert.Contains(t, err.Error(), "create error")
}

func TestDomainUpdateMeeting(t *testing.T) {
	id := "1"
	meeting := &models.Meeting{
		Title:       "Updated Board Meeting",
		Description: "Updated Annual Board Meeting",
	}

	mockMeetingRepo := new(MockMeetingRepo)
	mockMeetingRepo.On("UpdateMeeting", id, meeting).Return(nil).Once()
	md := domains.NewMeetingDomain(mockMeetingRepo)
	err := md.UpdateMeeting(id, meeting)

	assert.NoError(t, err)

	// Test error case
	mockMeetingRepo.On("UpdateMeeting", id, meeting).Return(errors.New("update error")).Once()
	err = md.UpdateMeeting(id, meeting)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "update error")
}

func TestDomainDeleteMeeting(t *testing.T) {
	id := "1"

	mockMeetingRepo := new(MockMeetingRepo)
	mockMeetingRepo.On("DeleteMeeting", id).Return(nil).Once()
	md := domains.NewMeetingDomain(mockMeetingRepo)
	err := md.DeleteMeeting(id)

	assert.NoError(t, err)

	// Test error case
	mockMeetingRepo.On("DeleteMeeting", id).Return(errors.New("delete error")).Once()
	err = md.DeleteMeeting(id)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "delete error")
}

func TestDomainGetMeeting(t *testing.T) {
	id := "1"
	meeting := &models.Meeting{
		Title:       "Board Meeting",
		Description: "Annual Board Meeting",
	}

	mockMeetingRepo := new(MockMeetingRepo)
	mockMeetingRepo.On("GetMeeting", id).Return(meeting, nil).Once()
	md := domains.NewMeetingDomain(mockMeetingRepo)
	fetchedMeeting, err := md.GetMeeting(id)

	assert.NoError(t, err)
	assert.Equal(t, meeting.Title, fetchedMeeting.Title)
	assert.Equal(t, meeting.Description, fetchedMeeting.Description)

	// Test error case
	mockMeetingRepo.On("GetMeeting", id).Return(nil, errors.New("get error")).Once()
	fetchedMeeting, err = md.GetMeeting(id)
	assert.Error(t, err)
	assert.Nil(t, fetchedMeeting)
	assert.Contains(t, err.Error(), "get error")
}

func TestDomainGetAllMeetings(t *testing.T) {
	expectedMeetings := []*models.Meeting{{ID: "1", Title: "Board Meeting", Description: "Annual Board Meeting"}}

	mockMeetingRepo := new(MockMeetingRepo)
	mockMeetingRepo.On("GetAllMeetings").Return(expectedMeetings, nil).Once()
	md := domains.NewMeetingDomain(mockMeetingRepo)
	meetings, err := md.GetAllMeetings()

	assert.NoError(t, err)
	assert.Equal(t, expectedMeetings, meetings)

	// Test error case
	mockMeetingRepo.On("GetAllMeetings").Return(nil, errors.New("get all error")).Once()
	meetings, err = md.GetAllMeetings()
	assert.Error(t, err)
	assert.Nil(t, meetings)
	assert.Contains(t, err.Error(), "get all error")
}

func TestDomainGetMeetingsByRoomIdAndDate(t *testing.T) {
	roomID := 10
	date := time.Now()
	expectedMeetings := []*models.Meeting{{ID: "2", Title: "Strategy Meeting", Description: "Strategy planning session"}}

	mockMeetingRepo := new(MockMeetingRepo)
	mockMeetingRepo.On("GetMeetingsByRoomIdAndDate", roomID, date).Return(expectedMeetings, nil).Once()
	md := domains.NewMeetingDomain(mockMeetingRepo)
	meetings, err := md.GetMeetingsByRoomIdAndDate(roomID, date)

	assert.NoError(t, err)
	assert.Equal(t, expectedMeetings, meetings)

	// Test error case
	mockMeetingRepo.On("GetMeetingsByRoomIdAndDate", roomID, date).Return(nil, errors.New("get by room and date error")).Once()
	meetings, err = md.GetMeetingsByRoomIdAndDate(roomID, date)
	assert.Error(t, err)
	assert.Nil(t, meetings)
	assert.Contains(t, err.Error(), "get by room and date error")
}
