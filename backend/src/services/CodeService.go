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

func NewCodeService(codeDomainArgs ...domains.CodeDomain) CodeService {
	if len(codeDomainArgs) == 1 {
		return CodeService(&codeService{cd: codeDomainArgs[0]})
	} else if len(codeDomainArgs) == 0 {
		return CodeService(&codeService{cd: domains.NewCodeDomain()})
	} else {
		panic("NewCodeService: too many arguments")
	}
}

func (cs codeService) CreateCodeType(codeType *models.CodeType) error {
	return cs.cd.CreateCodeType(codeType)
}

func (cs codeService) CreateCodeValue(codeValue *models.CodeValue) error {
	return cs.cd.CreateCodeValue(codeValue)
}

func (cs codeService) GetAllCodeTypes() ([]models.CodeType, error) {
	return cs.cd.GetAllCodeTypes()
}

func (cs codeService) GetAllCodeValuesByType(codeTypeID int) ([]models.CodeValue, error) {
	return cs.cd.GetAllCodeValuesByType(codeTypeID)
}

func (cs codeService) UpdateCodeType(codeType *models.CodeType) error {
	return cs.cd.UpdateCodeType(codeType)
}

func (cs codeService) UpdateCodeValue(codeValue *models.CodeValue) error {
	return cs.cd.UpdateCodeValue(codeValue)
}

func (cs codeService) DeleteCodeType(codeTypeID int) error {
	return cs.cd.DeleteCodeType(codeTypeID)
}

func (cs codeService) DeleteCodeValue(codeValueID int) error {
	return cs.cd.DeleteCodeValue(codeValueID)
}
