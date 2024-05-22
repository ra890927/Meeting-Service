package repos

import (
	db "meeting-center/src/io"
	"meeting-center/src/models"

	"gorm.io/gorm"
)

type roomRepo struct {
	db *gorm.DB
}

type RoomRepo interface {
	CreateRoom(room *models.Room) (*models.Room, error)
	UpdateRoom(id string, room *models.Room) error
	DeleteRoom(id string) error
	GetRoom(id string) (*models.Room, error)
	GetAllRooms() ([]*models.Room, error) // 新增方法
}

func NewRoomRepo(dbArgs ...*gorm.DB) RoomRepo {
	if len(dbArgs) == 0 {
		return roomRepo{db: db.GetDBInstance()}
	} else if len(dbArgs) == 1 {
		return roomRepo{db: dbArgs[0]}
	} else {
		panic("Too many arguments")
	}
}

// CreateRoom creates a new room in the database
func (rr roomRepo) CreateRoom(room *models.Room) (*models.Room, error) {
	result := rr.db.Create(room)
	if result.Error != nil {
		return nil, result.Error
	}
	return room, nil
}

// UpdateRoom updates an existing room in the database
func (rr roomRepo) UpdateRoom(id string, room *models.Room) error {
	result := rr.db.Model(&models.Room{}).Where("id = ?", id).Updates(room)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteRoom deletes a room from the database
func (rr roomRepo) DeleteRoom(id string) error {
	result := rr.db.Delete(&models.Room{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetRoom retrieves a room from the database
func (rr roomRepo) GetRoom(id string) (*models.Room, error) {
	var room models.Room
	result := rr.db.First(&room, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &room, nil
}

// GetAllRooms retrieves all rooms from the database
func (rr roomRepo) GetAllRooms() ([]*models.Room, error) {
	var rooms []*models.Room
	result := rr.db.Find(&rooms)
	if result.Error != nil {
		return nil, result.Error
	}
	return rooms, nil
}
