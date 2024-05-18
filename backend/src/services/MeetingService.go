package services

import (
	"meeting-center/src/domains"
	"meeting-center/src/models"
)

type MeetingService interface {
	CreateMeeting(meeting *models.Meeting) (*models.Meeting, error)
	UpdateMeeting(id string, meeting *models.Meeting) error
	DeleteMeeting(id string) error
	GetMeeting(id string) (*models.Meeting, error)
}

type meetingService struct {
	MeetingDomain domains.MeetingDomain
}

func NewMeetingService(opt ...domains.MeetingDomain) MeetingService {
	if len(opt) == 1 {
		return &meetingService{
			MeetingDomain: opt[0],
		}
	} else {
		return &meetingService{
			MeetingDomain: domains.NewMeetingDomain(),
		}
	}
}

func (ms *meetingService) CreateMeeting(meeting *models.Meeting) (*models.Meeting, error) {
	createdMeeting, err := ms.MeetingDomain.CreateMeeting(meeting)
	if err != nil {
		return nil, err
	}
	return createdMeeting, nil
}

func (ms *meetingService) UpdateMeeting(id string, meeting *models.Meeting) error {
	err := ms.MeetingDomain.UpdateMeeting(id, meeting)
	if err != nil {
		return err
	}
	return nil
}

func (ms *meetingService) DeleteMeeting(id string) error {
	err := ms.MeetingDomain.DeleteMeeting(id)
	if err != nil {
		return err
	}
	return nil
}

func (ms *meetingService) GetMeeting(id string) (*models.Meeting, error) {
	meeting, err := ms.MeetingDomain.GetMeeting(id)
	if err != nil {
		return nil, err
	}
	return meeting, nil
}
