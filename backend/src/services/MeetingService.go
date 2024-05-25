package services

import (
	"errors"
	"fmt"
	"meeting-center/src/domains"
	"meeting-center/src/models"
	"time"
)

type MeetingService interface {
	CreateMeeting(operator *models.User, meeting *models.Meeting) error
	UpdateMeeting(operator *models.User, meeting *models.Meeting) error
	DeleteMeeting(operator *models.User, id int) error
	GetMeeting(id int) (*models.Meeting, error)
	GetAllMeetings() ([]*models.Meeting, error)
	GetMeetingsByRoomIdAndDate(roomID int, date time.Time) ([]*models.Meeting, error)
}

type meetingService struct {
	MeetingDomain domains.MeetingDomain
}

func NewMeetingService(roomDomainArg ...domains.MeetingDomain) MeetingService {
	if len(roomDomainArg) == 1 {
		return MeetingService(&meetingService{MeetingDomain: roomDomainArg[0]})
	} else if len(roomDomainArg) == 0 {
		return MeetingService(&meetingService{MeetingDomain: domains.NewMeetingDomain()})
	} else {
		panic("too many arguments")
	}
}

// check if the meeting time is valid and not overlapping with other meetings
func (ms meetingService) checkValidMeetingTime(targetMeeting *models.Meeting) error {
	if !targetMeeting.StartTime.Before(targetMeeting.EndTime) {
		return errors.New("end time should be after start time")
	}
	fmt.Println(targetMeeting.Title)

	existingMeetings, err := ms.MeetingDomain.GetMeetingsByRoomIdAndDate(targetMeeting.RoomID, targetMeeting.StartTime)
	if err != nil {
		return errors.New("error when getting existing meetings")
	}
	fmt.Println(existingMeetings)
	for _, m := range existingMeetings {
		// except the target meeting itself
		if m.ID == targetMeeting.ID {
			continue
		}

		// check if the meeting overlaps with existing meetings
		if targetMeeting.StartTime.Before(m.EndTime) && m.StartTime.Before(targetMeeting.EndTime) {
			return errors.New("meeting overlaps with existing meeting")
		}
	}
	return nil
}

func (ms meetingService) CreateMeeting(operator *models.User, meeting *models.Meeting) error {
	// modyfing the OrganizerID to the operator's ID
	meeting.OrganizerID = operator.ID

	err := ms.checkValidMeetingTime(meeting)
	if err != nil {
		return err
	}

	err = ms.MeetingDomain.CreateMeeting(meeting)
	if err != nil {
		return errors.New("error when creating meeting")
	}
	return nil
}

func (ms meetingService) UpdateMeeting(operator *models.User, meeting *models.Meeting) error {
	// check if the operator is the creator of the meeting
	if operator.ID != meeting.OrganizerID {
		return errors.New("only the organizer can update the meeting")
	}

	err := ms.checkValidMeetingTime(meeting)
	if err != nil {
		return err
	}

	err = ms.MeetingDomain.UpdateMeeting(meeting)
	if err != nil {
		return err
	}
	return nil
}

func (ms meetingService) DeleteMeeting(operator *models.User, id int) error {
	// only the organizer can delete the meeting
	meeting, err := ms.MeetingDomain.GetMeeting(id)
	if err != nil {
		return err
	}
	if operator.ID != meeting.OrganizerID {
		return errors.New("only the organizer can delete the meeting")
	}

	err = ms.MeetingDomain.DeleteMeeting(id)
	if err != nil {
		return err
	}
	return nil
}

func (ms meetingService) GetMeeting(id int) (*models.Meeting, error) {
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
