package presentations

import (
	"meeting-center/src/models"
	"meeting-center/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RoomPresentation 定義了房間的展示層接口
type RoomPresentation interface {
	CreateRoom(c *gin.Context)
	UpdateRoom(c *gin.Context)
	DeleteRoom(c *gin.Context)
	GetRoom(c *gin.Context)
	GetAllRooms(c *gin.Context)
}

type roomPresentation struct {
	RoomService services.RoomService
}

// NewRoomPresentation 創建一個新的 RoomPresentation 實例
func NewRoomPresentation(roomServiceArgs ...services.RoomService) RoomPresentation {
	if len(roomServiceArgs) == 1 {
		return &roomPresentation{
			RoomService: roomServiceArgs[0],
		}
	} else if len(roomServiceArgs) == 0 {
		return &roomPresentation{
			RoomService: services.NewRoomService(),
		}
	} else {
		panic("Too many arguments")
	}
}

// CreateRoom 處理創建房間的請求
// @Summary 創建房間
// @Description 創建一個新的房間
// @Tags rooms
// @Accept json
// @Produce json
// @Param room body models.Room true "房間信息"
// @Success 200 {object} models.Room
// @Router /rooms [post]
func (rp *roomPresentation) CreateRoom(c *gin.Context) {
	var room models.Room
	if err := c.BindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	createdRoom, err := rp.RoomService.CreateRoom(&room)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, createdRoom)
}

// UpdateRoom 處理更新房間的請求
// @Summary 更新房間
// @Description 更新房間信息
// @Tags rooms
// @Accept json
// @Produce json
// @Param id path int true "房間ID"
// @Param room body models.Room true "房間信息"
// @Success 200 {object} models.Room
// @Router /rooms/{id} [put]
func (rp *roomPresentation) UpdateRoom(c *gin.Context) {
	var room models.Room
	id := c.Param("id")
	if err := c.BindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	err := rp.RoomService.UpdateRoom(id, &room)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}

// DeleteRoom 處理刪除房間的請求
// @Summary 刪除房間
// @Description 刪除指定ID的房間
// @Tags rooms
// @Param id path int true "房間ID"
// @Success 200 {string} string "deleted"
// @Router /rooms/{id} [delete]
func (rp *roomPresentation) DeleteRoom(c *gin.Context) {
	id := c.Param("id")
	err := rp.RoomService.DeleteRoom(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}

// GetRoom 獲取指定ID的房間信息
// @Summary 獲取房間
// @Description 獲取指定ID的房間信息
// @Tags rooms
// @Param id path int true "房間ID"
// @Success 200 {object} models.Room
// @Router /rooms/{id} [get]
func (rp *roomPresentation) GetRoom(c *gin.Context) {
	id := c.Param("id")
	room, err := rp.RoomService.GetRoom(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, room)
}

// GetAllRooms 獲取所有房間信息
// @Summary 獲取所有房間
// @Description 獲取所有房間信息
// @Tags rooms
// @Success 200 {array} models.Room
// @Router /rooms [get]
func (rp *roomPresentation) GetAllRooms(c *gin.Context) {
	rooms, err := rp.RoomService.GetAllRooms()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, rooms)
}
