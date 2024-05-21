package services

import (
	"errors"
	"meeting-center/src/domains"
	"meeting-center/src/models"
	"time"
)

type MeetingService interface {
	CreateMeeting(meeting *models.Meeting) (*models.Meeting, error)
	UpdateMeeting(id string, meeting *models.Meeting) error
	DeleteMeeting(id string) error
	GetMeeting(id string) (*models.Meeting, error)
	GetAllMeetings() ([]*models.Meeting, error)
	GetMeetingsByRoomIdAndDate(roomID int, date time.Time) ([]*models.Meeting, error)
}

type meetingService struct {
	MeetingDomain domains.MeetingDomain
}

func NewMeetingService(roomRepoArg ...domains.MeetingDomain) MeetingService {
	if len(roomRepoArg) == 1 {
		return &meetingService{
			MeetingDomain: roomRepoArg[0],
		}
	} else {
		return &meetingService{
			MeetingDomain: domains.NewMeetingDomain(),
		}
	}
}

func (ms meetingService) CreateMeeting(meeting *models.Meeting) (*models.Meeting, error) {
	// Validate time
	if !meeting.StartTime.Before(meeting.EndTime) {
		return nil, errors.New("StartTime must be before EndTime")
	}

	// Check for overlapping meetings
	existingMeetings, err := ms.GetMeetingsByRoomIdAndDate(meeting.RoomID, meeting.StartTime)
	if err != nil {
		return nil, err
	}
	for _, m := range existingMeetings {
		if meeting.StartTime.Before(m.EndTime) && meeting.EndTime.After(m.StartTime) {
			return nil, errors.New("time slot is already booked")
		}
	}

	createdMeeting, err := ms.MeetingDomain.CreateMeeting(meeting)
	if err != nil {
		return nil, err
	}
	return createdMeeting, nil
}

func (ms meetingService) UpdateMeeting(id string, meeting *models.Meeting) error {
	err := ms.MeetingDomain.UpdateMeeting(id, meeting)
	if err != nil {
		return err
	}
	return nil
}

func (ms meetingService) DeleteMeeting(id string) error {
	err := ms.MeetingDomain.DeleteMeeting(id)
	if err != nil {
		return err
	}
	return nil
}

func (ms meetingService) GetMeeting(id string) (*models.Meeting, error) {
	meeting, err := ms.MeetingDomain.GetMeeting(id)
	if err != nil {
		return nil, err
	}
	return meeting, nil
}

func (ms meetingService) GetAllMeetings() ([]*models.Meeting, error) {
	meetings, err := ms.MeetingDomain.GetAllMeetings()
	if err != nil {
		return nil, err
	}
	return meetings, nil
}

func (ms meetingService) GetMeetingsByRoomIdAndDate(roomID int, date time.Time) ([]*models.Meeting, error) {
	meetings, err := ms.MeetingDomain.GetMeetingsByRoomIdAndDate(roomID, date)
	if err != nil {
		return nil, err
	}
	return meetings, nil
}
