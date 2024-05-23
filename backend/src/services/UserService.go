package services

import (
	"errors"
	"meeting-center/src/domains"
	"meeting-center/src/models"
	"net/mail"

	"golang.org/x/crypto/bcrypt"
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

	// Check if the user exists (by email)
	_, err = us.userDomain.GetUserByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	// Hash the password
	hashValue, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashValue)

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

	// Check if the user exists
	userByEmail, err := us.userDomain.GetUserByEmail(user.Email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// patch userByEmail with the new user (where the input is not empty)
	if user.Username != "" {
		userByEmail.Username = user.Username
	}
	if user.Password != "" {
		userByEmail.Password = user.Password
	} else {
		hashValue, err := bcrypt.GenerateFromPassword([]byte(userByEmail.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		userByEmail.Password = string(hashValue)
	}

	// Update a user
	updatedUser, err := us.userDomain.UpdateUser(user)

	// return the user if no errors
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}
