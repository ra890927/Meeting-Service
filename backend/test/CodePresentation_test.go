package test

import (
	"bytes"
	"fmt"
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

func (suite *CodePresentationTestSuite) TestNewCodePresentation() {

	mockService := new(MockCodeService)
	cp := presentations.NewCodePresentation(mockService)
	assert.NotNil(suite.T(), cp)

	// cp = presentations.NewCodePresentation()
	// assert.NotNil(suite.T(), cp)

	assert.Panics(suite.T(), func() {
		presentations.NewCodePresentation(mockService, mockService)
	})
}

func (suite *CodePresentationTestSuite) TestCreateCodeType() {
	// Arrange
	codeType := &models.CodeType{
		TypeName: "TestType",
		TypeDesc: "This is a test type",
	}

	suite.cs.On("CreateCodeType", codeType).Return(nil).Once()

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

	suite.cs.On("CreateCodeType", codeType).Return(fmt.Errorf("some error")).Once()

	// create a request to pass to the handler
	reqBody = `{"type_name": "TestType", "type_desc": "This is a test type"}`
	req = httptest.NewRequest("POST", "/code/type", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// create a response recorder
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), 500, w.Code)

	// create a request to pass to the handler with invalid JSON
	reqBody = `{"type_name": "", "type_desc": "This is a test type"}`
	req = httptest.NewRequest("POST", "/code/type", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// create a response recorder
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), 400, w.Code)
}

func (suite *CodePresentationTestSuite) TestCreateCodeValue() {
	// Arrange
	codeValue := &models.CodeValue{
		CodeTypeID:    1,
		CodeValue:     "TestValue",
		CodeValueDesc: "This is a test value",
	}

	suite.cs.On("CreateCodeValue", codeValue).Return(nil).Once()

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

	suite.cs.On("CreateCodeValue", codeValue).Return(fmt.Errorf("some error")).Once()

	// create a request to pass to the handler
	reqBody = `{"code_type_id": 1, "code_value": "TestValue", "code_value_desc": "This is a test value"}`
	req = httptest.NewRequest("POST", "/code/value", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// create a response recorder
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), 500, w.Code)

	// create a request to pass to the handler with invalid JSON
	reqBody = `{"code_type_id": 1, "code_value": "", "code_value_desc": "This is a test value"}`
	req = httptest.NewRequest("POST", "/code/value", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// create a response recorder
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), 400, w.Code)
}

func (suite *CodePresentationTestSuite) TestGetAllCodeTypes() {
	// Arrange
	codeTypes := []models.CodeType{
		{
			ID:       1,
			TypeName: "Type1",
			TypeDesc: "Description1",
		},
		{
			ID:       2,
			TypeName: "Type2",
			TypeDesc: "Description2",
		},
	}

	suite.cs.On("GetAllCodeTypes").Return(codeTypes, nil).Once()

	// Act
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/code/types", suite.CodeService.GetAllCodeTypes)

	// create a request to pass to the handler
	req := httptest.NewRequest("GET", "/code/types", nil)
	req.Header.Set("Content-Type", "application/json")

	// create a response recorder
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), 200, w.Code)

	suite.cs.On("GetAllCodeTypes").Return([]models.CodeType{}, fmt.Errorf("some error")).Once()

	// create a request to pass to the handler
	req = httptest.NewRequest("GET", "/code/types", nil)
	req.Header.Set("Content-Type", "application/json")

	// create a response recorder
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), 500, w.Code)
}

func (suite *CodePresentationTestSuite) TestGetCodeValueByID() {
	// Arrange
	codeValueID := 1
	codeValue := &models.CodeValue{
		ID:            codeValueID,
		CodeTypeID:    1,
		CodeValue:     "TestValue",
		CodeValueDesc: "This is a test value",
	}

	suite.cs.On("GetCodeValueByID", codeValueID).Return(codeValue, nil).Once()

	// Act
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/code/value", suite.CodeService.GetCodeValueByID)

	// create a request to pass to the handler
	req := httptest.NewRequest("GET", "/code/value?id=1", nil)
	req.Header.Set("Content-Type", "application/json")

	// create a response recorder
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), 200, w.Code)

	req = httptest.NewRequest("GET", "/code/value?id=invalid", nil)
	req.Header.Set("Content-Type", "application/json")

	// create a response recorder
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), 400, w.Code)

	suite.cs.On("GetCodeValueByID", codeValueID).Return(nil, fmt.Errorf("some error")).Once()

	// create a request to pass to the handler
	req = httptest.NewRequest("GET", "/code/value?id=1", nil)
	req.Header.Set("Content-Type", "application/json")

	// create a response recorder
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), 500, w.Code)
}

func (suite *CodePresentationTestSuite) TestGetCodeTypeByID() {
	// Arrange
	codeTypeID := 1
	codeType := &models.CodeType{
		ID:       codeTypeID,
		TypeName: "TestType",
		TypeDesc: "This is a test type",
	}

	suite.cs.On("GetCodeTypeByID", codeTypeID).Return(codeType, nil).Once()

	// Act
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/code/type", suite.CodeService.GetCodeTypeByID)

	// create a request to pass to the handler
	req := httptest.NewRequest("GET", "/code/type?id=1", nil)
	req.Header.Set("Content-Type", "application/json")

	// create a response recorder
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), 200, w.Code)

	req = httptest.NewRequest("GET", "/code/type?id=invalid", nil)
	req.Header.Set("Content-Type", "application/json")

	// create a response recorder
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), 400, w.Code)

	suite.cs.On("GetCodeTypeByID", codeTypeID).Return(nil, fmt.Errorf("some error")).Once()

	// create a request to pass to the handler
	req = httptest.NewRequest("GET", "/code/type?id=1", nil)
	req.Header.Set("Content-Type", "application/json")

	// create a response recorder
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), 500, w.Code)
}

func (suite *CodePresentationTestSuite) TestUpdateCodeType() {
	// Arrange
	codeType := &models.CodeType{
		ID:       1,
		TypeName: "TestType",
		TypeDesc: "This is a test type",
	}

	suite.cs.On("UpdateCodeType", codeType).Return(nil).Once()

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

	suite.cs.On("UpdateCodeType", codeType).Return(fmt.Errorf("some error")).Once()

	// create a request to pass to the handler
	reqBody = `{"id": 1, "type_name": "TestType", "type_desc": "This is a test type"}`
	req = httptest.NewRequest("PUT", "/code/type", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// create a response recorder
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), 500, w.Code)

	// create a request to pass to the handler with invalid JSON
	reqBody = `{"id": 1, "type_name": "", "type_desc": "This is a test type"}`
	req = httptest.NewRequest("PUT", "/code/type", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// create a response recorder
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), 400, w.Code)
}

func (suite *CodePresentationTestSuite) TestUpdateCodeValue() {
	// Arrange
	codeValue := &models.CodeValue{
		ID:            1,
		CodeTypeID:    1,
		CodeValue:     "TestValue",
		CodeValueDesc: "This is a test value",
	}

	suite.cs.On("UpdateCodeValue", codeValue).Return(nil).Once()

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

	suite.cs.On("UpdateCodeValue", codeValue).Return(fmt.Errorf("some error")).Once()

	// create a request to pass to the handler
	reqBody = `{"id": 1, "code_type_id": 1, "code_value": "TestValue", "code_value_desc": "This is a test value"}`
	req = httptest.NewRequest("PUT", "/code/value", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// create a response recorder
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), 500, w.Code)

	// create a request to pass to the handler with invalid JSON
	reqBody = `{"id": 1, "code_type_id": 1, "code_value": "", "code_value_desc": "This is a test value"}`
	req = httptest.NewRequest("PUT", "/code/value", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// create a response recorder
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), 400, w.Code)
}

func (suite *CodePresentationTestSuite) TestDeleteCodeType() {
	// Arrange
	codeTypeID := 1
	codeTypeIDStr := strconv.Itoa(codeTypeID)

	suite.cs.On("DeleteCodeType", codeTypeID).Return(nil).Once()

	// Act
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.DELETE("/code/type", suite.CodeService.DeleteCodeType)

	// create a request to pass to the handler
	req := httptest.NewRequest("DELETE", "/code/type?id="+codeTypeIDStr, nil)
	req.Header.Set("Content-Type", "application/json")

	// create a response recorder
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), 200, w.Code)

	suite.cs.On("DeleteCodeType", codeTypeID).Return(fmt.Errorf("some error")).Once()

	// create a request to pass to the handler
	req = httptest.NewRequest("DELETE", "/code/type?id="+codeTypeIDStr, nil)
	req.Header.Set("Content-Type", "application/json")

	// create a response recorder
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), 500, w.Code)

	req = httptest.NewRequest("DELETE", "/code/type?id=invalid", nil)
	req.Header.Set("Content-Type", "application/json")

	// create a response recorder
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), 400, w.Code)
}

func (suite *CodePresentationTestSuite) TestDeleteCodeValue() {
	// Arrange
	codeValueID := 1
	codeValueIDStr := strconv.Itoa(codeValueID)

	suite.cs.On("DeleteCodeValue", codeValueID).Return(nil).Once()

	// Act
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.DELETE("/code/value", suite.CodeService.DeleteCodeValue)

	// create a request to pass to the handler
	req := httptest.NewRequest("DELETE", "/code/value?id="+codeValueIDStr, nil)
	req.Header.Set("Content-Type", "application/json")

	// create a response recorder
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), 200, w.Code)

	suite.cs.On("DeleteCodeValue", codeValueID).Return(fmt.Errorf("some error")).Once()

	// create a request to pass to the handler
	req = httptest.NewRequest("DELETE", "/code/value?id="+codeValueIDStr, nil)
	req.Header.Set("Content-Type", "application/json")

	// create a response recorder
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), 500, w.Code)

	req = httptest.NewRequest("DELETE", "/code/value?id=invalid", nil)
	req.Header.Set("Content-Type", "application/json")

	// create a response recorder
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), 400, w.Code)
}
