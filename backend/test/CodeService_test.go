package test

import (
	"meeting-center/src/models"
	"meeting-center/src/services"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockCodeDomain struct {
	mock.Mock
}

func (m *MockCodeDomain) CreateCodeType(codeType *models.CodeType) error {
	args := m.Called(codeType)
	return args.Error(0)
}

func (m *MockCodeDomain) CreateCodeValue(codeValue *models.CodeValue) error {
	args := m.Called(codeValue)
	return args.Error(0)
}

func (m *MockCodeDomain) GetAllCodeTypes() ([]models.CodeType, error) {
	args := m.Called()
	return args.Get(0).([]models.CodeType), args.Error(1)
}

func (m *MockCodeDomain) GetAllCodeValuesByType(codeTypeID int) ([]models.CodeValue, error) {
	args := m.Called(codeTypeID)
	return args.Get(0).([]models.CodeValue), args.Error(1)
}

func (m *MockCodeDomain) UpdateCodeType(codeType *models.CodeType) error {
	args := m.Called(codeType)
	return args.Error(0)
}

func (m *MockCodeDomain) UpdateCodeValue(codeValue *models.CodeValue) error {
	args := m.Called(codeValue)
	return args.Error(0)
}

func (m *MockCodeDomain) DeleteCodeType(codeTypeID int) error {
	args := m.Called(codeTypeID)
	return args.Error(0)
}

func (m *MockCodeDomain) DeleteCodeValue(codeValueID int) error {
	args := m.Called(codeValueID)
	return args.Error(0)
}

type CodeServiceTestSuite struct {
	suite.Suite
	cd services.CodeService
	cr *MockCodeDomain
}

func (suite *CodeServiceTestSuite) SetupTest() {
	suite.cr = new(MockCodeDomain)
	suite.cd = services.NewCodeService(suite.cr)
}

func TestCodeServiceTestSuite(t *testing.T) {
	suite.Run(t, new(CodeServiceTestSuite))
}

func (suite *CodeServiceTestSuite) TestCreateCodeType() {
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

func (suite *CodeServiceTestSuite) TestCreateCodeValue() {
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

func (suite *CodeServiceTestSuite) TestGetAllCodeTypes() {
	// Arrange
	suite.cr.On("GetAllCodeTypes").Return([]models.CodeType{}, nil)

	// Act
	_, err := suite.cd.GetAllCodeTypes()

	// Assert
	assert.NoError(suite.T(), err)
}

func (suite *CodeServiceTestSuite) TestGetAllCodeValuesByType() {
	// Arrange
	suite.cr.On("GetAllCodeValuesByType", 1).Return([]models.CodeValue{}, nil)

	// Act
	_, err := suite.cd.GetAllCodeValuesByType(1)

	// Assert
	assert.NoError(suite.T(), err)
}

func (suite *CodeServiceTestSuite) TestUpdateCodeType() {
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

func (suite *CodeServiceTestSuite) TestUpdateCodeValue() {
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

func (suite *CodeServiceTestSuite) TestDeleteCodeType() {
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

func (suite *CodeServiceTestSuite) TestDeleteCodeValue() {
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
