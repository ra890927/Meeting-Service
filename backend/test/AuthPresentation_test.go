package test

import (
	"bytes"
	"meeting-center/src/models"
	"meeting-center/src/presentations"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/gin-gonic/gin"
)

type MockAuthService struct {
	mock.Mock
}

type AuthPresentationTestSuite struct {
	suite.Suite
	ap presentations.AuthPresentation
	as *MockAuthService
}

func (m *MockAuthService) Login(user *models.User) (*models.User, *string, error) {
	args := m.Called(user)
	return args.Get(0).(*models.User), args.Get(1).(*string), args.Error(2)
}

func (m *MockAuthService) Logout(token *string) error {
	args := m.Called(token)
	return args.Error(0)
}

func (m *MockAuthService) NewAuthService() *MockAuthService {
	return &MockAuthService{}
}

func (suite *AuthPresentationTestSuite) SetupTest() {
	suite.as = new(MockAuthService)
	suite.ap = presentations.NewAuthPresentation(suite.as)
}

func (suite *AuthPresentationTestSuite) TestNewAuthPresentation_1_input() {
	// Arrange
	as := new(MockAuthService)

	// Act
	ap := presentations.NewAuthPresentation(as)

	// Assert
	assert.NotNil(suite.T(), ap)
}

func (suite *AuthPresentationTestSuite) TestNewAuthPresentation_2_input() {
	// Arrange
	as := new(MockAuthService)

	// Act
	defer func() {
		if r := recover(); r == nil {
			assert.NotNil(suite.T(), r)
		}
	}()
	_ = presentations.NewAuthPresentation(as, as)
}

// func (suite *AuthPresentationTestSuite) TestLogout() {

func (suite *AuthPresentationTestSuite) TestLogin() {
	// Arrange
	user := &models.User{Email: "test", Password: "test"}
	token := ""
	suite.as.On("Login", user).Return(user, &token, nil)

	// Act
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/login", suite.ap.Login)
	reqBody := `{"email":"test","password":"test"}`
	req, _ := http.NewRequest("POST", "/login", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), 200, w.Code)
}

func (suite *AuthPresentationTestSuite) TestLogin_InvalidRequest() {
	// Arrange

	// Act
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/login", suite.ap.Login)
	reqBody := `{"email":"test","password":123}`
	req, _ := http.NewRequest("POST", "/login", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), 400, w.Code)
}

func (suite *AuthPresentationTestSuite) TestLogin_FailInService() {
	// Arrange
	user := &models.User{Email: "test", Password: "test"}
	token := ""
	suite.as.On("Login", user).Return(user, &token, assert.AnError)

	// Act
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/login", suite.ap.Login)
	reqBody := `{"email":"test","password":"test"}`
	req, _ := http.NewRequest("POST", "/login", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), 500, w.Code)
}

func (suite *AuthPresentationTestSuite) TestLogout_InvalidRequest() {
	// Arrange
	token := ""
	suite.as.On("Logout", &token).Return(nil)

	// Act
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/logout", suite.ap.Logout)
	req, _ := http.NewRequest("POST", "/logout", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), 400, w.Code)
}

func (suite *AuthPresentationTestSuite) TestLogout_TokenInCookie() {
	// Arrange
	token := "test"
	suite.as.On("Logout", &token).Return(nil)

	// Act
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/logout", suite.ap.Logout)
	req, _ := http.NewRequest("POST", "/logout", nil)
	req.AddCookie(&http.Cookie{Name: "token", Value: "test"})
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), 200, w.Code)
}

func (suite *AuthPresentationTestSuite) TestLogout_FailInService() {
	// Arrange
	token := "test"
	suite.as.On("Logout", &token).Return(assert.AnError)

	// Act
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/logout", suite.ap.Logout)
	req, _ := http.NewRequest("POST", "/logout", nil)
	req.AddCookie(&http.Cookie{Name: "token", Value: "test"})
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), 500, w.Code)
}

func (suite *AuthPresentationTestSuite) TestWhoAmI_UserNotFound() {
	// Arrange

	// Act
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/whoami", suite.ap.WhoAmI)
	req, _ := http.NewRequest("GET", "/whoami", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), 400, w.Code)
}

func addFakeUserToContext(c *gin.Context) {
	c.Set("validate_user", models.User{ID: 1, Email: "test"})
}

func (suite *AuthPresentationTestSuite) TestWhoAmI_Success() {
	// Arrange

	// Act
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/whoami", addFakeUserToContext, suite.ap.WhoAmI)
	// add c.Set("validate_user", "test") to the context

	req, _ := http.NewRequest("GET", "/whoami", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), 200, w.Code)
}

func TestAuthPresentationTestSuite(t *testing.T) {
	suite.Run(t, new(AuthPresentationTestSuite))
}
