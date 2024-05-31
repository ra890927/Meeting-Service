package domains

import (
	"meeting-center/src/models"
	"meeting-center/src/repos"
	"time"
)

type MeetingDomain interface {
	CreateMeeting(meeting *models.Meeting) error
	UpdateMeeting(meeting *models.Meeting) error
	DeleteMeeting(id string) error
	GetMeeting(id string) (models.Meeting, error)
	GetAllMeetings() ([]models.Meeting, error)
	GetMeetingsByRoomId(room int) ([]models.Meeting, error)
	GetMeetingsByDatePeriod(dateFrom time.Time, dateTo time.Time) ([]models.Meeting, error)
}

type meetingDomain struct {
	MeetingRepo repos.MeetingRepo
}

func NewMeetingDomain(meetingRepoArg ...repos.MeetingRepo) MeetingDomain {
	if len(meetingRepoArg) == 1 {
		return MeetingDomain(&meetingDomain{
			MeetingRepo: meetingRepoArg[0],
		})
	} else if len(meetingRepoArg) == 0 {
		return MeetingDomain(&meetingDomain{
			MeetingRepo: repos.NewMeetingRepo(),
		})
	} else {
		panic("too many arguments")
	}
}

func (md meetingDomain) CreateMeeting(meeting *models.Meeting) error {
	err := md.MeetingRepo.CreateMeeting(meeting)
	if err != nil {
		return err
	}
	return nil
}

func (md meetingDomain) UpdateMeeting(meeting *models.Meeting) error {
	err := md.MeetingRepo.UpdateMeeting(meeting)
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

func (md meetingDomain) GetMeeting(id string) (models.Meeting, error) {
	meeting, err := md.MeetingRepo.GetMeeting(id)
	if err != nil {
		return models.Meeting{}, err
	}
	return meeting, nil
}

func (md meetingDomain) GetAllMeetings() ([]models.Meeting, error) {
	meetings, err := md.MeetingRepo.GetAllMeetings()
	if err != nil {
		return []models.Meeting{}, err
	}
	return meetings, nil
}

func (md meetingDomain) GetMeetingsByRoomId(room int) ([]models.Meeting, error) {
	meetings, err := md.MeetingRepo.GetMeetingsByRoomId(room)
	if err != nil {
		return []models.Meeting{}, err
	}
	return meetings, nil
}

func (md meetingDomain) GetMeetingsByDatePeriod(dateFrom time.Time, dateTo time.Time) ([]models.Meeting, error) {
	meetings, err := md.MeetingRepo.GetMeetingsByDatePeriod(dateFrom, dateTo)
	if err != nil {
		return []models.Meeting{}, err
	}
	return meetings, nil
}
