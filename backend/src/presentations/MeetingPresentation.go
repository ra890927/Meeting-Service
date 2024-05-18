package presentations

import (
	"meeting-center/src/models"
	"meeting-center/src/services"

	"github.com/gin-gonic/gin"
)

type MeetingPresentation interface {
	CreateMeeting(c *gin.Context)
	UpdateMeeting(c *gin.Context)
	DeleteMeeting(c *gin.Context)
	GetMeeting(c *gin.Context)
}

type meetingPresentation struct {
	MeetingService services.MeetingService
}

func NewMeetingPresentation(opt ...services.MeetingService) MeetingPresentation {
	if len(opt) == 1 {
		return &meetingPresentation{
			MeetingService: opt[0],
		}
	} else {
		return &meetingPresentation{
			MeetingService: services.NewMeetingService(),
		}
	}
}

// CreateMeeting handles the creation of a new meeting
// @Summary Create a new meeting
// @Description Create a new meeting with details
// @Tags Meeting
// @Accept json
// @Produce json
// @Param meeting body models.Meeting true "Meeting object"
// @Success 200 {object} models.Meeting
// @Router /Meeting [post]
func (mp *meetingPresentation) CreateMeeting(c *gin.Context) {
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
// @Summary Update meeting details
// @Description Update the details of an existing meeting
// @Tags Meeting
// @Accept json
// @Produce json
// @Param id path string true "Meeting ID"
// @Param meeting body models.Meeting true "Meeting object"
// @Success 200 {string} status "updated"
// @Router /Meeting/{id} [put]
func (mp *meetingPresentation) UpdateMeeting(c *gin.Context) {
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
// @Summary Delete a meeting
// @Description Delete a meeting by ID
// @Tags Meeting
// @Accept json
// @Produce json
// @Param id path string true "Meeting ID"
// @Success 200 {string} status "deleted"
// @Router /Meeting/{id} [delete]
func (mp *meetingPresentation) DeleteMeeting(c *gin.Context) {
	id := c.Param("id")
	err := mp.MeetingService.DeleteMeeting(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(200, gin.H{"status": "deleted"})
}

// GetMeeting retrieves details of a specific meeting
// @Summary Get meeting details
// @Description Get details of a meeting by ID
// @Tags Meeting
// @Accept json
// @Produce json
// @Param id path string true "Meeting ID"
// @Success 200 {object} models.Meeting
// @Router /Meeting/{id} [get]
func (mp *meetingPresentation) GetMeeting(c *gin.Context) {
	id := c.Param("id")
	meeting, err := mp.MeetingService.GetMeeting(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(200, meeting)
}
