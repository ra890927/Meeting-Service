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
	UpdateUser(operator *models.User, user *models.User) (*models.User, error)
	GetAllUsers() ([]models.User, error)
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

	// Check if the user exists
	_, err = us.userDomain.GetUserByEmail(user.Email)
	if err == nil {
		return nil, errors.New("user already exists")
	}

	// Set default role to "user"
	user.Role = "user"

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

func (us userService) GetAllUsers() ([]models.User, error) {
	// Get all users
	users, err := us.userDomain.GetAllUsers()
	if err != nil {
		return nil, errors.New("error getting all users")
	}

	return users, nil
}

func (us userService) UpdateUser(operator *models.User, updatedUser *models.User) (*models.User, error) {
	// check if the operator is the user itself or admin
	if operator.ID != updatedUser.ID && operator.Role != "admin" {
		return nil, errors.New("only user itself or admin can update user")
	}

	// Check if the user exists
	userByID, err := us.userDomain.GetUserByID(updatedUser.ID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Validate the input email
	_, err = mail.ParseAddress(updatedUser.Email)
	if err != nil {
		return nil, err
	}

	// use the original [username, password] if the updated one is empty
	if updatedUser.Username == "" {
		updatedUser.Username = userByID.Username
	}
	if updatedUser.Password == "" {
		updatedUser.Password = userByID.Password
	} else {
		hashValue, err := bcrypt.GenerateFromPassword([]byte(updatedUser.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		updatedUser.Password = string(hashValue)
	}

	// check if the operator is admin and update the role or use the original role
	if updatedUser.Role == "admin" {
		if operator.Role != "admin" {
			return nil, errors.New("only admin can update user role")
		}
	} else {
		updatedUser.Role = userByID.Role
	}

	// Update a user
	_, err = us.userDomain.UpdateUser(updatedUser)

	// return the user if no errors
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}
