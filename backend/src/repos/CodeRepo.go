package repos

import (
	db "meeting-center/src/io"
	"meeting-center/src/models"

	"gorm.io/gorm"
)

type codeRepo struct {
	db *gorm.DB
}

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

func NewCodeRepo(dbArgs ...*gorm.DB) CodeRepo {
	if len(dbArgs) == 0 {
		return codeRepo{db: db.GetDBInstance()}
	} else if len(dbArgs) == 1 {
		return codeRepo{db: dbArgs[0]}
	} else {
		panic("Too many arguments")
	}
}

func (cr codeRepo) CreateCodeType(codeType *models.CodeType) error {
	return cr.db.Create(codeType).Error
}

func (cr codeRepo) CreateCodeValue(codeValue *models.CodeValue) error {
	return cr.db.Create(codeValue).Error
}

func (cr codeRepo) GetAllCodeTypes() ([]models.CodeType, error) {
	var codeTypes []models.CodeType
	err := cr.db.Preload("CodeValues").Find(&codeTypes).Error
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
	return cr.db.Save(codeType).Error
}

func (cr codeRepo) UpdateCodeValue(codeValue *models.CodeValue) error {
	return cr.db.Save(codeValue).Error
}

func (cr codeRepo) DeleteCodeType(codeTypeID int) error {
	return cr.db.Delete(&models.CodeType{}, codeTypeID).Error
}

func (cr codeRepo) DeleteCodeValue(codeValueID int) error {
	return cr.db.Delete(&models.CodeValue{}, codeValueID).Error
}
