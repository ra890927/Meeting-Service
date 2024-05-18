package test

import (
	"meeting-center/src/models"
	"meeting-center/src/repos"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockGormDB simulates the GormDB for testing
type MockGormDB struct {
	mock.Mock
}

func (db *MockGormDB) Create(value interface{}) *MockGormDB {
	args := db.Called(value)
	return args.Get(0).(*MockGormDB)
}

func (db *MockGormDB) Update(id string, value interface{}) *MockGormDB {
	args := db.Called(id, value)
	return args.Get(0).(*MockGormDB)
}

func (db *MockGormDB) Delete(value interface{}, id string) *MockGormDB {
	args := db.Called(value, id)
	return args.Get(0).(*MockGormDB)
}

func (db *MockGormDB) First(out interface{}, where ...interface{}) *MockGormDB {
	args := db.Called(out, where)
	return args.Get(0).(*MockGormDB)
}

func (db *MockGormDB) Find(out interface{}, where ...interface{}) *MockGormDB {
	args := db.Called(out, where)
	return args.Get(0).(*MockGormDB)
}

func TestMeetingRepo_CreateMeeting(t *testing.T) {
	mockDB := new(MockGormDB)
	meetingRepo := repos.MeetingRepo{DB: mockDB}

	meeting := &models.Meeting{
		Title:       "Board Meeting",
		Description: "Annual Board Meeting",
	}

	mockDB.On("Create", meeting).Return(mockDB)
	mockDB.On("Error").Return(nil)

	createdMeeting, err := meetingRepo.CreateMeeting(meeting)

	assert.NoError(t, err)
	assert.Equal(t, meeting.Title, createdMeeting.Title)
	assert.Equal(t, meeting.Description, createdMeeting.Description)
}

func TestMeetingRepo_UpdateMeeting(t *testing.T) {
	mockDB := new(MockGormDB)
	meetingRepo := repos.MeetingRepo{DB: mockDB}

	meeting := &models.Meeting{
		Title:       "Updated Board Meeting",
		Description: "Updated Annual Board Meeting",
	}
	id := "1"

	mockDB.On("Model", &models.Meeting{}).Return(mockDB)
	mockDB.On("Where", "id = ?", id).Return(mockDB)
	mockDB.On("Updates", meeting).Return(mockDB)
	mockDB.On("Error").Return(nil)

	err := meetingRepo.UpdateMeeting(id, meeting)

	assert.NoError(t, err)
}

func TestMeetingRepo_DeleteMeeting(t *testing.T) {
	mockDB := new(MockGormDB)
	meetingRepo := repos.MeetingRepo{DB: mockDB}
	id := "1"

	mockDB.On("Delete", &models.Meeting{}, id).Return(mockDB)
	mockDB.On("Error").Return(nil)

	err := meetingRepo.DeleteMeeting(id)

	assert.NoError(t, err)
}

func TestMeetingRepo_GetMeeting(t *testing.T) {
	mockDB := new(MockGormDB)
	meetingRepo := repos.MeetingRepo{DB: mockDB}
	id := "1"
	meeting := &models.Meeting{}

	mockDB.On("First", meeting, "id = ?", id).Return(mockDB)
	mockDB.On("Error").Return(nil)

	fetchedMeeting, err := meetingRepo.GetMeeting(id)

	assert.NoError(t, err)
	assert.Equal(t, meeting, fetchedMeeting)
}

func TestMeetingRepo_GetAllMeetings(t *testing.T) {
	mockDB := new(MockGormDB)
	meetingRepo := repos.MeetingRepo{DB: mockDB}
	expectedMeetings := []*models.Meeting{{ID: "1", Title: "Board Meeting"}}

	mockDB.On("Find", &[]*models.Meeting{}).Return(mockDB)
	mockDB.On("Error").Return(nil)
	mockDB.On("Value").Return(expectedMeetings)

	meetings, err := meetingRepo.GetAllMeetings()

	assert.NoError(t, err)
	assert.Equal(t, expectedMeetings, meetings)
}

func TestMeetingRepo_GetMeetingsByRoomIdAndDate(t *testing.T) {
	mockDB := new(MockGormDB)
	meetingRepo := repos.MeetingRepo{DB: mockDB}
	roomID := 1
	date := time.Now()
	expectedMeetings := []*models.Meeting{{ID: "1", Title: "Strategy Meeting"}}

	mockDB.On("Where", "room_id = ? AND date(start_time) = date(?)", roomID, date).Return(mockDB)
	mockDB.On("Find", &[]*models.Meeting{}).Return(mockDB)
	mockDB.On("Error").Return(nil)
	mockDB.On("Value").Return(expectedMeetings)

	meetings, err := meetingRepo.GetMeetingsByRoomIdAndDate(roomID, date)

	assert.NoError(t, err)
	assert.Equal(t, expectedMeetings, meetings)
}
