package test

import (
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

func (m *MockMeetingRepo) CreateMeeting(meeting *models.Meeting) error {
	args := m.Called(meeting)
	return args.Error(0)
}

func (m *MockMeetingRepo) UpdateMeeting(meeting *models.Meeting) error {
	args := m.Called(meeting)
	return args.Error(0)
}

func (m *MockMeetingRepo) DeleteMeeting(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockMeetingRepo) GetMeeting(id string) (*models.Meeting, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Meeting), args.Error(1)
}

func (m *MockMeetingRepo) GetAllMeetings() ([]*models.Meeting, error) {
	args := m.Called()
	return args.Get(0).([]*models.Meeting), args.Error(1)
}

func (m *MockMeetingRepo) GetMeetingsByRoomIdAndDate(roomID int, date time.Time) ([]*models.Meeting, error) {
	args := m.Called(roomID, date)
	return args.Get(0).([]*models.Meeting), args.Error(1)
}

func TestDomainCreateMeeting(t *testing.T) {
	meeting := &models.Meeting{
		Title:       "Board Meeting",
		Description: "Annual Board Meeting",
	}

	mockMeetingRepo := new(MockMeetingRepo)
	mockMeetingRepo.On("CreateMeeting", meeting).Return(nil)
	md := domains.NewMeetingDomain(mockMeetingRepo)
	err := md.CreateMeeting(meeting)

	assert.NoError(t, err)
}

func TestDomainUpdateMeeting(t *testing.T) {
	meeting := &models.Meeting{
		ID:          "1",
		Title:       "Updated Board Meeting",
		Description: "Updated Annual Board Meeting",
	}

	mockMeetingRepo := new(MockMeetingRepo)
	mockMeetingRepo.On("UpdateMeeting", meeting).Return(nil)
	md := domains.NewMeetingDomain(mockMeetingRepo)
	err := md.UpdateMeeting(meeting)

	assert.NoError(t, err)
}

func TestDomainDeleteMeeting(t *testing.T) {
	id := "1"

	mockMeetingRepo := new(MockMeetingRepo)
	mockMeetingRepo.On("DeleteMeeting", id).Return(nil)
	md := domains.NewMeetingDomain(mockMeetingRepo)
	err := md.DeleteMeeting(id)

	assert.NoError(t, err)
}

func TestDomainGetMeeting(t *testing.T) {
	id := "1"
	meeting := &models.Meeting{
		Title:       "Board Meeting",
		Description: "Annual Board Meeting",
	}

	mockMeetingRepo := new(MockMeetingRepo)
	mockMeetingRepo.On("GetMeeting", id).Return(meeting, nil)
	md := domains.NewMeetingDomain(mockMeetingRepo)
	fetchedMeeting, err := md.GetMeeting(id)

	assert.NoError(t, err)
	assert.Equal(t, meeting.Title, fetchedMeeting.Title)
	assert.Equal(t, meeting.Description, fetchedMeeting.Description)
}
