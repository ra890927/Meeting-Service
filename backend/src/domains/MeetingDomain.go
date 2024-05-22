package domains

import (
	"meeting-center/src/models"
	"meeting-center/src/repos"
	"time"
)

type MeetingDomain interface {
	CreateMeeting(meeting *models.Meeting) (*models.Meeting, error)
	UpdateMeeting(id string, meeting *models.Meeting) error
	DeleteMeeting(id string) error
	GetMeeting(id string) (*models.Meeting, error)
	GetAllMeetings() ([]*models.Meeting, error)
	GetMeetingsByRoomIdAndDate(roomID int, date time.Time) ([]*models.Meeting, error)
}

type meetingDomain struct {
	MeetingRepo repos.MeetingRepo
}

func NewMeetingDomain(meetingRepoArg ...repos.MeetingRepo) MeetingDomain {
	if len(meetingRepoArg) == 1 {
		return MeetingDomain(&meetingDomain{
			MeetingRepo: meetingRepoArg[0],
		})
	} else {
		return MeetingDomain(&meetingDomain{
			MeetingRepo: repos.NewMeetingRepo(),
		})
	}
}

func (md meetingDomain) CreateMeeting(meeting *models.Meeting) (*models.Meeting, error) {
	createdMeeting, err := md.MeetingRepo.CreateMeeting(meeting)
	if err != nil {
		return nil, err
	}
	return createdMeeting, nil
}

func (md meetingDomain) UpdateMeeting(id string, meeting *models.Meeting) error {
	err := md.MeetingRepo.UpdateMeeting(id, meeting)
	if err != nil {
		return err
	}
	return nil
}

func (md meetingDomain) DeleteMeeting(id string) error {
	err := md.MeetingRepo.DeleteMeeting(id)
	if err != nil {
		return err
	}
	return nil
}

func (md meetingDomain) GetMeeting(id string) (*models.Meeting, error) {
	meeting, err := md.MeetingRepo.GetMeeting(id)
	if err != nil {
		return nil, err
	}
	return meeting, nil
}

func (md meetingDomain) GetAllMeetings() ([]*models.Meeting, error) {
	return md.MeetingRepo.GetAllMeetings()
}

func (md meetingDomain) GetMeetingsByRoomIdAndDate(roomID int, date time.Time) ([]*models.Meeting, error) {
	return md.MeetingRepo.GetMeetingsByRoomIdAndDate(roomID, date)
}