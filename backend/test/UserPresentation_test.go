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

func addFakeUserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := models.User{ID: 1}
		c.Set("validate_user", user)
		c.Next()
	}
}

type mockUserService struct {
	mock.Mock
}

func (m *mockUserService) CreateUser(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *mockUserService) UpdateUser(operator models.User, user *models.User) error {
	args := m.Called(operator, user)
	return args.Error(0)
}

func (m *mockUserService) GetAllUsers() ([]models.User, error) {
	args := m.Called()
	return args.Get(0).([]models.User), nil
}

func TestRegisterUser(t *testing.T) {
	user := models.User{
		Username: "test-user",
		Email:    "test@test.com",
		Password: "test-password",
	}
	// Mock the user service
	mockUserService := new(mockUserService)
	mockUserService.On("CreateUser", &user).Return(nil)
	up := presentations.NewUserPresentation(mockUserService)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/user", up.RegisterUser)

	// Create a request
	jsonDataString := `{"username":"test-user","email":"test@test.com","password":"test-password","role":"test-role"}`
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
		ID:       1,
		Username: "test-user-updated",
		Email:    "test@test.com",
		Password: "test-password-updated",
		Role:     "test-role-updated",
	}
	// - Mock the user service
	mockUserService := new(mockUserService)
	mockUserService.On("UpdateUser", models.User{ID: 1}, &user).Return(nil)
	up := presentations.NewUserPresentation(mockUserService)

	// - Set the mode to test
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.PUT("/user", addFakeUserMiddleware(), up.UpdateUser)

	// - Create a request
	jsonDataString := `{"id":1, "username":"test-user-updated","email":"test@test.com","password":"test-password-updated","role":"test-role-updated"}`
	jsonData := []byte(jsonDataString)
	req := httptest.NewRequest("PUT", "/user", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	// Act
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, 200, w.Code)
}
