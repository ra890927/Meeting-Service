package services

import (
	"errors"
	"meeting-center/src/domains"
	"meeting-center/src/models"
	"meeting-center/src/utils"
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(user *models.User) error
	UpdateUser(operator models.User, user *models.User) error
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

func (us userService) CreateUser(user *models.User) error {
	// Validate the email
	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return errors.New("invalid email")
	}

	// Check if the user exists
	existingUser, _ := us.userDomain.GetUserByEmail(user.Email)
	if existingUser.ID != 0 {
		return errors.New("user already exists")
	}

	// Hash the password
	hashValue, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("error when hashing password")
	}

	// Set default role to "user"
	user.Role = "user"
	user.Password = string(hashValue)

	// Create a new user
	return us.userDomain.CreateUser(user)
}

func (us userService) GetAllUsers() ([]models.User, error) {
	// Get all users
	return us.userDomain.GetAllUsers()
}

func (us userService) UpdateUser(operator models.User, updatedUser *models.User) error {
	// check if the operator is the user itself or admin
	if operator.ID != updatedUser.ID && operator.Role != "admin" {
		return errors.New("only user itself or admin can update user")
	}

	// Check if the user exists
	userByID, err := us.userDomain.GetUserByID(updatedUser.ID)
	if err != nil {
		return errors.New("user not found")
	}

	// 0528: not allow to update email
	if updatedUser.Email != userByID.Email {
		return errors.New("email cannot be updated")
	}

	// check if the operator is admin and update the role or use the original role
	if updatedUser.Role != userByID.Role { // if the role is updated
		if operator.Role != "admin" { // only admin can update user role
			return errors.New("only admin can update user role")
		}
	}

	if len(updatedUser.Password) != 0 {
		// if updatedUser.Password have not be overwritten by OverwriteValue
		hashValue, err := bcrypt.GenerateFromPassword([]byte(updatedUser.Password), bcrypt.DefaultCost)
		if err != nil {
			return errors.New("error when hashing password")
		}
		updatedUser.Password = string(hashValue)
	}

	utils.OverwriteValue(updatedUser, &userByID)
	return us.userDomain.UpdateUser(updatedUser)
}
