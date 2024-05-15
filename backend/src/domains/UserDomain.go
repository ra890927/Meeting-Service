package domains

import (
	"meeting-center/src/models"
	"meeting-center/src/repos"
)

type UserDomain interface {
	CreateUser(user *models.User) (*models.User, error)
}

type userDomain struct {
	UserRepo repos.UserRepo
}

func NewUserDomain(opt ...repos.UserRepo) UserDomain {
	if len(opt) == 1 {
		return &userDomain{
			UserRepo: opt[0],
		}
	} else {
		return &userDomain{
			UserRepo: repos.NewUserRepo(),
		}
	}
}

func (ud *userDomain) CreateUser(user *models.User) (*models.User, error) {
	// Create a new user
	createdUser, err := ud.UserRepo.CreateUser(user)

	// return the user if no errors
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
