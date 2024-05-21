package repos

import (
	db "meeting-center/src/io"
	"meeting-center/src/models"
)

type RoomRepo interface {
	CreateRoom(room *models.Room) (*models.Room, error)
	UpdateRoom(id string, room *models.Room) error
	DeleteRoom(id string) error
	GetRoom(id string) (*models.Room, error)
	GetAllRooms() ([]*models.Room, error) // 新增方法
}

// CreateRoom creates a new room in the database
func CreateRoom(room *models.Room) (*models.Room, error) {
	db := db.GetDBInstance()
	result := db.Create(room)
	if result.Error != nil {
		return nil, result.Error
	}
	return room, nil
}

// UpdateRoom updates an existing room in the database
func UpdateRoom(id string, room *models.Room) error {
	db := db.GetDBInstance()
	result := db.Model(&models.Room{}).Where("id = ?", id).Updates(room)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteRoom deletes a room from the database
func DeleteRoom(id string) error {
	db := db.GetDBInstance()
	result := db.Delete(&models.Room{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetRoom retrieves a room from the database
func GetRoom(id string) (*models.Room, error) {
	db := db.GetDBInstance()
	var room models.Room
	result := db.First(&room, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &room, nil
}

// GetAllRooms retrieves all rooms from the database
func GetAllRooms() ([]*models.Room, error) {
	db := db.GetDBInstance()
	var rooms []*models.Room
	result := db.Find(&rooms)
	if result.Error != nil {
		return nil, result.Error
	}
	return rooms, nil
}
