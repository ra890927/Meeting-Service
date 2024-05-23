package presentations

import (
	"meeting-center/src/models"
	"meeting-center/src/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RoomPresentation interface {
	CreateRoom(c *gin.Context)
	UpdateRoom(c *gin.Context)
	DeleteRoom(c *gin.Context)
	GetRoomByID(c *gin.Context)
	GetAllRooms(c *gin.Context)
}

type roomPresentation struct {
	RoomService services.RoomService
}

func NewRoomPresentation(roomServiceArgs ...services.RoomService) RoomPresentation {
	if len(roomServiceArgs) == 1 {
		return RoomPresentation(&roomPresentation{
			RoomService: roomServiceArgs[0],
		})
	} else if len(roomServiceArgs) == 0 {
		return RoomPresentation(&roomPresentation{
			RoomService: services.NewRoomService(),
		})
	} else {
		panic("Too many arguments")
	}
}

type CreateRoomInput struct {
	RoomName string `json:"room_name"`
	Type     string `json:"type"`
	Rules    []int  `json:"rules"`
	Capacity int    `json:"capacity"`
}

type CreateRoomResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Room struct {
			ID       int    `json:"id"`
			RoomName string `json:"room_name"`
			Type     string `json:"type"`
			Rules    []int  `json:"rules"`
			Capacity int    `json:"capacity"`
		} `json:"room"`
	} `json:"data"`
}

type UpdateRoomInput struct {
	ID       int    `json:"id"`
	RoomName string `json:"room_name"`
	Type     string `json:"type"`
	Rules    []int  `json:"rules"`
	Capacity int    `json:"capacity" `
}

type UpdateRoomResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Room struct {
			ID       int    `json:"id"`
			RoomName string `json:"room_name"`
			Type     string `json:"type"`
			Rules    []int  `json:"rules"`
			Capacity int    `json:"capacity"`
		} `json:"room"`
	} `json:"data"`
}

type GetAllRoomsResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Room []struct {
			ID       int    `json:"id"`
			RoomName string `json:"room_name"`
			Type     string `json:"type"`
			Rules    []int  `json:"rules"`
			Capacity int    `json:"capacity"`
		} `json:"rooms"`
	} `json:"data"`
}

type GetRoomByIDResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Room struct {
			ID       int    `json:"id"`
			RoomName string `json:"room_name"`
			Type     string `json:"type"`
			Rules    []int  `json:"rules"`
			Capacity int    `json:"capacity"`
		} `json:"room"`
	} `json:"data"`
}

// @Summary Create a new room
// @Description Create a new room
// @Tags admin
// @Accept json
// @Produce json
// @Param room body CreateRoomInput true "Room information"
// @Success 200 {object} CreateRoomResponse
// @Router /admin/room [post]
func (rp *roomPresentation) CreateRoom(c *gin.Context) {
	var createRoomInput CreateRoomInput
	var createRoomResponse CreateRoomResponse
	if err := c.BindJSON(&createRoomInput); err != nil {
		createRoomResponse.Status = "error"
		createRoomResponse.Message = "Invalid request"
		c.JSON(http.StatusBadRequest, createRoomResponse)
		return
	}

	room := models.Room{
		RoomName: createRoomInput.RoomName,
		Type:     createRoomInput.Type,
		Rules:    createRoomInput.Rules,
		Capacity: createRoomInput.Capacity,
	}

	err := rp.RoomService.CreateRoom(&room)
	if err != nil {
		createRoomResponse.Status = "error"
		createRoomResponse.Message = "Failed to create room"
		c.JSON(http.StatusInternalServerError, createRoomResponse)
		return
	}

	createRoomResponse.Status = "success"
	createRoomResponse.Message = "Room created successfully"
	createRoomResponse.Data.Room.ID = room.ID
	createRoomResponse.Data.Room.RoomName = room.RoomName
	createRoomResponse.Data.Room.Type = room.Type
	createRoomResponse.Data.Room.Rules = room.Rules
	createRoomResponse.Data.Room.Capacity = room.Capacity

	c.JSON(http.StatusOK, createRoomResponse)
}

// @Summary Update room information
// @Description Update room information
// @Tags admin
// @Accept json
// @Produce json
// @Param room body UpdateRoomInput true "Room information"
// @Success 200 {object} UpdateRoomResponse
// @Router /admin/room [put]
func (rp *roomPresentation) UpdateRoom(c *gin.Context) {
	var updateRoomInput UpdateRoomInput
	var updateRoomResponse UpdateRoomResponse
	if err := c.BindJSON(&updateRoomInput); err != nil {
		updateRoomResponse.Status = "error"
		updateRoomResponse.Message = "Invalid request"
		c.JSON(http.StatusBadRequest, updateRoomResponse)
		return
	}

	room := models.Room{
		ID:       updateRoomInput.ID,
		RoomName: updateRoomInput.RoomName,
		Type:     updateRoomInput.Type,
		Rules:    updateRoomInput.Rules,
		Capacity: updateRoomInput.Capacity,
	}

	err := rp.RoomService.UpdateRoom(&room)
	if err != nil {
		updateRoomResponse.Status = "error"
		updateRoomResponse.Message = "Failed to update room"
		c.JSON(http.StatusInternalServerError, updateRoomResponse)
		return
	}

	updateRoomResponse.Status = "success"
	updateRoomResponse.Message = "Room updated successfully"
	updateRoomResponse.Data.Room.ID = room.ID
	updateRoomResponse.Data.Room.RoomName = room.RoomName
	updateRoomResponse.Data.Room.Type = room.Type
	updateRoomResponse.Data.Room.Rules = room.Rules
	updateRoomResponse.Data.Room.Capacity = room.Capacity

	c.JSON(http.StatusOK, updateRoomResponse)
}

type DeleteRoomResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// @Summary Delete room
// @Description Delete room by room ID
// @Tags admin
// @Param id path int true "Room ID"
// @Success 200 {string} string "deleted"
// @Router /admin/room/{id} [delete]
func (rp *roomPresentation) DeleteRoom(c *gin.Context) {
	var deleteRoomResponse DeleteRoomResponse
	roomID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		deleteRoomResponse.Status = "error"
		deleteRoomResponse.Message = "Invalid request"
		c.JSON(http.StatusBadRequest, deleteRoomResponse)
		return
	}

	err = rp.RoomService.DeleteRoom(roomID)
	if err != nil {
		deleteRoomResponse.Status = "error"
		deleteRoomResponse.Message = "Failed to delete room"
		c.JSON(http.StatusInternalServerError, deleteRoomResponse)
		return
	}

	deleteRoomResponse.Status = "success"
	deleteRoomResponse.Message = "Room deleted successfully"
	c.JSON(http.StatusOK, deleteRoomResponse)
}

// @Summary Get room by ID
// @Description Get room by ID
// @Tags room
// @Param id path int true "Room ID"
// @Success 200 {object} GetRoomByIDResponse
// @Router /room/{id} [get]
func (rp *roomPresentation) GetRoomByID(c *gin.Context) {
	var getRoomByIDResponse GetRoomByIDResponse
	roomID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		getRoomByIDResponse.Status = "error"
		getRoomByIDResponse.Message = "Invalid request"
		c.JSON(http.StatusBadRequest, getRoomByIDResponse)
		return
	}

	room, err := rp.RoomService.GetRoomByID(roomID)
	if err != nil {
		getRoomByIDResponse.Status = "error"
		getRoomByIDResponse.Message = "Failed to get room"
		c.JSON(http.StatusInternalServerError, getRoomByIDResponse)
		return
	}

	getRoomByIDResponse.Status = "success"
	getRoomByIDResponse.Message = "Room retrieved successfully"
	getRoomByIDResponse.Data.Room.ID = room.ID
	getRoomByIDResponse.Data.Room.RoomName = room.RoomName
	getRoomByIDResponse.Data.Room.Type = room.Type
	getRoomByIDResponse.Data.Room.Rules = room.Rules
	getRoomByIDResponse.Data.Room.Capacity = room.Capacity

	c.JSON(http.StatusOK, getRoomByIDResponse)
}

// @Summary Get all rooms
// @Description Get all rooms
// @Tags room
// @Success 200 {object} GetAllRoomsResponse
// @Router /room/getAllRooms [get]
func (rp *roomPresentation) GetAllRooms(c *gin.Context) {
	rooms, err := rp.RoomService.GetAllRooms()
	var getAllRoomsResponse GetAllRoomsResponse
	if err != nil {
		getAllRoomsResponse.Status = "error"
		getAllRoomsResponse.Message = "Failed to get rooms"
		c.JSON(http.StatusInternalServerError, getAllRoomsResponse)
		return
	}

	getAllRoomsResponse.Status = "success"
	getAllRoomsResponse.Message = "Rooms retrieved successfully"
	for _, room := range rooms {
		getAllRoomsResponse.Data.Room = append(getAllRoomsResponse.Data.Room, struct {
			ID       int    `json:"id"`
			RoomName string `json:"room_name"`
			Type     string `json:"type"`
			Rules    []int  `json:"rules"`
			Capacity int    `json:"capacity"`
		}{
			ID:       room.ID,
			RoomName: room.RoomName,
			Type:     room.Type,
			Rules:    room.Rules,
			Capacity: room.Capacity,
		})
	}

	c.JSON(http.StatusOK, getAllRoomsResponse)
}
