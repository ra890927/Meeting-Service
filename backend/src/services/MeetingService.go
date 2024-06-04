package services

import (
	"errors"
	"meeting-center/src/domains"
	"meeting-center/src/mail"
	"meeting-center/src/models"
	"meeting-center/src/utils"
	"time"
)

type MeetingService interface {
	CreateMeeting(operator models.User, meeting *models.Meeting) error
	UpdateMeeting(operator models.User, meeting *models.Meeting) error
	DeleteMeeting(operator models.User, id string) error
	GetMeeting(id string) (models.Meeting, error)
	GetAllMeetings() ([]models.Meeting, error)
	GetMeetingsByRoomIdAndDatePeriod(roomID int, dateFrom time.Time, dateTo time.Time) ([]models.Meeting, error)
	GetMeetingsByParticipantId(participantID uint) ([]models.Meeting, error)
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

func (ms meetingService) intersectMeetingsById(meetingsArg ...[]models.Meeting) []models.Meeting {
	if len(meetingsArg) == 0 {
		return []models.Meeting{}
	}
	if len(meetingsArg) == 1 {
		return meetingsArg[0]
	}

	totalKeys := make(map[string]int)
	for _, meetings := range meetingsArg {
		for _, m := range meetings {
			totalKeys[m.ID]++
		}
	}

	var result []models.Meeting
	for _, m := range meetingsArg[0] {
		if totalKeys[m.ID] == len(meetingsArg) {
			result = append(result, m)
		}
	}
	return result
}

func (ms meetingService) CreateMeeting(operator models.User, meeting *models.Meeting) error {
	// modyfing the OrganizerID to the operator's ID
	meeting.OrganizerID = operator.ID

	err := ms.checkValidMeetingTime(*meeting)
	if err != nil {
		return err
	}

	err = ms.MeetingDomain.CreateMeeting(meeting)
	if err != nil {
		return errors.New("error when creating meeting")
	}

	mail.SendEmailByMeeting(*meeting, mail.NOTICE)

	return nil
}

func (ms meetingService) UpdateMeeting(operator models.User, meeting *models.Meeting) error {
	// find out the original meeting
	originalMeeting, err := ms.MeetingDomain.GetMeeting(meeting.ID)
	if err != nil {
		return errors.New("meeting not found")
	}

	permission := ms.getPermission(operator, originalMeeting)
	if !utils.CheckPermission(permission, utils.Update) {
		return errors.New("only the organizer can update the meeting")
	}

	err = ms.checkValidMeetingTime(*meeting)
	if err != nil {
		return err
	}

	err = ms.MeetingDomain.UpdateMeeting(meeting)
	if err != nil {
		return errors.New("error when updating meeting")
	}

	mail.SendEmailByMeeting(*meeting, mail.NOTICE)

	return nil
}

func (ms meetingService) DeleteMeeting(operator models.User, id string) error {
	// only the organizer can delete the meeting
	meeting, err := ms.MeetingDomain.GetMeeting(id)
	if err != nil {
		return errors.New("meeting not found")
	}

	permission := ms.getPermission(operator, meeting)
	if !utils.CheckPermission(permission, utils.Delete) {
		return errors.New("only the organizer can delete the meeting")
	}

	err = ms.MeetingDomain.DeleteMeeting(id)
	if err != nil {
		return err
	}
	return nil
}

func (ms meetingService) GetMeeting(id string) (models.Meeting, error) {
	meeting, err := ms.MeetingDomain.GetMeeting(id)
	if err != nil {
		return models.Meeting{}, errors.New("meeting not found")
	}
	return meeting, nil
}

func (ms meetingService) GetAllMeetings() ([]models.Meeting, error) {
	meetings, err := ms.MeetingDomain.GetAllMeetings()
	if err != nil {
		return []models.Meeting{}, errors.New("error when fetching meetings")
	}
	return meetings, nil
}

func (ms meetingService) GetMeetingsByRoomIdAndDatePeriod(roomID int, date_from time.Time, date_to time.Time) ([]models.Meeting, error) {
	meetingsByRoomId, err1 := ms.MeetingDomain.GetMeetingsByRoomId(roomID)
	meetingsByDatePeriod, err2 := ms.MeetingDomain.GetMeetingsByDatePeriod(date_from, date_to)
	if err1 != nil || err2 != nil {
		return []models.Meeting{}, errors.New("error when fetching meetings")
	}
	meetings := ms.intersectMeetingsById(meetingsByRoomId, meetingsByDatePeriod)
	return meetings, nil
}

func (ms meetingService) GetMeetingsByParticipantId(participantID uint) ([]models.Meeting, error) {
	// TODO: find a method to get all meetings by participant ID in repo layer
	meetings, err := ms.MeetingDomain.GetAllMeetings()
	if err != nil {
		return []models.Meeting{}, errors.New("error when fetching meetings")
	}

	var participantMeetings []models.Meeting
	for _, m := range meetings {
		for _, p := range m.Participants {
			if p == participantID {
				participantMeetings = append(participantMeetings, m)
				break
			}
		}
	}

	return participantMeetings, nil
}

func (ms meetingService) getPermission(operater models.User, meeting models.Meeting) utils.Permission {
	if operater.Role == "admin" || operater.ID == meeting.OrganizerID {
		return utils.Create | utils.Update | utils.Delete | utils.Read
	} else {
		return utils.Read
	}
}

// check if the meeting time is valid and not overlapping with other meetings
func (ms meetingService) checkValidMeetingTime(targetMeeting models.Meeting) error {
	if !targetMeeting.StartTime.Before(targetMeeting.EndTime) {
		return errors.New("end time should be after start time")
	}

	meetingsByRoomId, err1 := ms.MeetingDomain.GetMeetingsByRoomId(targetMeeting.RoomID)
	meetingsByDatePeriod, err2 := ms.MeetingDomain.GetMeetingsByDatePeriod(targetMeeting.StartTime, targetMeeting.EndTime)
	if err1 != nil || err2 != nil {
		return errors.New("error when fetching meetings")
	}

	// take the intersection of the two sets
	existingMeetings := ms.intersectMeetingsById(meetingsByRoomId, meetingsByDatePeriod)

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
