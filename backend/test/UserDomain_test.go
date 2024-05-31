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

func (m *MockUserRepo) CreateUser(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepo) UpdateUser(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepo) GetUserByEmail(email string) (models.User, error) {
	args := m.Called(email)
	return args.Get(0).(models.User), nil
}

func (m *MockUserRepo) GetAllUsers() ([]models.User, error) {
	args := m.Called()
	return args.Get(0).([]models.User), nil
}

func (m *MockUserRepo) GetUserByID(id uint) (models.User, error) {
	args := m.Called(id)
	return args.Get(0).(models.User), nil
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
	mockUserRepo.On("CreateUser", user).Return(nil)
	ud := domains.NewUserDomain(mockUserRepo)

	// Act
	err := ud.CreateUser(user)

	// Assert
	assert.NoError(t, err)
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
	mockUserRepo.On("UpdateUser", user).Return(nil)
	ud := domains.NewUserDomain(mockUserRepo)

	// Act
	err := ud.UpdateUser(user)

	// Assert
	assert.NoError(t, err)
}

func TestDomainGetAllUsers(t *testing.T) {
	// Arrange
	// mock the user repo
	mockUserRepo := new(MockUserRepo)
	mockUserRepo.On("GetAllUsers").Return([]models.User{}, nil)
	ud := domains.NewUserDomain(mockUserRepo)

	// Act
	users, err := ud.GetAllUsers()

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, users)
}

func TestDomainGetUserByEmail(t *testing.T) {
	// Arrange
	targetEmail := "a@a.com"
	// mock the user repo
	mockUserRepo := new(MockUserRepo)
	mockUserRepo.On("GetUserByEmail", targetEmail).Return(models.User{}, nil)
	ud := domains.NewUserDomain(mockUserRepo)

	// Act
	user, err := ud.GetUserByEmail(targetEmail)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestDomainGetUserByID(t *testing.T) {
	// Arrange
	targetID := uint(1)
	// mock the user repo
	mockUserRepo := new(MockUserRepo)
	mockUserRepo.On("GetUserByID", targetID).Return(models.User{}, nil)
	ud := domains.NewUserDomain(mockUserRepo)

	// Act
	user, err := ud.GetUserByID(targetID)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, user)
}
