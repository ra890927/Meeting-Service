package domains

import (
	"meeting-center/src/models"
	"meeting-center/src/repos"
)

type UserDomain interface {
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	GetAllUsers() ([]models.User, error)
	GetUserByEmail(email string) (models.User, error)
	GetUserByID(id uint) (models.User, error)
}

type userDomain struct {
	userRepo repos.UserRepo
}

func NewUserDomain(userRepoArgs ...repos.UserRepo) UserDomain {
	if len(userRepoArgs) == 1 {
		return UserDomain(&userDomain{userRepo: userRepoArgs[0]})
	} else if len(userRepoArgs) == 0 {
		return UserDomain(&userDomain{userRepo: repos.NewUserRepo()})
	} else {
		panic("Too many arguments")
	}
}

func (ud userDomain) CreateUser(user *models.User) error {
	// Create a new user
	err := ud.userRepo.CreateUser(user)

	// return the user if no errors
	if err != nil {
		return err
	}

	return nil
}

func (ud userDomain) UpdateUser(user *models.User) error {
	// Update a user
	err := ud.userRepo.UpdateUser(user)

	// return the user if no errors
	if err != nil {
		return err
	}

	return nil
}

func (ud userDomain) GetAllUsers() ([]models.User, error) {
	// Get all users
	allUsers, err := ud.userRepo.GetAllUsers()
	if err != nil {
		return []models.User{}, err
	}

	return allUsers, nil
}

func (ud userDomain) GetUserByEmail(email string) (models.User, error) {
	// Get a user by email
	userByEmail, err := ud.userRepo.GetUserByEmail(email)

	// return the user if no errors
	if err != nil {
		return models.User{}, err
	}

	return userByEmail, nil
}

func (ud userDomain) GetUserByID(id uint) (models.User, error) {
	// Get a user by ID
	user, err := ud.userRepo.GetUserByID(id)

	// return the user if no errors
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
