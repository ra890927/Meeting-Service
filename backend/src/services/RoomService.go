package services

import (
	"meeting-center/src/domains"
	"meeting-center/src/models"
)

type RoomService interface {
	CreateRoom(room *models.Room) error
	GetRoomByID(id int) (*models.Room, error)
	GetAllRooms() ([]*models.Room, error)
	UpdateRoom(room *models.Room) error
	DeleteRoom(id int) error
}

type roomService struct {
	RoomDomain domains.RoomDomain
}

// NewRoomService constructs a new RoomService. Optionally a specific RoomDomain can be injected.
func NewRoomService(roomDomainArgs ...domains.RoomDomain) RoomService {
	if len(roomDomainArgs) == 1 {
		return RoomService(&roomService{
			RoomDomain: roomDomainArgs[0],
		})
	} else if len(roomDomainArgs) == 0 {
		return RoomService(&roomService{
			RoomDomain: domains.NewRoomDomain(),
		})
	} else {
		panic("Too many arguments")
	}
}

// CreateRoom creates a new room using the RoomDomain
func (rs roomService) CreateRoom(room *models.Room) error {
	err := rs.RoomDomain.CreateRoom(room)
	if err != nil {
		return err
	}
	return nil
}

// UpdateRoom updates an existing room using the RoomDomain
func (rs roomService) UpdateRoom(room *models.Room) error {
	err := rs.RoomDomain.UpdateRoom(room)
	if err != nil {
		return err
	}
	return nil
}

// DeleteRoom deletes a room by its ID using the RoomDomain
func (rs roomService) DeleteRoom(id int) error {
	err := rs.RoomDomain.DeleteRoom(id)
	if err != nil {
		return err
	}
	return nil
}

// GetRoom retrieves a room by its ID using the RoomDomain
func (rs roomService) GetRoomByID(id int) (*models.Room, error) {
	room, err := rs.RoomDomain.GetRoomByID(id)
	if err != nil {
		return nil, err
	}
	return room, nil
}

// GetAllRooms retrieves all rooms using the RoomDomain
func (rs roomService) GetAllRooms() ([]*models.Room, error) {
	rooms, err := rs.RoomDomain.GetAllRooms()
	if err != nil {
		return nil, err
	}
	return rooms, nil
}
