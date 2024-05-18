package test

import (
	"meeting-center/src/domains"
	"meeting-center/src/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockMeetingRepo struct {
	mock.Mock
}

func (m *MockMeetingRepo) CreateMeeting(meeting *models.Meeting) (*models.Meeting, error) {
	args := m.Called(meeting)
	return args.Get(0).(*models.Meeting), args.Error(1)
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
	return args.Get(0).(*models.Meeting), args.Error(1)
}

func TestDomainCreateMeeting(t *testing.T) {
	meeting := &models.Meeting{
		Title:       "Board Meeting",
		Description: "Annual Board Meeting",
	}

	mockMeetingRepo := new(MockMeetingRepo)
	mockMeetingRepo.On("CreateMeeting", meeting).Return(meeting, nil)
	md := domains.NewMeetingDomain(mockMeetingRepo)
	createdMeeting, err := md.CreateMeeting(meeting)

	assert.NoError(t, err)
	assert.Equal(t, meeting.Title, createdMeeting.Title)
	assert.Equal(t, meeting.Description, createdMeeting.Description)
}

func TestDomainUpdateMeeting(t *testing.T) {
	id := "1"
	meeting := &models.Meeting{
		Title:       "Updated Board Meeting",
		Description: "Updated Annual Board Meeting",
	}

	mockMeetingRepo := new(MockMeetingRepo)
	mockMeetingRepo.On("UpdateMeeting", id, meeting).Return(nil)
	md := domains.NewMeetingDomain(mockMeetingRepo)
	err := md.UpdateMeeting(id, meeting)

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
