package test

import (
	"bytes"
	"meeting-center/src/models"
	"meeting-center/src/presentations"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/gin-gonic/gin"
)

type MockCodeService struct {
	mock.Mock
}

func (m *MockCodeService) CreateCodeType(codeType *models.CodeType) error {
	args := m.Called(codeType)
	return args.Error(0)
}

func (m *MockCodeService) CreateCodeValue(codeValue *models.CodeValue) error {
	args := m.Called(codeValue)
	return args.Error(0)
}

func (m *MockCodeService) GetAllCodeTypes() ([]models.CodeType, error) {
	args := m.Called()
	return args.Get(0).([]models.CodeType), args.Error(1)
}

func (m *MockCodeService) GetCodeTypeByID(codeTypeID int) (*models.CodeType, error) {
	args := m.Called(codeTypeID)
	return args.Get(0).(*models.CodeType), args.Error(1)
}

func (m *MockCodeService) GetCodeValueByID(codeValueID int) (*models.CodeValue, error) {
	args := m.Called(codeValueID)
	return args.Get(0).(*models.CodeValue), args.Error(1)
}

func (m *MockCodeService) UpdateCodeType(codeType *models.CodeType) error {
	args := m.Called(codeType)
	return args.Error(0)
}

func (m *MockCodeService) UpdateCodeValue(codeValue *models.CodeValue) error {
	args := m.Called(codeValue)
	return args.Error(0)
}

func (m *MockCodeService) DeleteCodeType(codeTypeID int) error {
	args := m.Called(codeTypeID)
	return args.Error(0)
}

func (m *MockCodeService) DeleteCodeValue(codeValueID int) error {
	args := m.Called(codeValueID)
	return args.Error(0)
}

type CodePresentationTestSuite struct {
	suite.Suite
	CodeService presentations.CodePresentation
	cs          *MockCodeService
}

func (suite *CodePresentationTestSuite) SetupTest() {
	suite.cs = new(MockCodeService)
	suite.CodeService = presentations.NewCodePresentation(suite.cs)
}

func TestCodePresentationTestSuite(t *testing.T) {
	suite.Run(t, new(CodePresentationTestSuite))
}

func (suite *CodePresentationTestSuite) TestCreateCodeType() {
	// Arrange
	codeType := &models.CodeType{
		TypeName: "TestType",
		TypeDesc: "This is a test type",
	}

	suite.cs.On("CreateCodeType", codeType).Return(nil)

	// Act
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/code/type", suite.CodeService.CreateCodeType)

	// create a request to pass to the handler
	reqBody := `{"type_name": "TestType", "type_desc": "This is a test type"}`
	req := httptest.NewRequest("POST", "/code/type", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// create a response recorder
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), 200, w.Code)
}

func (suite *CodePresentationTestSuite) TestCreateCodeValue() {
	// Arrange
	codeValue := &models.CodeValue{
		CodeTypeID:    1,
		CodeValue:     "TestValue",
		CodeValueDesc: "This is a test value",
	}

	suite.cs.On("CreateCodeValue", codeValue).Return(nil)

	// Act
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/code/value", suite.CodeService.CreateCodeValue)

	// create a request to pass to the handler
	reqBody := `{"code_type_id": 1, "code_value": "TestValue", "code_value_desc": "This is a test value"}`
	req := httptest.NewRequest("POST", "/code/value", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// create a response recorder
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), 200, w.Code)
}

func (suite *CodePresentationTestSuite) TestGetAllCodeTypes() {
	// Arrange
	suite.cs.On("GetAllCodeTypes").Return([]models.CodeType{}, nil)

	// Act
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/code/types", suite.CodeService.GetAllCodeTypes)

	// create a request to pass to the handler
	req := httptest.NewRequest("GET", "/code/types", nil)

	// create a response recorder
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), 200, w.Code)
}

func (suite *CodePresentationTestSuite) TestGetCodeTypeByID() {
	// Arrange
	codeTypeID := 1
	suite.cs.On("GetCodeTypeByID", codeTypeID).Return(&models.CodeType{}, nil)

	// Act
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/code/type", suite.CodeService.GetCodeTypeByID)

	// create a request to pass to the handler
	req := httptest.NewRequest("GET", "/code/type?id=1", nil)

	// create a response recorder
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), 200, w.Code)
}

func (suite *CodePresentationTestSuite) TestUpdateCodeType() {
	// Arrange
	codeType := &models.CodeType{
		ID:       1,
		TypeName: "TestType",
		TypeDesc: "This is a test type",
	}

	suite.cs.On("UpdateCodeType", codeType).Return(nil)

	// Act
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.PUT("/code/type", suite.CodeService.UpdateCodeType)

	// create a request to pass to the handler
	reqBody := `{"id": 1, "type_name": "TestType", "type_desc": "This is a test type"}`
	req := httptest.NewRequest("PUT", "/code/type", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// create a response recorder
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), 200, w.Code)
}

func (suite *CodePresentationTestSuite) TestUpdateCodeValue() {
	// Arrange
	codeValue := &models.CodeValue{
		ID:            1,
		CodeTypeID:    1,
		CodeValue:     "TestValue",
		CodeValueDesc: "This is a test value",
	}

	suite.cs.On("UpdateCodeValue", codeValue).Return(nil)

	// Act
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.PUT("/code/value", suite.CodeService.UpdateCodeValue)

	// create a request to pass to the handler
	reqBody := `{"id": 1, "code_type_id": 1, "code_value": "TestValue", "code_value_desc": "This is a test value"}`
	req := httptest.NewRequest("PUT", "/code/value", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// create a response recorder
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), 200, w.Code)
}

func (suite *CodePresentationTestSuite) TestDeleteCodeType() {
	// Arrange
	codeValueID := 1
	codeValieIDStr := strconv.Itoa(codeValueID)
	suite.cs.On("DeleteCodeType", codeValueID).Return(nil)

	// Act
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.DELETE("/code/type", suite.CodeService.DeleteCodeType)

	// create a request to pass to the handler
	req := httptest.NewRequest("DELETE", "/code/type", nil)
	q := req.URL.Query()
	q.Add("id", codeValieIDStr)
	req.URL.RawQuery = q.Encode()
	req.Header.Set("Content-Type", "application/json")

	// create a response recorder
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), 200, w.Code)
}

func (suite *CodePresentationTestSuite) TestDeleteCodeValue() {
	// Arrange
	codeValueID := 1
	codeValueIDStr := strconv.Itoa(codeValueID)
	suite.cs.On("DeleteCodeValue", codeValueID).Return(nil)

	// Act
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.DELETE("/code/value", suite.CodeService.DeleteCodeValue)

	// create a request to pass to the handler
	req := httptest.NewRequest("DELETE", "/code/value", nil)
	q := req.URL.Query()
	q.Add("id", codeValueIDStr)
	req.URL.RawQuery = q.Encode()
	req.Header.Set("Content-Type", "application/json")

	// create a response recorder
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), 200, w.Code)
}
