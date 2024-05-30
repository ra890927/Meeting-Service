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
	return ud.userRepo.CreateUser(user)
}

func (ud userDomain) UpdateUser(user *models.User) error {
	// Update a user
	return ud.userRepo.UpdateUser(user)
}

func (ud userDomain) GetAllUsers() ([]models.User, error) {
	// Get all users
	return ud.userRepo.GetAllUsers()
}

func (ud userDomain) GetUserByEmail(email string) (models.User, error) {
	// Get a user by email
	return ud.userRepo.GetUserByEmail(email)
}

func (ud userDomain) GetUserByID(id uint) (models.User, error) {
	// Get a user by ID
	return ud.userRepo.GetUserByID(id)
}
