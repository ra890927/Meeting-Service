package repos

import (
	"meeting-center/src/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type MeetingRepo interface {
	CreateMeeting(meeting *models.Meeting) (*models.Meeting, error)
	UpdateMeeting(id string, meeting *models.Meeting) error
	DeleteMeeting(id string) error
	GetMeeting(id string) (*models.Meeting, error)
}

type meetingRepo struct {
	dsn string
}

func NewMeetingRepo(opt ...string) MeetingRepo {
	dsn := "../sqlite.db"
	if len(opt) == 1 {
		dsn = opt[0]
	}
	return &meetingRepo{
		dsn: dsn,
	}
}

func (mr *meetingRepo) CreateMeeting(meeting *models.Meeting) (*models.Meeting, error) {
	db, err := gorm.Open(sqlite.Open(mr.dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	result := db.Create(meeting)
	if result.Error != nil {
		return nil, result.Error
	}

	return meeting, nil
}

func (mr *meetingRepo) UpdateMeeting(id string, meeting *models.Meeting) error {
	db, err := gorm.Open(sqlite.Open(mr.dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	result := db.Model(&models.Meeting{}).Where("id = ?", id).Updates(meeting)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (mr *meetingRepo) DeleteMeeting(id string) error {
	db, err := gorm.Open(sqlite.Open(mr.dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	result := db.Delete(&models.Meeting{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (mr *meetingRepo) GetMeeting(id string) (*models.Meeting, error) {
	db, err := gorm.Open(sqlite.Open(mr.dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	var meeting models.Meeting
	result := db.First(&meeting, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &meeting, nil
}
