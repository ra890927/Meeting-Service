package services

import (
	"meeting-center/src/domains"
	"meeting-center/src/models"
)

type RoomService interface {
	CreateRoom(room *models.Room) (*models.Room, error)
	GetRoom(id string) (*models.Room, error)
	GetAllRooms() ([]*models.Room, error)
	UpdateRoom(id string, room *models.Room) error
	DeleteRoom(id string) error
}

type roomService struct {
	RoomDomain domains.RoomDomain
}

// NewRoomService constructs a new RoomService. Optionally a specific RoomDomain can be injected.
func NewRoomService(roomDomainArgs ...domains.RoomDomain) RoomService {
	if len(roomDomainArgs) == 1 {
		return &roomService{
			RoomDomain: roomDomainArgs[0],
		}
	} else if len(roomDomainArgs) == 0 {
		return &roomService{
			RoomDomain: domains.NewRoomDomain(),
		}
	} else {
		panic("Too many arguments")
	}
}

// CreateRoom creates a new room using the RoomDomain
func (rs roomService) CreateRoom(room *models.Room) (*models.Room, error) {
	createdRoom, err := rs.RoomDomain.CreateRoom(room)
	if err != nil {
		return nil, err
	}
	return createdRoom, nil
}

// UpdateRoom updates an existing room using the RoomDomain
func (rs roomService) UpdateRoom(id string, room *models.Room) error {
	err := rs.RoomDomain.UpdateRoom(id, room)
	if err != nil {
		return err
	}
	return nil
}

// DeleteRoom deletes a room by its ID using the RoomDomain
func (rs roomService) DeleteRoom(id string) error {
	err := rs.RoomDomain.DeleteRoom(id)
	if err != nil {
		return err
	}
	return nil
}

// GetRoom retrieves a room by its ID using the RoomDomain
func (rs roomService) GetRoom(id string) (*models.Room, error) {
	room, err := rs.RoomDomain.GetRoom(id)
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
