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
	GetMeetingsByRoomIdAndDate(c *gin.Context)
}

type meetingPresentation struct {
	MeetingService services.MeetingService
}

func NewMeetingPresentation(roomServiceArg ...services.MeetingService) MeetingPresentation {
	if len(roomServiceArg) == 1 {
		return MeetingPresentation(&meetingPresentation{
			MeetingService: roomServiceArg[0],
		})
	} else {
		return MeetingPresentation(&meetingPresentation{
			MeetingService: services.NewMeetingService(),
		})
	}
}

// CreateMeeting handles the creation of a new meeting
func (mp meetingPresentation) CreateMeeting(c *gin.Context) {
	var meeting models.Meeting
	if err := c.BindJSON(&meeting); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	createdMeeting, err := mp.MeetingService.CreateMeeting(&meeting)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(200, createdMeeting)
}

// UpdateMeeting handles the updating of meeting details
func (mp meetingPresentation) UpdateMeeting(c *gin.Context) {
	var meeting models.Meeting
	id := c.Param("id")
	if err := c.BindJSON(&meeting); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	err := mp.MeetingService.UpdateMeeting(id, &meeting)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(200, gin.H{"status": "updated"})
}

// DeleteMeeting handles the deletion of a meeting
func (mp meetingPresentation) DeleteMeeting(c *gin.Context) {
	id := c.Param("id")
	err := mp.MeetingService.DeleteMeeting(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(200, gin.H{"status": "deleted"})
}

// GetMeeting retrieves details of a specific meeting
func (mp meetingPresentation) GetMeeting(c *gin.Context) {
	id := c.Param("id")
	meeting, err := mp.MeetingService.GetMeeting(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(200, meeting)
}

// GetAllMeetings retrieves all meetings
func (mp meetingPresentation) GetAllMeetings(c *gin.Context) {
	meetings, err := mp.MeetingService.GetAllMeetings()
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(200, meetings)
}

// GetMeetingsByRoomIdAndDate retrieves meetings based on room ID and specific date
func (mp meetingPresentation) GetMeetingsByRoomIdAndDate(c *gin.Context) {
	roomID, roomErr := c.GetQuery("room_id")
	date, dateErr := c.GetQuery("date")
	if !roomErr || !dateErr {
		c.JSON(400, gin.H{"error": "Room ID and Date required"})
		return
	}
	roomIdInt, _ := strconv.Atoi(roomID)
	parsedDate, _ := time.Parse("2006-01-02", date)
	fmt.Println(roomIdInt, parsedDate)
	meetings, err := mp.MeetingService.GetMeetingsByRoomIdAndDate(roomIdInt, parsedDate)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(200, meetings)
}
