package presentations

import (
	"fmt"
	"meeting-center/src/models"
	"meeting-center/src/services"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type MeetingPresentation interface {
	CreateMeeting(c *gin.Context)
	UpdateMeeting(c *gin.Context)
	DeleteMeeting(c *gin.Context)
	GetMeeting(c *gin.Context)
	GetAllMeetings(c *gin.Context)
	GetMeetingsByRoomIdAndDatePeriod(c *gin.Context)
	GetMeetingsByParticipantId(c *gin.Context)
}

type meetingPresentation struct {
	meetingService services.MeetingService
}

type CreateMeetingBody struct {
	RoomID       int       `json:"room_id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	StartTime    time.Time `json:"start_time"`
	EndTime      time.Time `json:"end_time"`
	OrganizerID  uint      `json:"organizer"`
	Participants []uint    `json:"participants"`
	StatusType   string    `json:"status_type"`
}

type CreateUpdateGetMeetingResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Meeting struct {
			ID           string    `json:"id"`
			RoomID       int       `json:"room_id"`
			Title        string    `json:"title"`
			OrganizerID  uint      `json:"organizer"`
			Participants []uint    `json:"participants"`
			Description  string    `json:"description"`
			StartTime    time.Time `json:"start_time"`
			EndTime      time.Time `json:"end_time"`
			StatusType   string    `json:"status_type"`
		} `json:"meeting"`
	} `json:"data"`
}

type UpdateMeetingBody struct {
	ID           string    `json:"id"`
	RoomID       int       `json:"room_id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	OrganizerID  uint      `json:"organizer"`
	Participants []uint    `json:"participants"`
	StartTime    time.Time `json:"start_time"`
	EndTime      time.Time `json:"end_time"`
	StatusType   string    `json:"status_type"`
}

type DeleteMeetingResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type GetAllMeetingsResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Meetings []struct {
			ID           string    `json:"id"`
			RoomID       int       `json:"room_id"`
			Title        string    `json:"title"`
			OrganizerID  uint      `json:"organizer"`
			Participants []uint    `json:"participants"`
			Description  string    `json:"description"`
			StartTime    time.Time `json:"start_time"`
			EndTime      time.Time `json:"end_time"`
			StatusType   string    `json:"status_type"`
		} `json:"meetings"`
	} `json:"data"`
}

func NewMeetingPresentation(roomServiceArg ...services.MeetingService) MeetingPresentation {
	if len(roomServiceArg) == 1 {
		return MeetingPresentation(&meetingPresentation{
			meetingService: roomServiceArg[0],
		})
	} else if len(roomServiceArg) == 0 {
		return MeetingPresentation(&meetingPresentation{
			meetingService: services.NewMeetingService(),
		})
	} else {
		panic("too many arguments")
	}
}

// @Summary Create a meeting
// @Description Create a meeting
// @Tags Meeting
// @Accept json
// @Produce json
// @Param meeting body CreateMeetingBody true "Meeting details"
// @Success 200 {object} CreateUpdateGetMeetingResponse
// @Router /meeting [post]
func (mp meetingPresentation) CreateMeeting(c *gin.Context) {
	operator := c.MustGet("validate_user").(models.User)
	var body CreateMeetingBody
	var response CreateUpdateGetMeetingResponse
	if err := c.BindJSON(&body); err != nil {
		response.Status = "error"
		response.Message = "Invalid request"
		response.Message = fmt.Sprintf("%v", err)
		c.JSON(400, response)
		return
	}

	meeting := models.Meeting{
		RoomID:       body.RoomID,
		Title:        body.Title,
		Description:  body.Description,
		OrganizerID:  body.OrganizerID,
		Participants: body.Participants,
		StartTime:    body.StartTime,
		EndTime:      body.EndTime,
		StatusType:   body.StatusType,
	}

	err := mp.meetingService.CreateMeeting(operator, &meeting)
	if err != nil {
		response.Status = "error"
		response.Message = err.Error()
		c.JSON(500, response)
		return
	}

	response.Status = "success"
	response.Message = "Meeting created"
	response.Data.Meeting.ID = meeting.ID
	response.Data.Meeting.RoomID = meeting.RoomID
	response.Data.Meeting.Title = meeting.Title
	response.Data.Meeting.Description = meeting.Description
	response.Data.Meeting.OrganizerID = meeting.OrganizerID
	response.Data.Meeting.Participants = meeting.Participants
	response.Data.Meeting.StartTime = meeting.StartTime
	response.Data.Meeting.EndTime = meeting.EndTime
	response.Data.Meeting.StatusType = meeting.StatusType
	c.JSON(200, response)
}

// @Summary Update a meeting
// @Description Update a meeting
// @Tags Meeting
// @Accept json
// @Produce json
// @Param meeting body UpdateMeetingBody true "Meeting details"
// @Success 200 {object} CreateUpdateGetMeetingResponse
// @Router /meeting [put]
func (mp meetingPresentation) UpdateMeeting(c *gin.Context) {
	var operator models.User
	operator = c.MustGet("validate_user").(models.User)

	var body UpdateMeetingBody
	var response CreateUpdateGetMeetingResponse
	if err := c.BindJSON(&body); err != nil {
		response.Status = "error"
		response.Message = fmt.Sprintf("%v", err)
		c.JSON(400, response)
		return
	}

	meeting := models.Meeting{
		ID:           body.ID,
		RoomID:       body.RoomID,
		Title:        body.Title,
		Description:  body.Description,
		OrganizerID:  body.OrganizerID,
		Participants: body.Participants,
		StartTime:    body.StartTime,
		EndTime:      body.EndTime,
		StatusType:   body.StatusType,
	}

	err := mp.meetingService.UpdateMeeting(operator, &meeting)
	if err != nil {
		response.Status = "error"
		response.Message = "Internal server error"
		c.JSON(500, response)
		return
	}

	response.Status = "success"
	response.Message = "Meeting updated"
	response.Data.Meeting.ID = meeting.ID
	response.Data.Meeting.RoomID = meeting.RoomID
	response.Data.Meeting.Title = meeting.Title
	response.Data.Meeting.Description = meeting.Description
	response.Data.Meeting.OrganizerID = meeting.OrganizerID
	response.Data.Meeting.Participants = meeting.Participants
	response.Data.Meeting.StartTime = meeting.StartTime
	response.Data.Meeting.EndTime = meeting.EndTime
	response.Data.Meeting.StatusType = meeting.StatusType
	c.JSON(200, response)
}

// @Summary Delete a meeting
// @Description Delete a meeting
// @Tags Meeting
// @Param id path string true "Meeting ID"
// @Success 200 {object} DeleteMeetingResponse
// @Router /meeting/{id} [delete]
func (mp meetingPresentation) DeleteMeeting(c *gin.Context) {
	operator := c.MustGet("validate_user").(models.User)
	var response DeleteMeetingResponse
	id := c.Param("id")
	if id == "" {
		response.Status = "error"
		response.Message = "Invalid request"
		c.JSON(400, response)
		return
	}

	err := mp.meetingService.DeleteMeeting(operator, id)
	if err != nil {
		response.Status = "error"
		response.Message = err.Error()
		c.JSON(500, response)
		return
	}

	response.Status = "success"
	response.Message = "Meeting deleted"
	c.JSON(200, response)
}

// @Summary Get a meeting
// @Description Get a meeting
// @Tags Meeting
// @Param id path string true "Meeting ID"
// @Success 200 {object} CreateUpdateGetMeetingResponse
// @Router /meeting/{id} [get]
func (mp meetingPresentation) GetMeeting(c *gin.Context) {
	id := c.Param("id")
	var response CreateUpdateGetMeetingResponse
	if id == "" {
		response.Status = "error"
		response.Message = "Invalid request"
		c.JSON(400, response)
		return
	}

	meeting, err := mp.meetingService.GetMeeting(id)
	if err != nil {
		response.Status = "error"
		response.Message = "Internal server error"
		c.JSON(500, response)
		return
	}

	response.Status = "success"
	response.Message = "Meeting retrieved"
	response.Data.Meeting.ID = meeting.ID
	response.Data.Meeting.RoomID = meeting.RoomID
	response.Data.Meeting.Title = meeting.Title
	response.Data.Meeting.Description = meeting.Description
	response.Data.Meeting.OrganizerID = meeting.OrganizerID
	response.Data.Meeting.Participants = meeting.Participants
	response.Data.Meeting.StartTime = meeting.StartTime
	response.Data.Meeting.EndTime = meeting.EndTime
	response.Data.Meeting.StatusType = meeting.StatusType
	c.JSON(200, response)
}

// @Summary Get all meetings
// @Description Get all meetings
// @Tags Meeting
// @Success 200 {object} GetAllMeetingsResponse
// @Router /meeting/getAllMeetings [get]
func (mp meetingPresentation) GetAllMeetings(c *gin.Context) {
	var response GetAllMeetingsResponse
	meetings, err := mp.meetingService.GetAllMeetings()
	if err != nil {
		response.Status = "error"
		response.Message = "Internal server error"
		c.JSON(500, response)
		return
	}

	response.Status = "success"
	response.Message = "Meetings retrieved"
	for _, meeting := range meetings {
		var meetingResponse struct {
			ID           string    `json:"id"`
			RoomID       int       `json:"room_id"`
			Title        string    `json:"title"`
			OrganizerID  uint      `json:"organizer"`
			Participants []uint    `json:"participants"`
			Description  string    `json:"description"`
			StartTime    time.Time `json:"start_time"`
			EndTime      time.Time `json:"end_time"`
			StatusType   string    `json:"status_type"`
		}
		meetingResponse.ID = meeting.ID
		meetingResponse.RoomID = meeting.RoomID
		meetingResponse.Title = meeting.Title
		meetingResponse.Description = meeting.Description
		meetingResponse.OrganizerID = meeting.OrganizerID
		meetingResponse.Participants = meeting.Participants
		meetingResponse.StartTime = meeting.StartTime
		meetingResponse.EndTime = meeting.EndTime
		meetingResponse.StatusType = meeting.StatusType
		response.Data.Meetings = append(response.Data.Meetings, meetingResponse)
	}
	c.JSON(200, response)
}

// GetMeetingsByRoomIdAndDatePeriod retrieves meetings based on room ID and specific date
// @Summary Get meetings by room ID and date
// @Description Get meetings by room ID and date
// @Tags Meeting
// @Param room_id query int true "Room ID"
// @Param date_from query string true "Date"
// @Param date_to query string false "Date"
// @Success 200 {object} GetAllMeetingsResponse
// @Router /meeting/getMeetingsByRoomIdAndDatePeriod [get]
func (mp meetingPresentation) GetMeetingsByRoomIdAndDatePeriod(c *gin.Context) {
	var response GetAllMeetingsResponse
	roomID, success_1 := c.GetQuery("room_id")
	date_from, success_2 := c.GetQuery("date_from")
	date_to, _ := c.GetQuery("date_to")
	if !success_1 || !success_2 {
		response.Status = "error"
		response.Message = "Invalid requestaaa"
		c.JSON(400, response)
		return
	}
	roomIdInt, err1 := strconv.Atoi(roomID)
	parsedDateFrom, err2 := time.Parse("2006-01-02", date_from)
	parsedDateTo, err3 := time.Parse("2006-01-02", date_to)
	if err1 != nil || err2 != nil {
		response.Status = "error"
		response.Message = "Invalid request"
		c.JSON(400, response)
		return
	}
	if err3 != nil {
		parsedDateTo = parsedDateFrom
	}

	meetings, err := mp.meetingService.GetMeetingsByRoomIdAndDatePeriod(roomIdInt, parsedDateFrom, parsedDateTo)
	if err != nil {
		response.Status = "error"
		response.Message = "Internal server error"
		c.JSON(500, response)
		return
	}

	response.Status = "success"
	response.Message = "Meetings retrieved"
	for _, meeting := range meetings {
		var meetingResponse struct {
			ID           string    `json:"id"`
			RoomID       int       `json:"room_id"`
			Title        string    `json:"title"`
			OrganizerID  uint      `json:"organizer"`
			Participants []uint    `json:"participants"`
			Description  string    `json:"description"`
			StartTime    time.Time `json:"start_time"`
			EndTime      time.Time `json:"end_time"`
			StatusType   string    `json:"status_type"`
		}
		meetingResponse.ID = meeting.ID
		meetingResponse.RoomID = meeting.RoomID
		meetingResponse.Title = meeting.Title
		meetingResponse.Description = meeting.Description
		meetingResponse.OrganizerID = meeting.OrganizerID
		meetingResponse.Participants = meeting.Participants
		meetingResponse.StartTime = meeting.StartTime
		meetingResponse.EndTime = meeting.EndTime
		meetingResponse.StatusType = meeting.StatusType
		response.Data.Meetings = append(response.Data.Meetings, meetingResponse)
	}
	if len(response.Data.Meetings) == 0 {
		response.Data.Meetings = []struct {
			ID           string    `json:"id"`
			RoomID       int       `json:"room_id"`
			Title        string    `json:"title"`
			OrganizerID  uint      `json:"organizer"`
			Participants []uint    `json:"participants"`
			Description  string    `json:"description"`
			StartTime    time.Time `json:"start_time"`
			EndTime      time.Time `json:"end_time"`
			StatusType   string    `json:"status_type"`
		}{}
	}
	c.JSON(200, response)
}

// @Summary Get meetings by participant ID
// @Description Get meetings by participant ID
// @Tags Meeting
// @Param id query int true "Participant ID"
// @Success 200 {object} GetAllMeetingsResponse
// @Router /meeting/getMeetingsByParticipantId [get]
func (mp meetingPresentation) GetMeetingsByParticipantId(c *gin.Context) {
	var response GetAllMeetingsResponse
	participantID, success := c.GetQuery("id")
	if !success {
		response.Status = "error"
		response.Message = "Invalid request"
		c.JSON(400, response)
		return
	}
	participantIdInt, err1 := strconv.Atoi(participantID)
	if err1 != nil {
		response.Status = "error"
		response.Message = "Invalid request"
		c.JSON(400, response)
		return
	}
	participantIdUint := uint(participantIdInt)

	meetings, err := mp.meetingService.GetMeetingsByParticipantId(participantIdUint)
	if err != nil {
		response.Status = "error"
		response.Message = "Internal server error"
		c.JSON(500, response)
		return
	}

	response.Status = "success"
	response.Message = "Meetings retrieved"
	for _, meeting := range meetings {
		var meetingResponse struct {
			ID           string    `json:"id"`
			RoomID       int       `json:"room_id"`
			Title        string    `json:"title"`
			OrganizerID  uint      `json:"organizer"`
			Participants []uint    `json:"participants"`
			Description  string    `json:"description"`
			StartTime    time.Time `json:"start_time"`
			EndTime      time.Time `json:"end_time"`
			StatusType   string    `json:"status_type"`
		}
		meetingResponse.ID = meeting.ID
		meetingResponse.RoomID = meeting.RoomID
		meetingResponse.Title = meeting.Title
		meetingResponse.Description = meeting.Description
		meetingResponse.OrganizerID = meeting.OrganizerID
		meetingResponse.Participants = meeting.Participants
		meetingResponse.StartTime = meeting.StartTime
		meetingResponse.EndTime = meeting.EndTime
		meetingResponse.StatusType = meeting.StatusType
		response.Data.Meetings = append(response.Data.Meetings, meetingResponse)
	}
	if len(response.Data.Meetings) == 0 {
		response.Data.Meetings = []struct {
			ID           string    `json:"id"`
			RoomID       int       `json:"room_id"`
			Title        string    `json:"title"`
			OrganizerID  uint      `json:"organizer"`
			Participants []uint    `json:"participants"`
			Description  string    `json:"description"`
			StartTime    time.Time `json:"start_time"`
			EndTime      time.Time `json:"end_time"`
			StatusType   string    `json:"status_type"`
		}{}
	}
	c.JSON(200, response)
}
