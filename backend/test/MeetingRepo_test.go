package test

import (
	"meeting-center/src/models"
	. "meeting-center/src/repos"
	"net/http"

	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type MeetingRepoTestSuite struct {
	suite.Suite
	mr MeetingRepo
	db *gorm.DB
}

func (suite *MeetingRepoTestSuite) SetupTest() {
	db, err := gorm.Open(sqlite.Open("./test.sqlite"), &gorm.Config{})
	assert.NoError(suite.T(), err)

	err = db.AutoMigrate(&models.Meeting{})
	assert.NoError(suite.T(), err)

	suite.db = db
	suite.mr = NewMeetingRepo(db)
}

func (suite *MeetingRepoTestSuite) TearDownTest() {
	db, err := suite.db.DB()
	assert.NoError(suite.T(), err)
	err = db.Close()
	assert.NoError(suite.T(), err)
}

func (suite *MeetingRepoTestSuite) TestCreateMeeting() {
	meeting := &models.Meeting{
		StartTime: time.Now(),
		EndTime:   time.Now().Add(time.Hour),
	}
	err := suite.mr.CreateMeeting(meeting)
	assert.NoError(suite.T(), err)
	assert.NotZero(suite.T(), meeting.ID)
}

func (suite *MeetingRepoTestSuite) TestUpdateMeeting() {
	meeting := &models.Meeting{
		RoomID:    1,
		StartTime: time.Now().Add(-time.Hour),
		EndTime:   time.Now().Add(time.Hour),
	}
	err := suite.mr.CreateMeeting(meeting)
	assert.NoError(suite.T(), err)

	meeting.StartTime = meeting.StartTime.Add(time.Hour)
	meeting.EndTime = meeting.EndTime.Add(time.Hour)
	err = suite.mr.UpdateMeeting(meeting)
	assert.NoError(suite.T(), err)
}

func (suite *MeetingRepoTestSuite) TestDeleteMeeting() {
	meeting := &models.Meeting{
		StartTime: time.Now().Add(-time.Hour),
		EndTime:   time.Now().Add(time.Hour),
	}
	err := suite.mr.CreateMeeting(meeting)
	assert.NoError(suite.T(), err)

	err = suite.mr.DeleteMeeting(meeting.ID)
	assert.NoError(suite.T(), err)
}

func (suite *MeetingRepoTestSuite) TestGetMeeting() {
	meeting := &models.Meeting{
		StartTime: time.Now().Add(-time.Hour),
		EndTime:   time.Now().Add(time.Hour),
	}
	err := suite.mr.CreateMeeting(meeting)
	assert.NoError(suite.T(), err)

	meeting2, err := suite.mr.GetMeeting(meeting.ID)
	assert.NoError(suite.T(), err)
	timeFormat := http.TimeFormat
	assert.Equal(suite.T(), meeting.StartTime.Format(timeFormat), meeting2.StartTime.Format(timeFormat))
	assert.Equal(suite.T(), meeting.EndTime.Format(timeFormat), meeting2.EndTime.Format(timeFormat))
}

func (suite *MeetingRepoTestSuite) TestGetAllMeetings() {
	// check the original length of the meetings
	meetings, err := suite.mr.GetAllMeetings()
	assert.NoError(suite.T(), err)
	originalLength := len(meetings)

	meeting := &models.Meeting{
		StartTime: time.Now().Add(-time.Hour),
		EndTime:   time.Now().Add(time.Hour),
	}
	err = suite.mr.CreateMeeting(meeting)
	assert.NoError(suite.T(), err)

	meetings, err = suite.mr.GetAllMeetings()
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), meetings, originalLength+1)
}

func (suite *MeetingRepoTestSuite) TestGetMeetingsByRoomIdAndDate() {
	// check the original length of the meetings
	meetings, err := suite.mr.GetMeetingsByRoomIdAndDate(1, time.Now())
	assert.NoError(suite.T(), err)
	originalLength := len(meetings)

	meeting := &models.Meeting{
		RoomID:    1,
		StartTime: time.Now().Add(-time.Hour),
		EndTime:   time.Now().Add(time.Hour),
	}
	err = suite.mr.CreateMeeting(meeting)
	assert.NoError(suite.T(), err)

	meetings, err = suite.mr.GetMeetingsByRoomIdAndDate(1, time.Now())
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), meetings, originalLength+1)
}

func TestMeetingRepoTestSuite(t *testing.T) {
	suite.Run(t, new(MeetingRepoTestSuite))
}
