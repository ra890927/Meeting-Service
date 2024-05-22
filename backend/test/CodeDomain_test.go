package test

import (
	"meeting-center/src/domains"
	"meeting-center/src/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockCodeRepo struct {
	mock.Mock
}

func (m *MockCodeRepo) CreateCodeType(codeType *models.CodeType) error {
	args := m.Called(codeType)
	return args.Error(0)
}

func (m *MockCodeRepo) CreateCodeValue(codeValue *models.CodeValue) error {
	args := m.Called(codeValue)
	return args.Error(0)
}

func (m *MockCodeRepo) GetAllCodeTypes() ([]models.CodeType, error) {
	args := m.Called()
	return args.Get(0).([]models.CodeType), args.Error(1)
}

func (m *MockCodeRepo) GetAllCodeValuesByType(codeTypeID int) ([]models.CodeValue, error) {
	args := m.Called(codeTypeID)
	return args.Get(0).([]models.CodeValue), args.Error(1)
}

func (m *MockCodeRepo) UpdateCodeType(codeType *models.CodeType) error {
	args := m.Called(codeType)
	return args.Error(0)
}

func (m *MockCodeRepo) UpdateCodeValue(codeValue *models.CodeValue) error {
	args := m.Called(codeValue)
	return args.Error(0)
}

func (m *MockCodeRepo) DeleteCodeType(codeTypeID int) error {
	args := m.Called(codeTypeID)
	return args.Error(0)
}

func (m *MockCodeRepo) DeleteCodeValue(codeValueID int) error {
	args := m.Called(codeValueID)
	return args.Error(0)
}

func (m *MockCodeRepo) GetCodeTypeByID(codeTypeID int) (*models.CodeType, error) {
	args := m.Called(codeTypeID)
	return args.Get(0).(*models.CodeType), args.Error(1)
}

func (m *MockCodeRepo) GetCodeValueByID(codeValueID int) (*models.CodeValue, error) {
	args := m.Called(codeValueID)
	return args.Get(0).(*models.CodeValue), args.Error(1)
}

type CodeDomainTestSuite struct {
	suite.Suite
	cd domains.CodeDomain
	cr *MockCodeRepo
}

func (suite *CodeDomainTestSuite) SetupTest() {
	suite.cr = new(MockCodeRepo)
	suite.cd = domains.NewCodeDomain(suite.cr)
}

func TestCodeDomainTestSuite(t *testing.T) {
	suite.Run(t, new(CodeDomainTestSuite))
}

func (suite *CodeDomainTestSuite) TestCreateCodeType() {
	// Arrange
	codeType := &models.CodeType{
		TypeName: "TestType",
		TypeDesc: "This is a test type",
	}
	suite.cr.On("CreateCodeType", codeType).Return(nil)

	// Act
	err := suite.cd.CreateCodeType(codeType)

	// Assert
	assert.NoError(suite.T(), err)
}

func (suite *CodeDomainTestSuite) TestCreateCodeValue() {
	// Arrange
	codeValue := &models.CodeValue{
		CodeTypeID:    1,
		CodeValue:     "TestValue",
		CodeValueDesc: "This is a test value",
	}
	suite.cr.On("CreateCodeValue", codeValue).Return(nil)

	// Act
	err := suite.cd.CreateCodeValue(codeValue)

	// Assert
	assert.NoError(suite.T(), err)
}

func (suite *CodeDomainTestSuite) TestGetAllCodeTypes() {
	// Arrange
	suite.cr.On("GetAllCodeTypes").Return([]models.CodeType{}, nil)

	// Act
	_, err := suite.cd.GetAllCodeTypes()

	// Assert
	assert.NoError(suite.T(), err)
}

func (suite *CodeDomainTestSuite) TestGetCodeTypeByID() {
	// Arrange
	suite.cr.On("GetCodeTypeByID", 1).Return(&models.CodeType{}, nil)

	// Act
	_, err := suite.cd.GetCodeTypeByID(1)

	// Assert
	assert.NoError(suite.T(), err)
}

func (suite *CodeDomainTestSuite) TestGetCodeValueByID() {
	// Arrange
	suite.cr.On("GetCodeValueByID", 1).Return(&models.CodeValue{}, nil)

	// Act
	_, err := suite.cd.GetCodeValueByID(1)

	// Assert
	assert.NoError(suite.T(), err)
}

func (suite *CodeDomainTestSuite) TestUpdateCodeType() {
	// Arrange
	codeType := &models.CodeType{
		ID:       1,
		TypeName: "TestType",
		TypeDesc: "This is a test type",
	}
	suite.cr.On("UpdateCodeType", codeType).Return(nil)

	// Act
	err := suite.cd.UpdateCodeType(codeType)

	// Assert
	assert.NoError(suite.T(), err)
}

func (suite *CodeDomainTestSuite) TestUpdateCodeValue() {
	// Arrange
	codeValue := &models.CodeValue{
		ID:            1,
		CodeTypeID:    1,
		CodeValue:     "TestValue",
		CodeValueDesc: "This is a test value",
	}
	suite.cr.On("UpdateCodeValue", codeValue).Return(nil)

	// Act
	err := suite.cd.UpdateCodeValue(codeValue)

	// Assert
	assert.NoError(suite.T(), err)
}

func (suite *CodeDomainTestSuite) TestDeleteCodeType() {
	// Arrange
	codeType := &models.CodeType{
		ID: 1,
	}
	suite.cr.On("DeleteCodeType", codeType.ID).Return(nil)

	// Act
	err := suite.cd.DeleteCodeType(1)

	// Assert
	assert.NoError(suite.T(), err)
}

func (suite *CodeDomainTestSuite) TestDeleteCodeValue() {
	// Arrange
	codeValue := &models.CodeValue{
		ID: 1,
	}
	suite.cr.On("DeleteCodeValue", codeValue.ID).Return(nil)

	// Act
	err := suite.cd.DeleteCodeValue(codeValue.ID)

	// Assert
	assert.NoError(suite.T(), err)
}
