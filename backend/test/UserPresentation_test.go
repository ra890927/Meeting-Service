package test

import (
	"bytes"
	"meeting-center/src/models"
	"meeting-center/src/presentations"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/gin-gonic/gin"
)

type mockUserService struct {
	mock.Mock
}

func (m *mockUserService) CreateUser(user *models.User) (*models.User, error) {
	args := m.Called(user)
	return args.Get(0).(*models.User), nil
}

func (m *mockUserService) UpdateUser(user *models.User) (*models.User, error) {
	args := m.Called(user)
	return args.Get(0).(*models.User), nil
}

func TestRegisterUser(t *testing.T) {
	user := models.User{
		Username: "test-user",
		Email:    "test-email",
		Password: "test-password",
		Role:     "test-role",
	}
	// Mock the user service
	mockUserService := new(mockUserService)
	mockUserService.On("CreateUser", &user).Return(&user, nil)
	up := presentations.NewUserPresentation(mockUserService)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/user", up.RegisterUser)

	// Create a request
	jsonDataString := `{"username":"test-user","email":"test-email","password":"test-password","role":"test-role"}`
	jsonData := []byte(jsonDataString)
	req := httptest.NewRequest("POST", "/user", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

}

func TestUpdateUser(t *testing.T) {
	// Arrange
	user := models.User{
		Username: "test-user-updated",
		Email:    "test-email-updated",
		Password: "test-password-updated",
		Role:     "test-role-updated",
	}
	// - Mock the user service
	mockUserService := new(mockUserService)
	mockUserService.On("UpdateUser", &user).Return(&user, nil)
	up := presentations.NewUserPresentation(mockUserService)

	// - Set the mode to test
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.PUT("/user", up.UpdateUser)

	// - Create a request
	jsonDataString := `{"username":"test-user-updated","email":"test-email-updated","password":"test-password-updated","role":"test-role-updated"}`
	jsonData := []byte(jsonDataString)
	req := httptest.NewRequest("PUT", "/user", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	// Act
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, 200, w.Code)
}
