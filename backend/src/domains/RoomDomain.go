package domains

import (
	"meeting-center/src/models"
	"meeting-center/src/repos"
)

type RoomDomain interface {
	CreateRoom(room *models.Room) (*models.Room, error)
	UpdateRoom(id string, room *models.Room) error
	DeleteRoom(id string) error
	GetRoom(id string) (*models.Room, error)
}

type roomDomain struct {
	RoomRepo repos.RoomRepo
}

// NewRoomDomain constructs a new RoomDomain. Optionally a specific RoomRepo can be injected.
func NewRoomDomain(opt ...repos.RoomRepo) RoomDomain {
	if len(opt) == 1 {
		return &roomDomain{
			RoomRepo: opt[0],
		}
	} else {
		return &roomDomain{
			RoomRepo: repos.NewRoomRepo(),
		}
	}
}

// CreateRoom creates a new room using the RoomRepo
func (rd *roomDomain) CreateRoom(room *models.Room) (*models.Room, error) {
	createdRoom, err := rd.RoomRepo.CreateRoom(room)
	if err != nil {
		return nil, err
	}
	return createdRoom, nil
}

// UpdateRoom updates an existing room using the RoomRepo
func (rd *roomDomain) UpdateRoom(id string, room *models.Room) error {
	err := rd.RoomRepo.UpdateRoom(id, room)
	if err != nil {
		return err
	}
	return nil
}

// DeleteRoom deletes a room by its ID using the RoomRepo
func (rd *roomDomain) DeleteRoom(id string) error {
	err := rd.RoomRepo.DeleteRoom(id)
	if err != nil {
		return err
	}
	return nil
}

// GetRoom retrieves a room by its ID using the RoomRepo
func (rd *roomDomain) GetRoom(id string) (*models.Room, error) {
	room, err := rd.RoomRepo.GetRoom(id)
	if err != nil {
		return nil, err
	}
	return room, nil
}
