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
	GetMeetingsByRoomId(roomID int) ([]*models.Meeting, error)
	GetMeetingsByDatePeriod(dateFrom time.Time, dateTo time.Time) ([]*models.Meeting, error)
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

func (mr meetingRepo) GetMeetingsByRoomId(roomID int) ([]*models.Meeting, error) {
	var meetings []*models.Meeting
	result := mr.db.Where("room_id = ?", roomID).Find(&meetings)
	if result.Error != nil {
		return []*models.Meeting{}, result.Error
	}
	return meetings, nil
}

func (mr meetingRepo) GetMeetingsByDatePeriod(dateFrom time.Time, dateTo time.Time) ([]*models.Meeting, error) {
	// search for meetings is in the period of dateFrom and dateTo
	var meetings []*models.Meeting
	result := mr.db.Where("(DATE(start_time) BETWEEN ? AND ?) OR (DATE(end_time) BETWEEN ? AND ?) OR (DATE(start_time) <= ? AND DATE(end_time) >= ?)",
		dateFrom.Format("2006-01-02"), dateTo.Format("2006-01-02"),
		dateFrom.Format("2006-01-02"), dateTo.Format("2006-01-02"),
		dateFrom.Format("2006-01-02"), dateTo.Format("2006-01-02")).Find(&meetings)
	if result.Error != nil {
		return []*models.Meeting{}, result.Error
	}
	return meetings, nil
}
