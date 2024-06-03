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
	return args.Get(0).(models.User), args.Error(1)
}

func (m *MockUserDomain) GetAllUsers() ([]models.User, error) {
	args := m.Called()
	return args.Get(0).([]models.User), args.Error(1)
}

func (m *MockUserDomain) GetUserByID(id uint) (models.User, error) {
	args := m.Called(id)
	return args.Get(0).(models.User), args.Error(1)
}

func TestServiceCreateUser(t *testing.T) {
	// Arrange
	user := &models.User{
		Username: "test-username",
		Email:    "test@test.com",
		Password: "test-password",
		Role:     "test-role",
	}
	mockUserDomain := new(MockUserDomain)
	mockUserDomain.On("GetUserByEmail", user.Email).Return(models.User{}, nil)
	mockUserDomain.On("CreateUser", user).Return(nil)
	us := services.NewUserService(mockUserDomain)

	// Act
	err := us.CreateUser(user)

	// Assert
	assert.NoError(t, err)
}

func TestServiceCreateUser_InvalidEmail(t *testing.T) {
	// Arrange
	user := &models.User{
		Email: "invalid-email",
	}
	mockUserDomain := new(MockUserDomain)
	us := services.NewUserService(mockUserDomain)

	// Act
	err := us.CreateUser(user)

	// Assert
	assert.EqualError(t, err, "invalid email")
}

func TestServiceCreateUser_UserAlreadyExists(t *testing.T) {
	// Arrange
	user := &models.User{
		Email: "test@test.com",
	}
	existingUser := models.User{
		ID: 1,
	}
	mockUserDomain := new(MockUserDomain)
	mockUserDomain.On("GetUserByEmail", user.Email).Return(existingUser, nil)
	us := services.NewUserService(mockUserDomain)

	// Act
	err := us.CreateUser(user)

	// Assert
	assert.EqualError(t, err, "user already exists")
}

func TestServiceUpdateUser(t *testing.T) {
	// Arrange
	operator := models.User{
		ID:   2,
		Role: "admin",
	}
	user := &models.User{
		ID:       1,
		Username: "test-username-updated",
		Email:    "test@test.com",
		Password: "test-password-updated",
		Role:     "test-role-updated",
	}
	userByID := models.User{
		ID:       1,
		Email:    "test@test.com",
		Username: "old-username",
		Password: "old-password",
		Role:     "user",
	}
	mockUserDomain := new(MockUserDomain)
	mockUserDomain.On("GetUserByID", user.ID).Return(userByID, nil)
	mockUserDomain.On("UpdateUser", user).Return(nil)
	us := services.NewUserService(mockUserDomain)

	// Act
	err := us.UpdateUser(operator, user)

	// Assert
	assert.NoError(t, err)
}

func TestServiceUpdateUser_NotAdminOrSelf(t *testing.T) {
	// Arrange
	operator := models.User{
		ID:   2,
		Role: "user",
	}
	user := &models.User{
		ID: 1,
	}
	mockUserDomain := new(MockUserDomain)
	us := services.NewUserService(mockUserDomain)

	// Act
	err := us.UpdateUser(operator, user)

	// Assert
	assert.EqualError(t, err, "only user itself or admin can update user")
}

func TestServiceUpdateUser_UserNotFound(t *testing.T) {
	// Arrange
	operator := models.User{
		ID:   1,
		Role: "user",
	}
	user := &models.User{
		ID: 1,
	}
	mockUserDomain := new(MockUserDomain)
	mockUserDomain.On("GetUserByID", user.ID).Return(models.User{}, errors.New("user not found"))
	us := services.NewUserService(mockUserDomain)

	// Act
	err := us.UpdateUser(operator, user)

	// Assert
	assert.EqualError(t, err, "user not found")
}

func TestServiceUpdateUser_EmailCannotBeUpdated(t *testing.T) {
	// Arrange
	operator := models.User{
		ID:   1,
		Role: "user",
	}
	user := &models.User{
		ID:    1,
		Email: "new-email@test.com",
	}
	userByID := models.User{
		ID:    1,
		Email: "test@test.com",
	}
	mockUserDomain := new(MockUserDomain)
	mockUserDomain.On("GetUserByID", user.ID).Return(userByID, nil)
	us := services.NewUserService(mockUserDomain)

	// Act
	err := us.UpdateUser(operator, user)

	// Assert
	assert.EqualError(t, err, "email cannot be updated")
}

func TestServiceUpdateUser_OnlyAdminCanUpdateUserRole(t *testing.T) {
	// Arrange
	operator := models.User{
		ID:   1,
		Role: "user",
	}
	user := &models.User{
		ID:   1,
		Role: "admin",
	}
	userByID := models.User{
		ID:   1,
		Role: "user",
	}
	mockUserDomain := new(MockUserDomain)
	mockUserDomain.On("GetUserByID", user.ID).Return(userByID, nil)
	us := services.NewUserService(mockUserDomain)

	// Act
	err := us.UpdateUser(operator, user)

	// Assert
	assert.EqualError(t, err, "only admin can update user role")
}

func TestServiceGetAllUsers(t *testing.T) {
	// Arrange
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