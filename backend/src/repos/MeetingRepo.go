package repos

import (
	db "meeting-center/src/io"
	"meeting-center/src/models"
	"time"

	"gorm.io/gorm"
)

type MeetingRepo interface {
	CreateMeeting(meeting *models.Meeting) error
	UpdateMeeting(meeting *models.Meeting) error
	DeleteMeeting(id string) error
	GetMeeting(id string) (*models.Meeting, error)
	GetAllMeetings() ([]*models.Meeting, error)
	GetMeetingsByRoomIdAndDate(roomID int, date time.Time) ([]*models.Meeting, error)
}

type meetingRepo struct {
	db *gorm.DB
}

func NewMeetingRepo(dbArgs ...*gorm.DB) MeetingRepo {
	if len(dbArgs) == 0 {
		return MeetingRepo(&meetingRepo{db: db.GetDBInstance()})
	} else if len(dbArgs) == 1 {
		return MeetingRepo(&meetingRepo{db: dbArgs[0]})
	} else {
		panic("Too many arguments")
	}
}

func (mr meetingRepo) CreateMeeting(meeting *models.Meeting) error {
	result := mr.db.Create(meeting)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (mr meetingRepo) UpdateMeeting(meeting *models.Meeting) error {
	result := mr.db.Model(&models.Meeting{}).Where("id = ?", meeting.ID).Updates(meeting)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (mr meetingRepo) DeleteMeeting(id string) error {
	result := mr.db.Where("id = ?", id).Delete(&models.Meeting{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (mr meetingRepo) GetMeeting(id string) (*models.Meeting, error) {
	var meeting models.Meeting
	result := mr.db.Where("id = ?", id).First(&meeting)
	if result.Error != nil {
		return nil, result.Error
	}
	return &meeting, nil
}

func (mr meetingRepo) GetAllMeetings() ([]*models.Meeting, error) {
	var meetings []*models.Meeting
	result := mr.db.Find(&meetings)
	if result.Error != nil {
		return nil, result.Error
	}
	return meetings, nil
}

func (mr meetingRepo) GetMeetingsByRoomIdAndDate(roomID int, date time.Time) ([]*models.Meeting, error) {
	// here get the meetings by room id and date(start_time)
	var meetings []*models.Meeting
	result := mr.db.Where("room_id = ? AND date(start_time) = date(?)", roomID, date).Find(&meetings)
	if result.Error != nil {
		return nil, result.Error
	}
	return meetings, nil
}
