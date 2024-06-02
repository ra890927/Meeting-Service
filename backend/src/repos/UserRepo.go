package repos

import (
	"meeting-center/src/clients"
	"meeting-center/src/models"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

type UserRepo interface {
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	GetAllUsers() ([]models.User, error)
	GetUserByEmail(email string) (models.User, error)
	GetUserByID(id uint) (models.User, error)
}

func NewUserRepo(dbArgs ...*gorm.DB) UserRepo {
	if len(dbArgs) == 0 {
		return userRepo{db: clients.GetDBInstance()}
	} else if len(dbArgs) == 1 {
		return userRepo{db: dbArgs[0]}
	} else {
		panic("Too many arguments")
	}
}

func (ur userRepo) CreateUser(user *models.User) error {
	result := ur.db.Create(user)
	// return the user if no errors
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (ur userRepo) UpdateUser(user *models.User) error {
	result := ur.db.Save(user)
	// return the user if no errors
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (ur userRepo) GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := ur.db.Find(&users)
	// return the users if no errors
	if result.Error != nil {
		return []models.User{}, result.Error
	}

	return users, nil
}

func (ur userRepo) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	result := ur.db.Where("email = ?", email).First(&user)
	// return the user if no errors
	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}

func (ur userRepo) GetUserByID(id uint) (models.User, error) {
	var user models.User
	result := ur.db.Where("id = ?", id).First(&user)
	// return the user if no errors
	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}
