package repos

import (
	db "meeting-center/src/io"
	"meeting-center/src/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

type UserRepo interface {
	CreateUser(user *models.User) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
}

func NewUserRepo(dbArgs ...*gorm.DB) UserRepo {
	if len(dbArgs) == 0 {
		return userRepo{db: db.GetDBInstance()}
	} else if len(dbArgs) == 1 {
		return userRepo{db: dbArgs[0]}
	} else {
		panic("Too many arguments")
	}
}

func (ur userRepo) CreateUser(user *models.User) (*models.User, error) {
	// hash the password of the input user's password
	hashValue, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashValue)

	result := ur.db.Create(user)
	// return the user if no errors
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (ur userRepo) UpdateUser(user *models.User) (*models.User, error) {
	// hash the password of the input user's password
	hashValue, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashValue)

	result := ur.db.Save(user)
	// return the user if no errors
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
