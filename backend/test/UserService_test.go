package test

import (
	"errors"
	"meeting-center/src/models"
	"meeting-center/src/services"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserDomain struct {
	mock.Mock
}

func (m *MockUserDomain) CreateUser(user *models.User) (*models.User, error) {
	args := m.Called(user)
	return args.Get(0).(*models.User), nil
}

func (m *MockUserDomain) UpdateUser(user *models.User) (*models.User, error) {
	args := m.Called(user)
	return args.Get(0).(*models.User), nil
}

func (m *MockUserDomain) GetUserByEmail(email string) (*models.User, error) {
	args := m.Called(email)
	return args.Get(0).(*models.User), nil
}

func (m *MockUserDomain) GetAllUsers() ([]models.User, error) {
	args := m.Called()
	return args.Get(0).([]models.User), nil
}

func (m *MockUserDomain) GetUserByID(id uint) (*models.User, error) {
	args := m.Called(id)
	return args.Get(0).(*models.User), nil
}

func TestServiceCreateUser(t *testing.T) {
	// Arrange
	// new user for testing input
	user := &models.User{
		Username: "test-username",
		Email:    "test@test.com",
		Password: "test-password",
		Role:     "test-role",
	}
	// mock the user domain
	mockUserDomain := new(MockUserDomain)
	mockUserDomain.On("GetUserByEmail", user.Email).Return(&models.User{}, errors.New("user not found"))
	mockUserDomain.On("CreateUser", user).Return(user)
	us := services.NewUserService(mockUserDomain)

	// Act
	createdUser, err := us.CreateUser(user)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, user.Username, createdUser.Username)
	assert.Equal(t, user.Email, createdUser.Email)
	assert.Equal(t, user.Password, createdUser.Password)
	assert.Equal(t, user.Role, createdUser.Role)

}

func TestServiceUpdateUser(t *testing.T) {
	// Arrange
	// new user for testing input
	user := &models.User{
		ID:       1,
		Username: "test-username-updated",
		Email:    "test@test.com",
		Password: "test-password-updated",
		Role:     "test-role-updated",
	}
	// mock the user domain
	mockUserDomain := new(MockUserDomain)
	mockUserDomain.On("GetUserByID", user.ID).Return(user, nil)
	mockUserDomain.On("UpdateUser", user).Return(user)
	us := services.NewUserService(mockUserDomain)

	// Act
	updatedUser, err := us.UpdateUser(user, user)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, user.Username, updatedUser.Username)
	assert.Equal(t, user.Email, updatedUser.Email)
	assert.Equal(t, user.Password, updatedUser.Password)
	assert.Equal(t, user.Role, updatedUser.Role)
}
