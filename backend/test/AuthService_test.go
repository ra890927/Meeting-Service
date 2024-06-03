package test

import (
	"meeting-center/src/models"
	"meeting-center/src/services"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockAuthDomain struct {
	mock.Mock
}

type AuthServiceTestSuite struct {
	suite.Suite
	ad services.AuthService
	as *MockAuthDomain
}

func (m *MockAuthDomain) Login(user *models.User) (*models.User, *string, error) {
	args := m.Called(user)
	return args.Get(0).(*models.User), args.Get(1).(*string), args.Error(2)
}

func (m *MockAuthDomain) Logout(token *string) error {
	args := m.Called(token)
	return args.Error(0)
}

func (m *MockAuthDomain) NewAuthDomain() *MockAuthDomain {
	return &MockAuthDomain{}
}

func (suite *AuthServiceTestSuite) SetupTest() {
	suite.as = new(MockAuthDomain)
	suite.ad = services.NewAuthService(suite.as)
}

func (suite *AuthServiceTestSuite) TestNewAuthService_1_input() {
	// Arrange
	ad := new(MockAuthDomain)

	// Act
	result := services.NewAuthService(ad)

	// Assert
	assert.NotNil(suite.T(), result)
}

func (suite *AuthServiceTestSuite) TestNewAuthService_2_input() {
	// Arrange
	ad := new(MockAuthDomain)

	// Act
	defer func() {
		if r := recover(); r == nil {
			assert.NotNil(suite.T(), r)
		}
	}()
	_ = services.NewAuthService(ad, ad)
}

func (suite *AuthServiceTestSuite) TestLogin_Success() {
	// Arrange
	user := &models.User{}
	suite.as.On("Login", user).Return(user, new(string), nil)

	// Act
	_, _, err := suite.ad.Login(user)

	// Assert
	assert.Nil(suite.T(), err)
}

func (suite *AuthServiceTestSuite) TestLogin_FailInDomain() {
	// Arrange
	user := &models.User{}
	stringPointer := new(string)
	suite.as.On("Login", user).Return(user, stringPointer, assert.AnError)

	// Act
	_, _, err := suite.ad.Login(user)

	// Assert
	assert.NotNil(suite.T(), err)
}

func (suite *AuthServiceTestSuite) TestLogout_Success() {
	// Arrange
	token := new(string)
	suite.as.On("Logout", token).Return(nil)

	// Act
	err := suite.ad.Logout(token)

	// Assert
	assert.Nil(suite.T(), err)
}

func (suite *AuthServiceTestSuite) TestLogout_FailInDomain() {
	// Arrange
	token := new(string)
	suite.as.On("Logout", token).Return(assert.AnError)

	// Act
	err := suite.ad.Logout(token)

	// Assert
	assert.NotNil(suite.T(), err)
}

func TestAuthServiceTestSuite(t *testing.T) {
	suite.Run(t, new(AuthServiceTestSuite))
}
