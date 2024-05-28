package test

import (
	"meeting-center/src/models"
	. "meeting-center/src/repos"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type UserRepoTestSuite struct {
	suite.Suite
	ur UserRepo
	db *gorm.DB
}

func (suite *UserRepoTestSuite) SetupTest() {
	db, err := gorm.Open(sqlite.Open("./test.sqlite"), &gorm.Config{})
	assert.NoError(suite.T(), err)

	err = db.AutoMigrate(&models.User{})
	assert.NoError(suite.T(), err)

	suite.db = db
	suite.ur = NewUserRepo(db)
}

func (suite *UserRepoTestSuite) TearDownTest() {
	db, err := suite.db.DB()
	assert.NoError(suite.T(), err)
	err = db.Close()
	assert.NoError(suite.T(), err)
}

func (suite *UserRepoTestSuite) TestCreateUser() {
	user := &models.User{
		Username: "TestCreateUser",
		Email:    "TestCreateUser@test.com",
		Password: "test-password",
		Role:     "test-role",
	}
	_, err := suite.ur.CreateUser(user)
	assert.NoError(suite.T(), err)
	assert.NotZero(suite.T(), user.ID)
}

func (suite *UserRepoTestSuite) TestGetUserByEmail() {
	user := &models.User{
		Username: "TestGetUserByEmail",
		Email:    "TestGetUserByEmail@test.com",
	}
	_, err := suite.ur.CreateUser(user)
	assert.NoError(suite.T(), err)

	foundUser, err := suite.ur.GetUserByEmail(user.Email)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), user.Username, foundUser.Username)
	assert.Equal(suite.T(), user.Email, foundUser.Email)
}

func (suite *UserRepoTestSuite) TestGetUserByID() {
	user := &models.User{
		Username: "TestGetUserByID",
		Email:    "TestGetUserByID@test.com",
	}
	_, err := suite.ur.CreateUser(user)
	assert.NoError(suite.T(), err)

	foundUser, err := suite.ur.GetUserByID(user.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), user.Username, foundUser.Username)
	assert.Equal(suite.T(), user.Email, foundUser.Email)
}

func (suite *UserRepoTestSuite) TestGetAllUsers() {
	users := []models.User{
		{
			Username: "TestGetAllUsers1",
			Email:    "TestGetAllUsers1@test.com",
		},
		{
			Username: "TestGetAllUsers2",
			Email:    "TestGetAllUsers2@test.com",
		},
	}

	orifinalFoundUsers, err := suite.ur.GetAllUsers()

	for _, user := range users {
		_, err := suite.ur.CreateUser(&user)
		assert.NoError(suite.T(), err)
	}

	foundUsers, err := suite.ur.GetAllUsers()
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), foundUsers, len(orifinalFoundUsers)+len(users))
}

func TestUserRepoTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepoTestSuite))
}
