package test

import (
	"meeting-center/src/domains"
	"meeting-center/src/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) CreateUser(user *models.User) (*models.User, error) {
	args := m.Called(user)
	return args.Get(0).(*models.User), nil
}

func TestDomainCreateUser(t *testing.T) {
	// Arrange
	// new user for testing input
	user := &models.User{
		Username: "test",
		Email:    "qwe",
		Password: "password",
		Role:     "user",
	}
	// mock the user repo
	mockUserRepo := new(MockUserRepo)
	mockUserRepo.On("CreateUser", user).Return(user, nil)
	ud := domains.NewUserDomain(mockUserRepo)
	createdUser, err := ud.CreateUser(user)

	// Check if the user is created
	if err != nil {
		t.Error("Error while creating a user")
	} else {
		// to check if the attribute is right(ID, Username, Email, Password, Role)
		assert.Equal(t, user.Username, createdUser.Username)
		assert.Equal(t, user.Email, createdUser.Email)
		assert.Equal(t, user.Password, createdUser.Password)
		assert.Equal(t, user.Role, createdUser.Role)
	}
}
