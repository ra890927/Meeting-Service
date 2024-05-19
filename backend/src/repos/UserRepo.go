package repos

import (
	"meeting-center/src/models"

	"gorm.io/driver/sqlite"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepo interface {
	CreateUser(user *models.User) (*models.User, error)

	UpdateUser(user *models.User) (*models.User, error)
}

type userRepo struct {
	dsn string
}

func NewUserRepo(dsnArgs ...string) UserRepo {
	if len(dsnArgs) == 1 {
		return UserRepo(&userRepo{dsn: dsnArgs[0]})
	} else if len(dsnArgs) == 0 {
		return UserRepo(&userRepo{dsn: "../sqlite.db"})
	} else {
		panic("too many arguments")
	}
}

func (ur userRepo) CreateUser(user *models.User) (*models.User, error) {
	// Create a new user
	db, err := gorm.Open(sqlite.Open(ur.dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// hash the password of the input user's password
	hashValue, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashValue)

	result := db.Create(user)
	// return the user if no errors
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (ur userRepo) UpdateUser(user *models.User) (*models.User, error) {
	// Update a user
	db, err := gorm.Open(sqlite.Open(ur.dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// hash the password of the input user's password
	hashValue, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashValue)

	result := db.Save(user)
	// return the user if no errors
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
