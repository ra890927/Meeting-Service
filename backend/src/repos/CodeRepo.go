package repos

import (
	"meeting-center/src/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type CodeRepo interface {
	// C
	CreateCodeType(codeType *models.CodeType) error
	CreateCodeValue(codeValue *models.CodeValue) error

	// R
	GetAllCodeTypes() ([]models.CodeType, error)
	GetCodeTypeByID(codeTypeID int) (*models.CodeType, error)
	GetCodeValueByID(codeValueID int) (*models.CodeValue, error)

	// U
	UpdateCodeType(codeType *models.CodeType) error
	UpdateCodeValue(codeValue *models.CodeValue) error

	// D
	DeleteCodeType(codeTypeID int) error
	DeleteCodeValue(codeValueID int) error
}

type codeRepo struct {
	dsn string
}

func NewCodeRepo(dsnArgs ...string) CodeRepo {
	dsn := "../sqlite.db"
	if len(dsnArgs) == 1 {
		dsn = dsnArgs[0]
	} else if len(dsnArgs) == 0 {
		dsn = "../sqlite.db"
	} else {
		panic("NewCodeRepo: too many arguments")
	}
	return CodeRepo(&codeRepo{dsn: dsn})
}

func (cr codeRepo) CreateCodeType(codeType *models.CodeType) error {
	db, err := gorm.Open(sqlite.Open(cr.dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	return db.Create(codeType).Error
}

func (cr codeRepo) CreateCodeValue(codeValue *models.CodeValue) error {
	db, err := gorm.Open(sqlite.Open(cr.dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	return db.Create(codeValue).Error
}

func (cr codeRepo) GetAllCodeTypes() ([]models.CodeType, error) {
	db, err := gorm.Open(sqlite.Open(cr.dsn), &gorm.Config{})
	if err != nil {
		return []models.CodeType{}, err
	}

	var codeTypes []models.CodeType
	err = db.Preload("CodeValues").Find(&codeTypes).Error
	if err != nil {
		return nil, err
	}

	return codeTypes, nil
}

func (cr codeRepo) GetCodeTypeByID(codeTypeID int) (*models.CodeType, error) {
	db, err := gorm.Open(sqlite.Open(cr.dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	var codeType models.CodeType
	err = db.Preload("CodeValues").Find(&codeType, codeTypeID).Error
	if err != nil {
		return nil, err
	}

	return &codeType, nil
}

func (cr codeRepo) GetCodeValueByID(codeValueID int) (*models.CodeValue, error) {
	db, err := gorm.Open(sqlite.Open(cr.dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	var codeValue models.CodeValue
	err = db.First(&codeValue, codeValueID).Error
	if err != nil {
		return nil, err
	}

	return &codeValue, nil
}

func (cr codeRepo) UpdateCodeType(codeType *models.CodeType) error {
	db, err := gorm.Open(sqlite.Open(cr.dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	return db.Save(codeType).Error
}

func (cr codeRepo) UpdateCodeValue(codeValue *models.CodeValue) error {
	db, err := gorm.Open(sqlite.Open(cr.dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	return db.Save(codeValue).Error
}

func (cr codeRepo) DeleteCodeType(codeTypeID int) error {
	db, err := gorm.Open(sqlite.Open(cr.dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	return db.Delete(&models.CodeType{}, codeTypeID).Error
}

func (cr codeRepo) DeleteCodeValue(codeValueID int) error {
	db, err := gorm.Open(sqlite.Open(cr.dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	return db.Delete(&models.CodeValue{}, codeValueID).Error
}
