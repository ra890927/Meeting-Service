package test

import (
	"meeting-center/src/models"
	"meeting-center/src/services"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockMeetingDomain struct {
	mock.Mock
}

func (m *MockMeetingDomain) CreateMeeting(meeting *models.Meeting) (*models.Meeting, error) {
	args := m.Called(meeting)
	return args.Get(0).(*models.Meeting), args.Error(1)
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
	return args.Get(0).(*models.Meeting), args.Error(1)
}

func TestServiceCreateMeeting(t *testing.T) {
	mockMeetingDomain := new(MockMeetingDomain)
	ms := services.NewMeetingService(mockMeetingDomain)

	meeting := &models.Meeting{
		Title:       "Board Meeting",
		Description: "Annual Board Meeting",
	}

	mockMeetingDomain.On("CreateMeeting", meeting).Return(meeting, nil)

	createdMeeting, err := ms.CreateMeeting(meeting)

	assert.NoError(t, err)
	assert.Equal(t, meeting, createdMeeting)
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

func TestServiceDeleteMeeting(t *testing.T) {
	mockMeetingDomain := new(MockMeetingDomain)
	ms := services.NewMeetingService(mockMeetingDomain)

	id := "1"

	mockMeetingDomain.On("DeleteMeeting", id).Return(nil)

	err := ms.DeleteMeeting(id)

	assert.NoError(t, err)
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
