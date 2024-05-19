package services

import (
	"meeting-center/src/domains"
	"meeting-center/src/models"
	"net/mail"
)

type UserService interface {
	CreateUser(user *models.User) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
}

type userService struct {
	userDomain domains.UserDomain
}

func NewUserService(userDomainArgs ...domains.UserDomain) UserService {
	if len(userDomainArgs) == 1 {
		return UserService(&userService{userDomain: userDomainArgs[0]})
	} else if len(userDomainArgs) == 0 {
		return UserService(&userService{userDomain: domains.NewUserDomain()})
	} else {
		panic("Too many arguments")
	}
}

func (us userService) CreateUser(user *models.User) (*models.User, error) {
	// Validate the email
	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return nil, err
	}

	// Create a new user
	createdUser, err := us.userDomain.CreateUser(user)

	// return the user if no errors
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (us userService) UpdateUser(user *models.User) (*models.User, error) {
	// Validate the email
	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return nil, err
	}

	// Update a user
	updatedUser, err := us.userDomain.UpdateUser(user)

	// return the user if no errors
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}
