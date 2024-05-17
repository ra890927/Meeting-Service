package services

import (
	"meeting-center/src/domains"
	"meeting-center/src/models"
)

type CodeService interface {
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

type codeService struct {
	cd domains.CodeDomain
}

func NewCodeService(opt ...domains.CodeDomain) CodeService {
	cd := domains.NewCodeDomain()
	if len(opt) == 1 {
		cd = opt[0]
	}
	return &codeService{cd: cd}
}

func (cs *codeService) CreateCodeType(codeType *models.CodeType) error {
	return cs.cd.CreateCodeType(codeType)
}

func (cs *codeService) CreateCodeValue(codeValue *models.CodeValue) error {
	return cs.cd.CreateCodeValue(codeValue)
}

func (cs *codeService) GetAllCodeTypes() ([]models.CodeType, error) {
	return cs.cd.GetAllCodeTypes()
}

func (cs *codeService) GetAllCodeValuesByType(codeTypeID int) ([]models.CodeValue, error) {
	return cs.cd.GetAllCodeValuesByType(codeTypeID)
}

func (cs *codeService) UpdateCodeType(codeType *models.CodeType) error {
	return cs.cd.UpdateCodeType(codeType)
}

func (cs *codeService) UpdateCodeValue(codeValue *models.CodeValue) error {
	return cs.cd.UpdateCodeValue(codeValue)
}

func (cs *codeService) DeleteCodeType(codeTypeID int) error {
	return cs.cd.DeleteCodeType(codeTypeID)
}

func (cs *codeService) DeleteCodeValue(codeValueID int) error {
	return cs.cd.DeleteCodeValue(codeValueID)
}
