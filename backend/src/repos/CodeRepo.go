package repos

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"meeting-center/src/models"
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

func NewCodeRepo(opt ...string) CodeRepo {
	dsn := "../sqlite.db"
	if len(opt) == 1 {
		dsn = opt[0]
	}
	return &codeRepo{dsn: dsn}
}

func (cr *codeRepo) CreateCodeType(codeType *models.CodeType) error {
	db, err := gorm.Open(sqlite.Open(cr.dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	return db.Create(codeType).Error
}

func (cr *codeRepo) CreateCodeValue(codeValue *models.CodeValue) error {
	db, err := gorm.Open(sqlite.Open(cr.dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	return db.Create(codeValue).Error
}

func (cr *codeRepo) GetAllCodeTypes() ([]models.CodeType, error) {
	db, err := gorm.Open(sqlite.Open(cr.dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	var codeTypes []models.CodeType
	err = db.Preload("CodeValues").Find(&codeTypes).Error
	if err != nil {
		return nil, err
	}

	return codeTypes, nil
}

func (cr *codeRepo) GetAllCodeValuesByType(codeTypeID int) ([]models.CodeValue, error) {
	db, err := gorm.Open(sqlite.Open(cr.dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	var codeValues []models.CodeValue
	err = db.Where("code_type_id = ?", codeTypeID).Find(&codeValues).Error
	if err != nil {
		return nil, err
	}

	return codeValues, nil
}

func (cr *codeRepo) UpdateCodeType(codeType *models.CodeType) error {
	db, err := gorm.Open(sqlite.Open(cr.dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	return db.Save(codeType).Error
}

func (cr *codeRepo) UpdateCodeValue(codeValue *models.CodeValue) error {
	db, err := gorm.Open(sqlite.Open(cr.dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	return db.Save(codeValue).Error
}

func (cr *codeRepo) DeleteCodeType(codeTypeID int) error {
	db, err := gorm.Open(sqlite.Open(cr.dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	return db.Delete(&models.CodeType{}, codeTypeID).Error
}

func (cr *codeRepo) DeleteCodeValue(codeValueID int) error {
	db, err := gorm.Open(sqlite.Open(cr.dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	return db.Delete(&models.CodeValue{}, codeValueID).Error
}
