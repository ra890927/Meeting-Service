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

func (m *MockMeetingRepo) GetMeeting(id string) (models.Meeting, error) {
	args := m.Called(id)
	return args.Get(0).(models.Meeting), args.Error(1)
}

func (m *MockMeetingRepo) GetAllMeetings() ([]models.Meeting, error) {
	args := m.Called()
	return args.Get(0).([]models.Meeting), args.Error(1)
}

func (m *MockMeetingRepo) GetMeetingsByRoomId(roomID int) ([]models.Meeting, error) {
	args := m.Called(roomID)
	return args.Get(0).([]models.Meeting), args.Error(1)
}

func (m *MockMeetingRepo) GetMeetingsByDatePeriod(dateFrom time.Time, dateTo time.Time) ([]models.Meeting, error) {
	args := m.Called(dateFrom, dateTo)
	return args.Get(0).([]models.Meeting), args.Error(1)
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
	meeting := models.Meeting{
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

func TestDomainGetAllMeetings(t *testing.T) {
	meetings := []models.Meeting{
		{
			Title:       "Board Meeting",
			Description: "Annual Board Meeting",
		},
		{
			Title:       "Team Meeting",
			Description: "Weekly Team Meeting",
		},
	}

	mockMeetingRepo := new(MockMeetingRepo)
	mockMeetingRepo.On("GetAllMeetings").Return(meetings, nil)
	md := domains.NewMeetingDomain(mockMeetingRepo)
	fetchedMeetings, err := md.GetAllMeetings()

	assert.NoError(t, err)
	assert.Equal(t, len(meetings), len(fetchedMeetings))
}

func TestDomainGetMeetingsByRoomId(t *testing.T) {
	roomID := 1
	meetings := []models.Meeting{
		{
			RoomID:      roomID,
			Title:       "Board Meeting",
			Description: "Annual Board Meeting",
		},
		{
			RoomID:      roomID,
			Title:       "Team Meeting",
			Description: "Weekly Team Meeting",
		},
	}

	mockMeetingRepo := new(MockMeetingRepo)
	mockMeetingRepo.On("GetMeetingsByRoomId", roomID).Return(meetings, nil)
	md := domains.NewMeetingDomain(mockMeetingRepo)
	fetchedMeetings, err := md.GetMeetingsByRoomId(roomID)

	assert.NoError(t, err)
	assert.Equal(t, len(meetings), len(fetchedMeetings))
}

func TestDomainGetMeetingsByDatePeriod(t *testing.T) {
	dateFrom := time.Now()
	dateTo := time.Now().AddDate(0, 0, 1)
	meetings := []models.Meeting{
		{
			Title:       "Board Meeting",
			Description: "Annual Board Meeting",
			StartTime:   dateFrom,
			EndTime:     dateTo,
		},
		{
			Title:       "Team Meeting",
			Description: "Weekly Team Meeting",
			StartTime:   dateFrom,
			EndTime:     dateTo,
		},
	}

	mockMeetingRepo := new(MockMeetingRepo)
	mockMeetingRepo.On("GetMeetingsByDatePeriod", dateFrom, dateTo).Return(meetings, nil)
	md := domains.NewMeetingDomain(mockMeetingRepo)
	fetchedMeetings, err := md.GetMeetingsByDatePeriod(dateFrom, dateTo)

	assert.NoError(t, err)
	assert.Equal(t, len(meetings), len(fetchedMeetings))
}
