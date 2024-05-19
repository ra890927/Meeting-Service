package test

import (
	"meeting-center/src/models"
	"meeting-center/src/repos"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type MockGormDB struct {
	mock.Mock
}

func TestRepoCreateUser(t *testing.T) {
	// Create a new user
	user := &models.User{
		Username: "test-username",
		Email:    "test@test.com",
		Password: "test-password",
		Role:     "test-role",
	}

	// todo: use test container to create a test database

}
