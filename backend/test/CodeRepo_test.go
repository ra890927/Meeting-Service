package test

import (
	"meeting-center/src/models"
	. "meeting-center/src/repos"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type CodeRepoTestSuite struct {
	suite.Suite
	cr CodeRepo
	db *gorm.DB
}

func (suite *CodeRepoTestSuite) SetupTest() {
	db, err := gorm.Open(sqlite.Open("./test.sqlite"), &gorm.Config{})
	assert.NoError(suite.T(), err)

	err = db.AutoMigrate(&models.CodeType{}, &models.CodeValue{})
	assert.NoError(suite.T(), err)

	suite.db = db
	suite.cr = NewCodeRepo(db)
}

func (suite *CodeRepoTestSuite) TearDownTest() {
	db, err := suite.db.DB()
	assert.NoError(suite.T(), err)
	err = db.Close()
	assert.NoError(suite.T(), err)
}

func (suite *CodeRepoTestSuite) TestCreateCodeType() {
	codeType := &models.CodeType{
		TypeName: "TestCreateCodeType",
		TypeDesc: "This is a test type",
	}
	err := suite.cr.CreateCodeType(codeType)
	assert.NoError(suite.T(), err)
	assert.NotZero(suite.T(), codeType.ID)
}

func (suite *CodeRepoTestSuite) TestCreateCodeValue() {
	codeType := &models.CodeType{
		TypeName: "TestCreateCodeValue",
		TypeDesc: "This is a test type",
	}
	err := suite.cr.CreateCodeType(codeType)
	assert.NoError(suite.T(), err)

	codeValue := &models.CodeValue{
		CodeTypeID:    codeType.ID,
		CodeValue:     "TestValue",
		CodeValueDesc: "This is a test value",
	}
	err = suite.cr.CreateCodeValue(codeValue)
	assert.NoError(suite.T(), err)
	assert.NotZero(suite.T(), codeValue.ID)
}

func (suite *CodeRepoTestSuite) TestGetAllCodeTypes() {
	// check the original length of the codeTypes
	codeTypes, err := suite.cr.GetAllCodeTypes()
	assert.NoError(suite.T(), err)
	originalLength := len(codeTypes)

	codeType := &models.CodeType{
		TypeName: "TestGetAllCodeTypes",
		TypeDesc: "This is a test type",
	}
	err = suite.cr.CreateCodeType(codeType)
	assert.NoError(suite.T(), err)

	codeTypes, err = suite.cr.GetAllCodeTypes()
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), codeTypes, originalLength+1)
}

func (suite *CodeRepoTestSuite) TestGetCodeTypeByID() {
	codeType := &models.CodeType{
		TypeName: "TestGetCodeTypeByID",
		TypeDesc: "This is a test type",
	}
	err := suite.cr.CreateCodeType(codeType)
	assert.NoError(suite.T(), err)

	foundCodeType, err := suite.cr.GetCodeTypeByID(codeType.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), codeType.ID, foundCodeType.ID)
}

func (suite *CodeRepoTestSuite) TestGetCodeValueByID() {
	codeType := &models.CodeType{
		TypeName: "TestGetCodeValueByID",
		TypeDesc: "This is a test type",
	}
	err := suite.cr.CreateCodeType(codeType)
	assert.NoError(suite.T(), err)

	codeValue := &models.CodeValue{
		CodeTypeID:    codeType.ID,
		CodeValue:     "TestValue",
		CodeValueDesc: "This is a test value",
	}
	err = suite.cr.CreateCodeValue(codeValue)
	assert.NoError(suite.T(), err)

	foundCodeValue, err := suite.cr.GetCodeValueByID(codeValue.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), codeValue.ID, foundCodeValue.ID)
}

func (suite *CodeRepoTestSuite) TestUpdateCodeType() {
	codeType := &models.CodeType{
		TypeName: "TestUpdateCodeType",
		TypeDesc: "This is a test type",
	}
	err := suite.cr.CreateCodeType(codeType)
	assert.NoError(suite.T(), err)

	codeType.TypeDesc = "Updated test type"
	err = suite.cr.UpdateCodeType(codeType)
	assert.NoError(suite.T(), err)
}

func (suite *CodeRepoTestSuite) TestUpdateCodeValue() {
	codeType := &models.CodeType{
		TypeName: "TestUpdateCodeValue",
		TypeDesc: "This is a test type",
	}
	err := suite.cr.CreateCodeType(codeType)
	assert.NoError(suite.T(), err)

	codeValue := &models.CodeValue{
		CodeTypeID:    codeType.ID,
		CodeValue:     "TestValue",
		CodeValueDesc: "This is a test value",
	}
	err = suite.cr.CreateCodeValue(codeValue)
	assert.NoError(suite.T(), err)

	codeValue.CodeValueDesc = "Updated test value"
	err = suite.cr.UpdateCodeValue(codeValue)
	assert.NoError(suite.T(), err)
}

func (suite *CodeRepoTestSuite) TestDeleteCodeValue() {
	codeType := &models.CodeType{
		TypeName: "TestDeleteCodeValue",
		TypeDesc: "This is a test type",
	}
	err := suite.cr.CreateCodeType(codeType)
	assert.NoError(suite.T(), err)

	codeValue := &models.CodeValue{
		CodeTypeID:    codeType.ID,
		CodeValue:     "TestValue",
		CodeValueDesc: "This is a test value",
	}
	err = suite.cr.CreateCodeValue(codeValue)
	assert.NoError(suite.T(), err)

	err = suite.cr.DeleteCodeValue(codeValue.ID)
	assert.NoError(suite.T(), err)
}

func (suite *CodeRepoTestSuite) TestDeleteCodeType() {
	codeType := &models.CodeType{
		TypeName: "TestDeleteCodeType",
		TypeDesc: "This is a test type",
	}
	err := suite.cr.CreateCodeType(codeType)
	assert.NoError(suite.T(), err)

	err = suite.cr.DeleteCodeType(codeType.ID)
	assert.NoError(suite.T(), err)
}

func TestCodeRepoTestSuite(t *testing.T) {
	suite.Run(t, new(CodeRepoTestSuite))
}
