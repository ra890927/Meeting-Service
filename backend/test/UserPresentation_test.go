package test

import (
	"bytes"

	"errors"
	"meeting-center/src/models"
	"meeting-center/src/presentations"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

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
	return args.Get(0).([]models.User), args.Error(1)
}

func addFakeUserMiddlewareUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := models.User{ID: 1}
		c.Set("validate_user", user)
		c.Next()
	}
}

func TestNewUserPresentation(t *testing.T) {
	// Test with one argument
	mockService := new(mockUserService)
	up := presentations.NewUserPresentation(mockService)
	assert.NotNil(t, up)

	// Test with no arguments
	up = presentations.NewUserPresentation()
	assert.NotNil(t, up)

	// Test with multiple arguments should panic
	assert.Panics(t, func() {
		presentations.NewUserPresentation(mockService, mockService)
	})
}

func TestRegisterUser(t *testing.T) {
	user := models.User{
		Username: "test-user",
		Email:    "test@test.com",
		Password: "test-password",
	}

	mockUserService := new(mockUserService)
	mockUserService.On("CreateUser", &user).Return(nil).Once()
	up := presentations.NewUserPresentation(mockUserService)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/user", up.RegisterUser)

	// Test successful registration
	jsonData := `{"username":"test-user","email":"test@test.com","password":"test-password"}`
	req := httptest.NewRequest("POST", "/user", bytes.NewBufferString(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	mockUserService.AssertExpectations(t)

	// Test invalid JSON
	invalidJSON := `{"username":12345}`
	req = httptest.NewRequest("POST", "/user", bytes.NewBufferString(invalidJSON))
	req.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)

	// Test service error
	mockUserService.On("CreateUser", &user).Return(errors.New("some error")).Once()
	req = httptest.NewRequest("POST", "/user", bytes.NewBufferString(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Code)
}

func TestUpdateUser(t *testing.T) {
	user := models.User{
		ID:       1,
		Username: "test-user-updated",
		Email:    "test@test.com",
		Password: "test-password-updated",
		Role:     "test-role-updated",
	}

	mockUserService := new(mockUserService)
	mockUserService.On("UpdateUser", models.User{ID: 1}, &user).Return(nil).Once()
	up := presentations.NewUserPresentation(mockUserService)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.PUT("/user", addFakeUserMiddlewareUser(), up.UpdateUser)

	// Test successful update
	jsonData := `{"id":1, "username":"test-user-updated","email":"test@test.com","password":"test-password-updated","role":"test-role-updated"}`
	req := httptest.NewRequest("PUT", "/user", bytes.NewBufferString(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	mockUserService.AssertExpectations(t)

	// Test invalid JSON
	invalidJSON := `{"id":"invalid"}`
	req = httptest.NewRequest("PUT", "/user", bytes.NewBufferString(invalidJSON))
	req.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)

	// Test service error
	mockUserService.On("UpdateUser", models.User{ID: 1}, &user).Return(errors.New("some error")).Once()
	req = httptest.NewRequest("PUT", "/user", bytes.NewBufferString(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Code)
}

func TestGetAllUsers(t *testing.T) {
	users := []models.User{
		{ID: 1, Username: "user1", Email: "user1@test.com", Role: "role1"},
		{ID: 2, Username: "user2", Email: "user2@test.com", Role: "role2"},
	}

	mockUserService := new(mockUserService)
	mockUserService.On("GetAllUsers").Return(users, nil).Once()
	up := presentations.NewUserPresentation(mockUserService)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/users", up.GetAllUsers)

	req := httptest.NewRequest("GET", "/users", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	mockUserService.AssertExpectations(t)

	// Test service error
	mockUserService.On("GetAllUsers").Return(nil, errors.New("some error")).Once()
	req = httptest.NewRequest("GET", "/users", nil)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Code)
}
