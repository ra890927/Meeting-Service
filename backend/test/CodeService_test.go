package test

import (
	"errors"
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

func (m *MockCodeDomain) GetCodeTypeByID(codeTypeID int) (*models.CodeType, error) {
	args := m.Called(codeTypeID)
	return args.Get(0).(*models.CodeType), args.Error(1)
}

func (m *MockCodeDomain) GetCodeValueByID(codeValueID int) (*models.CodeValue, error) {
	args := m.Called(codeValueID)
	return args.Get(0).(*models.CodeValue), args.Error(1)
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

func (suite *CodeServiceTestSuite) TestNewCodeService() {
	// Arrange
	mockDomain := new(MockCodeDomain)

	// Act and Assert

	// Test case with one argument
	cs := services.NewCodeService(mockDomain)
	assert.NotNil(suite.T(), cs)

	// Test case with no arguments
	// cs = services.NewCodeService()
	// assert.NotNil(suite.T(), cs)

	// Test case with multiple arguments should panic
	assert.Panics(suite.T(), func() {
		services.NewCodeService(mockDomain, mockDomain)
	})
}

func (suite *CodeServiceTestSuite) TestCreateCodeType() {
	codeType := &models.CodeType{
		TypeName: "TestType",
		TypeDesc: "This is a test type",
	}
	suite.cr.On("CreateCodeType", codeType).Return(nil)

	err := suite.cd.CreateCodeType(codeType)

	assert.NoError(suite.T(), err)
}

func (suite *CodeServiceTestSuite) TestCreateCodeValue() {
	codeValue := &models.CodeValue{
		CodeTypeID:    1,
		CodeValue:     "TestValue",
		CodeValueDesc: "This is a test value",
	}
	suite.cr.On("CreateCodeValue", codeValue).Return(nil)

	err := suite.cd.CreateCodeValue(codeValue)

	assert.NoError(suite.T(), err)
}

func (suite *CodeServiceTestSuite) TestGetAllCodeTypes() {
	suite.cr.On("GetAllCodeTypes").Return([]models.CodeType{}, nil)

	_, err := suite.cd.GetAllCodeTypes()

	assert.NoError(suite.T(), err)
}

func (suite *CodeServiceTestSuite) TestGetCodeTypeByID() {
	suite.cr.On("GetCodeTypeByID", 1).Return(&models.CodeType{}, nil)

	_, err := suite.cd.GetCodeTypeByID(1)

	assert.NoError(suite.T(), err)
}

func (suite *CodeServiceTestSuite) TestGetCodeValueByID() {
	suite.cr.On("GetCodeValueByID", 1).Return(&models.CodeValue{}, nil)

	_, err := suite.cd.GetCodeValueByID(1)

	assert.NoError(suite.T(), err)
}

func (suite *CodeServiceTestSuite) TestUpdateCodeType() {
	codeType := &models.CodeType{
		ID:       1,
		TypeName: "TestType",
		TypeDesc: "This is a test type",
	}
	suite.cr.On("GetCodeTypeByID", codeType.ID).Return(codeType, nil).Once()
	suite.cr.On("UpdateCodeType", codeType).Return(nil).Once()

	err := suite.cd.UpdateCodeType(codeType)

	assert.NoError(suite.T(), err)

	suite.cr.On("GetCodeTypeByID", codeType.ID).Return((*models.CodeType)(nil), errors.New("not found")).Once()

	err = suite.cd.UpdateCodeType(codeType)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "codeType not found", err.Error())
}

func (suite *CodeServiceTestSuite) TestUpdateCodeValue() {
	codeValue := &models.CodeValue{
		ID:            1,
		CodeTypeID:    1,
		CodeValue:     "TestValue",
		CodeValueDesc: "This is a test value",
	}
	suite.cr.On("GetCodeValueByID", codeValue.ID).Return(codeValue, nil).Once()
	suite.cr.On("UpdateCodeValue", codeValue).Return(nil).Once()

	err := suite.cd.UpdateCodeValue(codeValue)

	assert.NoError(suite.T(), err)

	suite.cr.On("GetCodeValueByID", codeValue.ID).Return((*models.CodeValue)(nil), errors.New("not found")).Once()

	err = suite.cd.UpdateCodeValue(codeValue)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "codeValue not found", err.Error())
}

func (suite *CodeServiceTestSuite) TestDeleteCodeType() {
	codeTypeID := 1
	codeType := &models.CodeType{
		ID: codeTypeID,
	}
	suite.cr.On("GetCodeTypeByID", codeTypeID).Return(codeType, nil).Once()
	suite.cr.On("DeleteCodeType", codeTypeID).Return(nil).Once()

	err := suite.cd.DeleteCodeType(codeTypeID)

	assert.NoError(suite.T(), err)

	suite.cr.On("GetCodeTypeByID", codeTypeID).Return((*models.CodeType)(nil), errors.New("not found")).Once()

	err = suite.cd.DeleteCodeType(codeTypeID)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "codeType not found", err.Error())
}

func (suite *CodeServiceTestSuite) TestDeleteCodeValue() {
	codeValueID := 1
	codeValue := &models.CodeValue{
		ID: codeValueID,
	}
	suite.cr.On("GetCodeValueByID", codeValueID).Return(codeValue, nil).Once()
	suite.cr.On("DeleteCodeValue", codeValueID).Return(nil).Once()

	err := suite.cd.DeleteCodeValue(codeValueID)

	assert.NoError(suite.T(), err)

	suite.cr.On("GetCodeValueByID", codeValueID).Return((*models.CodeValue)(nil), errors.New("not found")).Once()

	err = suite.cd.DeleteCodeValue(codeValueID)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "codeValue not found", err.Error())
}
