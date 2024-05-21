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
	GetAllCodeValuesByType(codeTypeID int) ([]models.CodeValue, error)

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

func (cr codeRepo) GetAllCodeValuesByType(codeTypeID int) ([]models.CodeValue, error) {
	var codeValues []models.CodeValue
	err := cr.db.Where("code_type_id = ?", codeTypeID).Find(&codeValues).Error
	if err != nil {
		return []models.CodeValue{}, err
	}

	return codeValues, nil
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
