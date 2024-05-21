package repos

import (
	db "meeting-center/src/io"
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

func CreateCodeType(codeType *models.CodeType) error {
	db := db.GetDBInstance()
	return db.Create(codeType).Error
}

func CreateCodeValue(codeValue *models.CodeValue) error {
	db := db.GetDBInstance()
	return db.Create(codeValue).Error
}

func GetAllCodeTypes() ([]models.CodeType, error) {
	db := db.GetDBInstance()

	var codeTypes []models.CodeType
	err := db.Preload("CodeValues").Find(&codeTypes).Error
	if err != nil {
		return nil, err
	}

	return codeTypes, nil
}

func GetAllCodeValuesByType(codeTypeID int) ([]models.CodeValue, error) {
	db := db.GetDBInstance()

	var codeValues []models.CodeValue
	err := db.Where("code_type_id = ?", codeTypeID).Find(&codeValues).Error
	if err != nil {
		return []models.CodeValue{}, err
	}

	return codeValues, nil
}

func UpdateCodeType(codeType *models.CodeType) error {
	db := db.GetDBInstance()
	return db.Save(codeType).Error
}

func UpdateCodeValue(codeValue *models.CodeValue) error {
	db := db.GetDBInstance()
	return db.Save(codeValue).Error
}

func DeleteCodeType(codeTypeID int) error {
	db := db.GetDBInstance()
	return db.Delete(&models.CodeType{}, codeTypeID).Error
}

func DeleteCodeValue(codeValueID int) error {
	db := db.GetDBInstance()
	return db.Delete(&models.CodeValue{}, codeValueID).Error
}
