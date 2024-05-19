package domains

import (
	"meeting-center/src/models"
	"meeting-center/src/repos"
)

type AuthDomain interface {
	Login(user *models.User) (*models.User, *string, error)
	Logout(token *string) error
}

type authDomain struct {
	authRepo repos.AuthRepo
}

func NewAuthDomain(authRepoArgs ...repos.AuthRepo) AuthDomain {
	if len(authRepoArgs) == 1 {
		return AuthDomain(&authDomain{authRepo: authRepoArgs[0]})
	} else if len(authRepoArgs) == 0 {
		return AuthDomain(&authDomain{authRepo: repos.NewAuthRepo()})
	} else {
		panic("Too many arguments")
	}
}

func (ad authDomain) Login(user *models.User) (*models.User, *string, error) {
	// Login a user
	loggedUser, token, err := ad.authRepo.Login(user)

	// return the user if no errors
	if err != nil {
		return nil, nil, err
	}

	return loggedUser, token, nil
}

func (ad authDomain) Logout(token *string) error {
	// Logout a user
	err := ad.authRepo.Logout(token)

	// return the user if no errors
	if err != nil {
		return err
	}

	return nil
}
