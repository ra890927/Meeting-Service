package domains

import (
	"meeting-center/src/models"
	"meeting-center/src/repos"
)

type MeetingDomain interface {
	CreateMeeting(meeting *models.Meeting) (*models.Meeting, error)
	UpdateMeeting(id string, meeting *models.Meeting) error
	DeleteMeeting(id string) error
	GetMeeting(id string) (*models.Meeting, error)
}

type meetingDomain struct {
	MeetingRepo repos.MeetingRepo
}

func NewMeetingDomain(opt ...repos.MeetingRepo) MeetingDomain {
	if len(opt) == 1 {
		return &meetingDomain{
			MeetingRepo: opt[0],
		}
	} else {
		return &meetingDomain{
			MeetingRepo: repos.NewMeetingRepo(),
		}
	}
}

func (md *meetingDomain) CreateMeeting(meeting *models.Meeting) (*models.Meeting, error) {
	createdMeeting, err := md.MeetingRepo.CreateMeeting(meeting)
	if err != nil {
		return nil, err
	}
	return createdMeeting, nil
}

func (md *meetingDomain) UpdateMeeting(id string, meeting *models.Meeting) error {
	err := md.MeetingRepo.UpdateMeeting(id, meeting)
	if err != nil {
		return err
	}
	return nil
}

func (md *meetingDomain) DeleteMeeting(id string) error {
	err := md.MeetingRepo.DeleteMeeting(id)
	if err != nil {
		return err
	}
	return nil
}

func (md *meetingDomain) GetMeeting(id string) (*models.Meeting, error) {
	meeting, err := md.MeetingRepo.GetMeeting(id)
	if err != nil {
		return nil, err
	}
	return meeting, nil
}
