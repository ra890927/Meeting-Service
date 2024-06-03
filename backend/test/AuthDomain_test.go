package test

import (
	"meeting-center/src/domains"
	"meeting-center/src/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockAuthRepo struct {
	mock.Mock
}

type AuthDomainTestSuite struct {
	suite.Suite
	ad domains.AuthDomain
	ar *MockAuthRepo
}

func (m *MockAuthRepo) Login(user *models.User) (*models.User, *string, error) {
	args := m.Called(user)
	return args.Get(0).(*models.User), args.Get(1).(*string), args.Error(2)
}

func (m *MockAuthRepo) Logout(token *string) error {
	args := m.Called(token)
	return args.Error(0)
}

func (m *MockAuthRepo) NewAuthRepository() *MockAuthRepo {
	return &MockAuthRepo{}
}

func (suite *AuthDomainTestSuite) SetupTest() {
	suite.ar = new(MockAuthRepo)
	suite.ad = domains.NewAuthDomain(suite.ar)
}

func (suite *AuthDomainTestSuite) TestNewAuthDomain_1_input() {
	// Arrange
	ar := new(MockAuthRepo)

	// Act
	result := domains.NewAuthDomain(ar)

	// Assert
	assert.NotNil(suite.T(), result)
}

func (suite *AuthDomainTestSuite) TestNewAuthDomain_2_input() {
	// Arrange
	ar := new(MockAuthRepo)

	// Act
	defer func() {
		if r := recover(); r == nil {
			assert.NotNil(suite.T(), r)
		}
	}()
	_ = domains.NewAuthDomain(ar, ar)
}

func (suite *AuthDomainTestSuite) Test_Login_Success() {
	// Arrange
	user := &models.User{}
	suite.ar.On("Login", user).Return(user, new(string), nil)

	// Act
	result, _, err := suite.ad.Login(user)

	// Assert
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), result)
}

func (suite *AuthDomainTestSuite) Test_Login_FailInRepo() {
	// Arrange
	user := &models.User{}
	suite.ar.On("Login", user).Return(user, new(string), assert.AnError)

	// Act
	_, _, err := suite.ad.Login(user)

	// Assert
	assert.NotNil(suite.T(), err)
}

func (suite *AuthDomainTestSuite) Test_Logout_Success() {
	// Arrange
	token := new(string)
	suite.ar.On("Logout", token).Return(nil)

	// Act
	err := suite.ad.Logout(token)

	// Assert
	assert.Nil(suite.T(), err)
}

func (suite *AuthDomainTestSuite) Test_Logout_FailInRepo() {
	// Arrange
	token := new(string)
	suite.ar.On("Logout", token).Return(assert.AnError)

	// Act
	err := suite.ad.Logout(token)

	// Assert
	assert.NotNil(suite.T(), err)
}

func TestAuthDomainTestSuite(t *testing.T) {
	suite.Run(t, new(AuthDomainTestSuite))
}
