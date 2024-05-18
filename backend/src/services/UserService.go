package services

import (
	"meeting-center/src/domains"
	"meeting-center/src/models"
)

type UserService interface {
	CreateUser(user *models.User) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
}

type userService struct {
	userDomain domains.UserDomain
}

func NewUserService(userDomainArgs ...domains.UserDomain) userService {
	if len(userDomainArgs) == 1 {
		return userService{
			userDomain: userDomainArgs[0],
		}
	} else if len(userDomainArgs) == 0 {
		return userService{
			userDomain: domains.NewUserDomain(),
		}
	} else {
		panic("Too many arguments")
	}
}

func (us userService) CreateUser(user *models.User) (*models.User, error) {
	// Create a new user
	createdUser, err := us.userDomain.CreateUser(user)

	// return the user if no errors
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (us userService) UpdateUser(user *models.User) (*models.User, error) {
	// Update a user
	updatedUser, err := us.userDomain.UpdateUser(user)

	// return the user if no errors
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}
