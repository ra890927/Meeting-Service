package repos

import (
	"meeting-center/src/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type RoomRepo interface {
	CreateRoom(room *models.Room) (*models.Room, error)
	UpdateRoom(id string, room *models.Room) error
	DeleteRoom(id string) error
	GetRoom(id string) (*models.Room, error)
	GetAllRooms() ([]*models.Room, error) // 新增方法
}

type roomRepo struct {
	dsn string
}

// NewRoomRepo constructs a new RoomRepo. Optionally a specific DSN can be injected.
func NewRoomRepo(dsnArgs ...string) RoomRepo {
	dsn := "../sqlite.db" // Default DSN
	if len(dsnArgs) == 1 {
		dsn = dsnArgs[0]
	}
	return &roomRepo{
		dsn: dsn,
	}
}

// CreateRoom creates a new room in the database
func (rr roomRepo) CreateRoom(room *models.Room) (*models.Room, error) {
	db, err := gorm.Open(sqlite.Open(rr.dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	result := db.Create(room)
	if result.Error != nil {
		return nil, result.Error
	}
	return room, nil
}

// UpdateRoom updates an existing room in the database
func (rr roomRepo) UpdateRoom(id string, room *models.Room) error {
	db, err := gorm.Open(sqlite.Open(rr.dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	result := db.Model(&models.Room{}).Where("id = ?", id).Updates(room)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteRoom deletes a room from the database
func (rr roomRepo) DeleteRoom(id string) error {
	db, err := gorm.Open(sqlite.Open(rr.dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	result := db.Delete(&models.Room{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetRoom retrieves a room from the database
func (rr roomRepo) GetRoom(id string) (*models.Room, error) {
	db, err := gorm.Open(sqlite.Open(rr.dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	var room models.Room
	result := db.First(&room, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &room, nil
}

// GetAllRooms retrieves all rooms from the database
func (rr roomRepo) GetAllRooms() ([]*models.Room, error) {
	db, err := gorm.Open(sqlite.Open(rr.dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	var rooms []*models.Room
	result := db.Find(&rooms)
	if result.Error != nil {
		return nil, result.Error
	}
	return rooms, nil
}
