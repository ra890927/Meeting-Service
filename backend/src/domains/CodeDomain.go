package domains

import (
	"meeting-center/src/models"
	"meeting-center/src/repos"
)

type CodeDomain interface {
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

type codeDomain struct {
	cr repos.CodeRepo
}

func NewCodeDomain(opt ...repos.CodeRepo) CodeDomain {
	cr := repos.NewCodeRepo()
	if len(opt) == 1 {
		cr = opt[0]
	}
	return &codeDomain{cr: cr}
}

func (cd *codeDomain) CreateCodeType(codeType *models.CodeType) error {
	return cd.cr.CreateCodeType(codeType)
}

func (cd *codeDomain) CreateCodeValue(codeValue *models.CodeValue) error {
	return cd.cr.CreateCodeValue(codeValue)
}

func (cd *codeDomain) GetAllCodeTypes() ([]models.CodeType, error) {
	return cd.cr.GetAllCodeTypes()
}

func (cd *codeDomain) GetAllCodeValuesByType(codeTypeID int) ([]models.CodeValue, error) {
	return cd.cr.GetAllCodeValuesByType(codeTypeID)
}

func (cd *codeDomain) UpdateCodeType(codeType *models.CodeType) error {
	return cd.cr.UpdateCodeType(codeType)
}

func (cd *codeDomain) UpdateCodeValue(codeValue *models.CodeValue) error {
	return cd.cr.UpdateCodeValue(codeValue)
}

func (cd *codeDomain) DeleteCodeType(codeTypeID int) error {
	return cd.cr.DeleteCodeType(codeTypeID)
}

func (cd *codeDomain) DeleteCodeValue(codeValueID int) error {
	return cd.cr.DeleteCodeValue(codeValueID)
}
