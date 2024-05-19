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

func (m *MockUserRepo) UpdateUser(user *models.User) (*models.User, error) {
	args := m.Called(user)
	return args.Get(0).(*models.User), nil
}

func TestDomainCreateUser(t *testing.T) {
	// Arrange
	// new user for testing input
	user := &models.User{
		Username: "test",
		Email:    "test@test.com",
		Password: "password",
		Role:     "user",
	}
	// mock the user repo
	mockUserRepo := new(MockUserRepo)
	mockUserRepo.On("CreateUser", user).Return(user, nil)
	ud := domains.NewUserDomain(mockUserRepo)

	// Act
	createdUser, err := ud.CreateUser(user)

	// Assert
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

func TestDomainUpdateUser(t *testing.T) {
	// Arrange
	// new user for testing input
	user := &models.User{
		Username: "test-username-updated",
		Email:    "test@test.com",
		Password: "test-password-updated",
		Role:     "test-role-updated",
	}
	// mock the user repo
	mockUserRepo := new(MockUserRepo)
	mockUserRepo.On("UpdateUser", user).Return(user, nil)
	ud := domains.NewUserDomain(mockUserRepo)

	// Act
	updatedUser, err := ud.UpdateUser(user)

	// Assert
	if err != nil {
		t.Error("Error while updating a user")
	} else {
		// to check if the attribute is right(ID, Username, Email, Password, Role)
		assert.Equal(t, user.Username, updatedUser.Username)
		assert.Equal(t, user.Email, updatedUser.Email)
		assert.Equal(t, user.Password, updatedUser.Password)
		assert.Equal(t, user.Role, updatedUser.Role)
	}
}
