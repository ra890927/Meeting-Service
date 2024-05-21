package repos

import (
	db "meeting-center/src/io"
	"meeting-center/src/models"

	"golang.org/x/crypto/bcrypt"
)

type UserRepo interface {
	CreateUser(user *models.User) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
}

func CreateUser(user *models.User) (*models.User, error) {
	db := db.GetDBInstance()

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

func UpdateUser(user *models.User) (*models.User, error) {
	db := db.GetDBInstance()

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
