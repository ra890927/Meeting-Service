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

func (m *MockUserDomain) CreateUser(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserDomain) UpdateUser(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserDomain) GetUserByEmail(email string) (models.User, error) {
	args := m.Called(email)
	return args.Get(0).(models.User), nil
}

func (m *MockUserDomain) GetAllUsers() ([]models.User, error) {
	args := m.Called()
	return args.Get(0).([]models.User), nil
}

func (m *MockUserDomain) GetUserByID(id uint) (models.User, error) {
	args := m.Called(id)
	return args.Get(0).(models.User), nil
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
	mockUserDomain.On("GetUserByEmail", user.Email).Return(models.User{}, errors.New("user not found"))
	mockUserDomain.On("CreateUser", user).Return(nil)
	us := services.NewUserService(mockUserDomain)

	// Act
	err := us.CreateUser(user)

	// Assert
	assert.NoError(t, err)
}

func TestServiceUpdateUser(t *testing.T) {
	// Arrange
	// new user for testing input
	operator := models.User{
		Role: "admin",
	}
	user := &models.User{
		ID:       1,
		Username: "test-username-updated",
		Email:    "test@test.com",
		Password: "test-password-updated",
		Role:     "test-role-updated",
	}
	// mock the user domain
	mockUserDomain := new(MockUserDomain)
	mockUserDomain.On("GetUserByID", user.ID).Return(*user, nil)
	mockUserDomain.On("UpdateUser", user).Return(nil)
	us := services.NewUserService(mockUserDomain)

	// Act
	err := us.UpdateUser(operator, user)

	// Assert
	assert.NoError(t, err)
}

func TestServiceGetAllUsers(t *testing.T) {
	// Arrange
	// mock the user domain
	mockUserDomain := new(MockUserDomain)
	mockUserDomain.On("GetAllUsers").Return([]models.User{}, nil)
	us := services.NewUserService(mockUserDomain)

	// Act
	users, err := us.GetAllUsers()

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, users)
}

func TestServiceGetAllUsersError(t *testing.T) {
	// Arrange
	// mock the user domain
	mockUserDomain := new(MockUserDomain)
	mockUserDomain.On("GetAllUsers").Return([]models.User{}, nil)
	us := services.NewUserService(mockUserDomain)

	// Act
	users, err := us.GetAllUsers()

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, users)
}
