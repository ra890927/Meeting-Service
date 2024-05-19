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
	GetAllCodeValuesByType(codeTypeID int) ([]models.CodeValue, error)

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

func (cr codeRepo) GetAllCodeValuesByType(codeTypeID int) ([]models.CodeValue, error) {
	db, err := gorm.Open(sqlite.Open(cr.dsn), &gorm.Config{})
	if err != nil {
		return []models.CodeValue{}, err
	}

	var codeValues []models.CodeValue
	err = db.Where("code_type_id = ?", codeTypeID).Find(&codeValues).Error
	if err != nil {
		return []models.CodeValue{}, err
	}

	return codeValues, nil
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
