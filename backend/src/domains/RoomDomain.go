package domains

import (
	"meeting-center/src/models"
	"meeting-center/src/repos"
)

type RoomDomain interface {
	CreateRoom(room *models.Room) error
	GetRoomByID(id int) (*models.Room, error)
	GetAllRooms() ([]*models.Room, error)
	UpdateRoom(room *models.Room) error
	DeleteRoom(id int) error
}

type roomDomain struct {
	roomRepo repos.RoomRepo
}

// NewRoomDomain constructs a new RoomDomain. Optionally a specific RoomRepo can be injected.
func NewRoomDomain(roomRepoArgs ...repos.RoomRepo) RoomDomain {
	if len(roomRepoArgs) == 1 {
		return RoomDomain(&roomDomain{roomRepo: roomRepoArgs[0]})
	} else if len(roomRepoArgs) == 0 {
		return RoomDomain(&roomDomain{roomRepo: repos.NewRoomRepo()})
	} else {
		panic("Too many arguments")
	}
}

// CreateRoom creates a new room using the RoomRepo
func (rd roomDomain) CreateRoom(room *models.Room) error {
	err := rd.roomRepo.CreateRoom(room)
	if err != nil {
		return err
	}
	return nil
}

// UpdateRoom updates an existing room using the RoomRepo
func (rd roomDomain) UpdateRoom(room *models.Room) error {
	err := rd.roomRepo.UpdateRoom(room)
	if err != nil {
		return err
	}
	return nil
}

// DeleteRoom deletes a room by its ID using the RoomRepo
func (rd roomDomain) DeleteRoom(id int) error {
	err := rd.roomRepo.DeleteRoom(id)
	if err != nil {
		return err
	}
	return nil
}

// GetRoomByID retrieves a room by its ID using the RoomRepo
func (rd roomDomain) GetRoomByID(id int) (*models.Room, error) {
	room, err := rd.roomRepo.GetRoomByID(id)
	if err != nil {
		return nil, err
	}
	return room, nil
}

func (rd roomDomain) GetAllRooms() ([]*models.Room, error) {
	rooms, err := rd.roomRepo.GetAllRooms()
	if err != nil {
		return nil, err
	}
	if rooms == nil {
		rooms = []*models.Room{}
	}
	return rooms, nil
}
