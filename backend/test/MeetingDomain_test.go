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

func TestNewMeetingDomain(t *testing.T) {
	mockMeetingRepo := new(MockMeetingRepo)

	// t.Run("With 0 input", func(t *testing.T) {
	// 	md := domains.NewMeetingDomain()
	// 	assert.NotNil(t, md)
	// })

	t.Run("With 1 input", func(t *testing.T) {
		md := domains.NewMeetingDomain(mockMeetingRepo)
		assert.NotNil(t, md)
	})

	t.Run("With more than 1 input", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("NewMeetingDomain did not panic")
			}
		}()
		domains.NewMeetingDomain(mockMeetingRepo, mockMeetingRepo)
	})
}

func TestDomainCreateMeeting(t *testing.T) {
	meeting := &models.Meeting{
		Title:       "Board Meeting",
		Description: "Annual Board Meeting",
	}

	mockMeetingRepo := new(MockMeetingRepo)
	mockMeetingRepo.On("CreateMeeting", meeting).Return(nil).Once()
	md := domains.NewMeetingDomain(mockMeetingRepo)

	t.Run("With Normal Case", func(t *testing.T) {
		err := md.CreateMeeting(meeting)

		assert.NoError(t, err)
	})

	t.Run("Error in Repo", func(t *testing.T) {
		mockMeetingRepo.On("CreateMeeting", meeting).Return(assert.AnError).Once()
		err := md.CreateMeeting(meeting)

		assert.Error(t, err)
	})
}

func TestDomainUpdateMeeting(t *testing.T) {
	meeting := &models.Meeting{
		ID:          "1",
		Title:       "Updated Board Meeting",
		Description: "Updated Annual Board Meeting",
	}

	mockMeetingRepo := new(MockMeetingRepo)
	md := domains.NewMeetingDomain(mockMeetingRepo)

	t.Run("With Normal Case", func(t *testing.T) {
		mockMeetingRepo.On("UpdateMeeting", meeting).Return(nil).Once()

		err := md.UpdateMeeting(meeting)

		assert.NoError(t, err)
	})

	t.Run("Error in Repo", func(t *testing.T) {
		mockMeetingRepo.On("UpdateMeeting", meeting).Return(assert.AnError).Once()

		err := md.UpdateMeeting(meeting)

		assert.Error(t, err)
	})

}

func TestDomainDeleteMeeting(t *testing.T) {
	id := "1"

	mockMeetingRepo := new(MockMeetingRepo)
	md := domains.NewMeetingDomain(mockMeetingRepo)

	t.Run("With Normal Case", func(t *testing.T) {
		mockMeetingRepo.On("DeleteMeeting", id).Return(nil).Once()

		err := md.DeleteMeeting(id)

		assert.NoError(t, err)
	})

	t.Run("Error in Repo", func(t *testing.T) {
		mockMeetingRepo.On("DeleteMeeting", id).Return(assert.AnError).Once()

		err := md.DeleteMeeting(id)

		assert.Error(t, err)
	})
}

func TestDomainGetMeeting(t *testing.T) {
	id := "1"
	meeting := models.Meeting{
		Title:       "Board Meeting",
		Description: "Annual Board Meeting",
	}

	mockMeetingRepo := new(MockMeetingRepo)

	t.Run("With Normal Case", func(t *testing.T) {
		mockMeetingRepo.On("GetMeeting", id).Return(meeting, nil).Once()
		md := domains.NewMeetingDomain(mockMeetingRepo)
		fetchedMeeting, err := md.GetMeeting(id)

		assert.NoError(t, err)
		assert.Equal(t, meeting.Title, fetchedMeeting.Title)
		assert.Equal(t, meeting.Description, fetchedMeeting.Description)
	})

	t.Run("Error in Repo", func(t *testing.T) {
		mockMeetingRepo.On("GetMeeting", id).Return(models.Meeting{}, assert.AnError).Once()
		md := domains.NewMeetingDomain(mockMeetingRepo)
		fetchedMeeting, err := md.GetMeeting(id)

		assert.Error(t, err)
		assert.Equal(t, models.Meeting{}, fetchedMeeting)
	})
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

	t.Run("With Normal Case", func(t *testing.T) {
		mockMeetingRepo.On("GetAllMeetings").Return(meetings, nil).Once()
		md := domains.NewMeetingDomain(mockMeetingRepo)
		fetchedMeetings, err := md.GetAllMeetings()

		assert.NoError(t, err)
		assert.Equal(t, len(meetings), len(fetchedMeetings))
	})

	t.Run("Error in Repo", func(t *testing.T) {
		mockMeetingRepo.On("GetAllMeetings").Return([]models.Meeting{}, assert.AnError).Once()
		md := domains.NewMeetingDomain(mockMeetingRepo)
		fetchedMeetings, err := md.GetAllMeetings()

		assert.Error(t, err)
		assert.Empty(t, fetchedMeetings)
	})
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

	t.Run("With Normal Case", func(t *testing.T) {
		mockMeetingRepo.On("GetMeetingsByRoomId", roomID).Return(meetings, nil).Once()
		md := domains.NewMeetingDomain(mockMeetingRepo)
		fetchedMeetings, err := md.GetMeetingsByRoomId(roomID)

		assert.NoError(t, err)
		assert.Equal(t, len(meetings), len(fetchedMeetings))
	})

	t.Run("Error in Repo", func(t *testing.T) {
		mockMeetingRepo.On("GetMeetingsByRoomId", roomID).Return([]models.Meeting{}, assert.AnError).Once()
		md := domains.NewMeetingDomain(mockMeetingRepo)
		_, err := md.GetMeetingsByRoomId(roomID)

		assert.Error(t, err)
	})
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

	t.Run("With Normal Case", func(t *testing.T) {
		mockMeetingRepo.On("GetMeetingsByDatePeriod", dateFrom, dateTo).Return(meetings, nil).Once()
		md := domains.NewMeetingDomain(mockMeetingRepo)
		fetchedMeetings, err := md.GetMeetingsByDatePeriod(dateFrom, dateTo)

		assert.NoError(t, err)
		assert.Equal(t, len(meetings), len(fetchedMeetings))
	})

	t.Run("Error in Repo", func(t *testing.T) {
		mockMeetingRepo.On("GetMeetingsByDatePeriod", dateFrom, dateTo).Return([]models.Meeting{}, assert.AnError).Once()
		md := domains.NewMeetingDomain(mockMeetingRepo)
		_, err := md.GetMeetingsByDatePeriod(dateFrom, dateTo)

		assert.Error(t, err)
	})
}
