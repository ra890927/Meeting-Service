package services

import (
	"meeting-center/src/domains"
	"meeting-center/src/models"
)

type UserService interface {
	CreateUser(user *models.User) (*models.User, error)
}

type userService struct {
	UserDomain domains.UserDomain
}

func NewUserService(opt ...domains.UserDomain) UserService {
	if len(opt) == 1 {
		return &userService{
			UserDomain: opt[0],
		}
	} else {
		return &userService{
			UserDomain: domains.NewUserDomain(),
		}
	}
}

func (us *userService) CreateUser(user *models.User) (*models.User, error) {
	// Create a new user
	createdUser, err := us.UserDomain.CreateUser(user)

	// return the user if no errors
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
