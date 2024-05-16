package presentations

import (
	"meeting-center/src/models"
	"meeting-center/src/services"

	"github.com/gin-gonic/gin"
)

type RoomPresentation interface {
	CreateRoom(c *gin.Context)
	UpdateRoom(c *gin.Context)
	DeleteRoom(c *gin.Context)
	GetRoom(c *gin.Context)
}

type roomPresentation struct {
	RoomService services.RoomService
}

func NewRoomPresentation(opt ...services.RoomService) RoomPresentation {
	if len(opt) == 1 {
		return &roomPresentation{
			RoomService: opt[0],
		}
	} else {
		return &roomPresentation{
			RoomService: services.NewRoomService(),
		}
	}
}

// CreateRoom handles the creation of a new room
// @Summary Create a new room
// @Description Create a new room with details
// @Tags Room
// @Accept json
// @Produce json
// @Param room body models.Room true "Room object"
// @Success 200 {object} models.Room
// @Router /Room [post]
func (rp *roomPresentation) CreateRoom(c *gin.Context) {
	var room models.Room
	if err := c.BindJSON(&room); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	createdRoom, err := rp.RoomService.CreateRoom(&room)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(200, createdRoom)
}

// UpdateRoom handles the updating of room details
// @Summary Update room details
// @Description Update the details of an existing room
// @Tags Room
// @Accept json
// @Produce json
// @Param id path int true "Room ID"
// @Param room body models.Room true "Room object"
// @Success 200 {string} status "updated"
// @Router /Room/{id} [put]
func (rp *roomPresentation) UpdateRoom(c *gin.Context) {
	var room models.Room
	id := c.Param("id")
	if err := c.BindJSON(&room); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	err := rp.RoomService.UpdateRoom(id, &room)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(200, gin.H{"status": "updated"})
}

// DeleteRoom handles the deletion of a room
// @Summary Delete a room
// @Description Delete a room by ID
// @Tags Room
// @Accept json
// @Produce json
// @Param id path int true "Room ID"
// @Success 200 {string} status "deleted"
// @Router /Room/{id} [delete]
func (rp *roomPresentation) DeleteRoom(c *gin.Context) {
	id := c.Param("id")
	err := rp.RoomService.DeleteRoom(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(200, gin.H{"status": "deleted"})
}

// GetRoom retrieves details of a specific room
// @Summary Get room details
// @Description Get details of a room by ID
// @Tags Room
// @Accept json
// @Produce json
// @Param id path int true "Room ID"
// @Success 200 {object} models.Room
// @Router /Room/{id} [get]
func (rp *roomPresentation) GetRoom(c *gin.Context) {
	id := c.Param("id")
	room, err := rp.RoomService.GetRoom(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(200, room)
}
