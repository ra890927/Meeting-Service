package repos

import (
	"meeting-center/src/models"

	"gorm.io/driver/sqlite"

	"gorm.io/gorm"
)

type UserRepo interface {
	CreateUser(user *models.User) (*models.User, error)
}

type userRepo struct {
	dsn string
}

func NewUserRepo(opt ...string) UserRepo {
	dsn := "../sqlite.db"
	if len(opt) == 1 {
		dsn = opt[0]
	}
	return &userRepo{
		dsn: dsn,
	}
}

func (ur *userRepo) CreateUser(user *models.User) (*models.User, error) {
	// Create a new user
	db, err := gorm.Open(sqlite.Open(ur.dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	result := db.Create(user)
	// return the user if no errors
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
