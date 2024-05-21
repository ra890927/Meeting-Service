package services

import (
	"errors"
	"meeting-center/src/domains"
	"meeting-center/src/models"
)

type CodeService interface {
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

func (cs codeService) GetCodeTypeByID(codeTypeID int) (*models.CodeType, error) {
	return cs.cd.GetCodeTypeByID(codeTypeID)
}

func (cs codeService) GetCodeValueByID(codeValueID int) (*models.CodeValue, error) {
	return cs.cd.GetCodeValueByID(codeValueID)
}

func (cs codeService) UpdateCodeType(codeType *models.CodeType) error {
	// check if such codeType exists
	_, err := cs.cd.GetCodeTypeByID(codeType.ID)
	if err != nil {
		return errors.New("codeType not found")
	}

	return cs.cd.UpdateCodeType(codeType)
}

func (cs codeService) UpdateCodeValue(codeValue *models.CodeValue) error {
	// check if such codeValue exists
	_, err := cs.cd.GetCodeValueByID(codeValue.ID)
	if err != nil {
		return errors.New("codeValue not found")
	}
	return cs.cd.UpdateCodeValue(codeValue)
}

func (cs codeService) DeleteCodeType(codeTypeID int) error {
	// check if such codeType exists
	_, err := cs.cd.GetCodeTypeByID(codeTypeID)
	if err != nil {
		return errors.New("codeType not found")
	}
	return cs.cd.DeleteCodeType(codeTypeID)
}

func (cs codeService) DeleteCodeValue(codeValueID int) error {
	// check if such codeValue exists
	_, err := cs.cd.GetCodeValueByID(codeValueID)
	if err != nil {
		return errors.New("codeValue not found")
	}
	return cs.cd.DeleteCodeValue(codeValueID)
}
