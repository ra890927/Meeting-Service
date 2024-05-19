package services

import (
	"meeting-center/src/domains"
	"meeting-center/src/models"
)

type AuthService interface {
	Login(user *models.User) (*models.User, *string, error)
	Logout(token *string) error
}

type authService struct {
	authDomain domains.AuthDomain
}

func NewAuthService(authDomainArgs ...domains.AuthDomain) AuthService {
	if len(authDomainArgs) == 1 {
		return AuthService(&authService{authDomain: authDomainArgs[0]})
	} else if len(authDomainArgs) == 0 {
		return AuthService(&authService{authDomain: domains.NewAuthDomain()})
	} else {
		panic("Too many arguments")
	}
}

func (as authService) Login(user *models.User) (*models.User, *string, error) {
	// Login a user
	loggedUser, token, err := as.authDomain.Login(user)

	// return the user if no errors
	if err != nil {
		return nil, nil, err
	}

	return loggedUser, token, nil
}

func (as authService) Logout(token *string) error {
	// Logout a user
	err := as.authDomain.Logout(token)

	// return the user if no errors
	if err != nil {
		return err
	}

	return nil
}
