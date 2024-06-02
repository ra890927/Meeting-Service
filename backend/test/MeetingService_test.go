package test

import (
	"meeting-center/src/models"
	"meeting-center/src/services"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/datatypes"
)

type MockMeetingDomain struct {
	mock.Mock
}

func (m *MockMeetingDomain) CreateMeeting(meeting *models.Meeting) error {
	args := m.Called(meeting)
	return args.Error(0)
}

func (m *MockMeetingDomain) UpdateMeeting(meeting *models.Meeting) error {
	args := m.Called(meeting)
	return args.Error(0)
}

func (m *MockMeetingDomain) DeleteMeeting(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockMeetingDomain) GetMeeting(id string) (models.Meeting, error) {
	args := m.Called(id)
	return args.Get(0).(models.Meeting), args.Error(1)
}

func (m *MockMeetingDomain) GetAllMeetings() ([]models.Meeting, error) {
	args := m.Called()
	return args.Get(0).([]models.Meeting), args.Error(1)
}

func (m *MockMeetingDomain) GetMeetingsByRoomId(roomID int) ([]models.Meeting, error) {
	args := m.Called(roomID)
	return args.Get(0).([]models.Meeting), args.Error(1)
}

func (m *MockMeetingDomain) GetMeetingsByDatePeriod(dateFrom time.Time, dateTo time.Time) ([]models.Meeting, error) {
	args := m.Called(dateFrom, dateTo)
	return args.Get(0).([]models.Meeting), args.Error(1)
}

func TestNewMeetingService(t *testing.T) {
	mockMeetingDomain := new(MockMeetingDomain)

	t.Run("With 0 input", func(t *testing.T) {
		ms := services.NewMeetingService()
		assert.NotNil(t, ms)
	})

	t.Run("With 1 input", func(t *testing.T) {
		ms := services.NewMeetingService(mockMeetingDomain)
		assert.NotNil(t, ms)
	})

	t.Run("With more than 1 input", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("NewMeetingService did not panic")
			}
		}()
		services.NewMeetingService(mockMeetingDomain, mockMeetingDomain)
	})
}

func TestServiceCreateMeeting(t *testing.T) {
	mockMeetingDomain := new(MockMeetingDomain)
	ms := services.NewMeetingService(mockMeetingDomain)

	user := models.User{ID: 1}

	t.Run("Normal Case", func(t *testing.T) {
		meeting := &models.Meeting{
			Title:       "Invalid Time Meeting",
			Description: "This meeting has invalid time",
			RoomID:      101,
			StartTime:   time.Now().Add(-2 * time.Hour),
			EndTime:     time.Now(),
		}
		mockMeetingDomain.On("CreateMeeting", meeting).Return(nil).Once()
		mockMeetingDomain.On("GetMeetingsByRoomId", meeting.RoomID).Return([]models.Meeting{}, nil).Once()
		mockMeetingDomain.On("GetMeetingsByDatePeriod", meeting.StartTime, meeting.EndTime).Return([]models.Meeting{}, nil).Once()
		err := ms.CreateMeeting(user, meeting)

		assert.NoError(t, err)
	})

	t.Run("With Invalid Time", func(t *testing.T) {
		meeting := &models.Meeting{
			Title:       "Invalid Time Meeting",
			Description: "This meeting has invalid time",
			RoomID:      101,
			StartTime:   time.Now().Add(2 * time.Hour),
			EndTime:     time.Now(),
		}

		err := ms.CreateMeeting(user, meeting)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "end time should be after start time")
	})

	t.Run("With Conflict", func(t *testing.T) {
		meeting := &models.Meeting{
			Title:       "Board Meeting",
			Description: "Annual Board Meeting",
			RoomID:      101,
			StartTime:   time.Now(),
			EndTime:     time.Now().Add(2 * time.Hour),
		}

		existingMeeting := models.Meeting{
			ID:        "2",
			Title:     "Strategy Meeting",
			StartTime: time.Now().Add(1 * time.Hour),
			EndTime:   time.Now().Add(3 * time.Hour),
		}

		mockMeetingDomain.On("CreateMeeting", meeting).Return(nil).Once()
		mockMeetingDomain.On("GetMeetingsByRoomId", meeting.RoomID).Return([]models.Meeting{existingMeeting}, nil).Once()
		mockMeetingDomain.On("GetMeetingsByDatePeriod", meeting.StartTime, meeting.EndTime).Return([]models.Meeting{existingMeeting}, nil).Once()

		err := ms.CreateMeeting(user, meeting)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "meeting overlaps with existing meeting")
	})

	t.Run("With Error in Domain", func(t *testing.T) {
		meeting := &models.Meeting{
			Title:       "Invalid Time Meeting",
			Description: "This meeting has invalid time",
			RoomID:      101,
			StartTime:   time.Now().Add(-2 * time.Hour),
			EndTime:     time.Now(),
		}

		mockMeetingDomain.On("CreateMeeting", meeting).Return(assert.AnError).Once()
		mockMeetingDomain.On("GetMeetingsByRoomId", meeting.RoomID).Return([]models.Meeting{}, assert.AnError).Once()
		mockMeetingDomain.On("GetMeetingsByDatePeriod", meeting.StartTime, meeting.EndTime).Return([]models.Meeting{}, assert.AnError).Once()

		err := ms.CreateMeeting(user, meeting)

		assert.Error(t, err)
	})
}

func TestServiceUpdateMeetingWithInvalidTime(t *testing.T) {
	mockMeetingDomain := new(MockMeetingDomain)
	ms := services.NewMeetingService(mockMeetingDomain)

	user := models.User{ID: 1}

	meeting := &models.Meeting{
		ID:          "1",
		Title:       "Invalid Time Meeting",
		Description: "This meeting has invalid time",
		OrganizerID: user.ID,
		StartTime:   time.Now().Add(2 * time.Hour),
		EndTime:     time.Now(),
	}

	mockMeetingDomain.On("GetMeeting", meeting.ID).Return(*meeting, nil)

	err := ms.UpdateMeeting(user, meeting)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "end time should be after start time")
}

func TestServiceUpdateMeetingWithConflict(t *testing.T) {
	mockMeetingDomain := new(MockMeetingDomain)
	ms := services.NewMeetingService(mockMeetingDomain)

	user := models.User{ID: 1}

	meeting := &models.Meeting{
		ID:          "1",
		Title:       "Updated Board Meeting",
		OrganizerID: user.ID,
		Description: "Updated Annual Board Meeting",
		StartTime:   time.Now(),
		EndTime:     time.Now().Add(2 * time.Hour),
	}

	existingMeeting := models.Meeting{
		ID:        "2",
		Title:     "Strategy Meeting",
		StartTime: time.Now().Add(1 * time.Hour),
		EndTime:   time.Now().Add(3 * time.Hour),
	}

	mockMeetingDomain.On("UpdateMeeting", meeting).Return(nil)
	mockMeetingDomain.On("GetMeeting", meeting.ID).Return(*meeting, nil)
	mockMeetingDomain.On("GetMeetingsByRoomId", meeting.RoomID).Return([]models.Meeting{existingMeeting}, nil)
	mockMeetingDomain.On("GetMeetingsByDatePeriod", meeting.StartTime, meeting.EndTime).Return([]models.Meeting{existingMeeting}, nil)

	err := ms.UpdateMeeting(user, meeting)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "meeting overlaps with existing meeting")

}

func TestServiceUpdateMeetingWithNoConflict(t *testing.T) {
	mockMeetingDomain := new(MockMeetingDomain)
	ms := services.NewMeetingService(mockMeetingDomain)

	user := models.User{ID: 1}

	meeting := &models.Meeting{
		ID:          "1",
		Title:       "Updated Board Meeting",
		Description: "Updated Annual Board Meeting",
		OrganizerID: user.ID,
		StartTime:   time.Now(),
		EndTime:     time.Now().Add(2 * time.Hour),
	}

	mockMeetingDomain.On("UpdateMeeting", meeting).Return(nil)
	mockMeetingDomain.On("GetMeeting", meeting.ID).Return(*meeting, nil)
	mockMeetingDomain.On("GetMeetingsByRoomId", meeting.RoomID).Return([]models.Meeting{}, nil)
	mockMeetingDomain.On("GetMeetingsByDatePeriod", meeting.StartTime, meeting.EndTime).Return([]models.Meeting{}, nil)

	err := ms.UpdateMeeting(user, meeting)
	assert.NoError(t, err)
}

func TestServiceDeleteMeeting(t *testing.T) {
	mockMeetingDomain := new(MockMeetingDomain)
	ms := services.NewMeetingService(mockMeetingDomain)

	user := models.User{ID: 1}
	id := "1"
	mockMeetingDomain.On("GetMeeting", id).Return(models.Meeting{ID: id, OrganizerID: user.ID}, nil)
	mockMeetingDomain.On("DeleteMeeting", id).Return(nil)

	err := ms.DeleteMeeting(user, id)

	assert.NoError(t, err)
}

func TestServiceGetMeeting(t *testing.T) {
	mockMeetingDomain := new(MockMeetingDomain)
	ms := services.NewMeetingService(mockMeetingDomain)

	id := "1"
	meeting := models.Meeting{
		Title:       "Board Meeting",
		Description: "Annual Board Meeting",
	}

	mockMeetingDomain.On("GetMeeting", id).Return(meeting, nil)

	fetchedMeeting, err := ms.GetMeeting(id)

	assert.NoError(t, err)
	assert.Equal(t, meeting, fetchedMeeting)
}

func TestServiceGetAllMeetings(t *testing.T) {
	mockMeetingDomain := new(MockMeetingDomain)
	ms := services.NewMeetingService(mockMeetingDomain)
	expectedMeetings := []models.Meeting{
		{ID: "1", Title: "Board Meeting", Description: "Annual Board Meeting"},
	}

	t.Run("Normal Case", func(t *testing.T) {
		mockMeetingDomain.On("GetAllMeetings").Return(expectedMeetings, nil).Once()

		meetings, err := ms.GetAllMeetings()

		assert.NoError(t, err)
		assert.Equal(t, expectedMeetings, meetings)
	})

	t.Run("With Error in Domain", func(t *testing.T) {
		mockMeetingDomain.On("GetAllMeetings").Return([]models.Meeting{}, assert.AnError).Once()

		meetings, err := ms.GetAllMeetings()

		assert.Error(t, err)
		assert.Empty(t, meetings)
	})

}

func TestServiceGetMeetingsByRoomIdAndDatePeriod(t *testing.T) {
	mockMeetingDomain := new(MockMeetingDomain)
	ms := services.NewMeetingService(mockMeetingDomain)
	roomID := 10
	date := time.Now()
	expectedMeetings := []models.Meeting{{ID: "2", Title: "Strategy Meeting", Description: "Strategy planning session"}}

	mockMeetingDomain.On("GetMeetingsByRoomId", roomID).Return(expectedMeetings, nil)
	mockMeetingDomain.On("GetMeetingsByDatePeriod", date, date).Return(expectedMeetings, nil)

	meetings, err := ms.GetMeetingsByRoomIdAndDatePeriod(roomID, date, date)

	assert.NoError(t, err)
	assert.Equal(t, expectedMeetings, meetings)
}

func TestServiceGetMeetingsByParticipantId(t *testing.T) {
	mockMeetingDomain := new(MockMeetingDomain)
	ms := services.NewMeetingService(mockMeetingDomain)
	participantID := uint(1)
	expectedMeetings := []models.Meeting{
		{
			ID:           "2",
			Participants: datatypes.JSONSlice[uint]{participantID},
		},
	}

	mockMeetingDomain.On("GetAllMeetings").Return(expectedMeetings, nil)

	meetings, err := ms.GetMeetingsByParticipantId(participantID)

	assert.NoError(t, err)
	assert.Len(t, meetings, len(expectedMeetings))
}
