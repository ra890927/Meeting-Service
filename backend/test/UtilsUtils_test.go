package test

import (
	"meeting-center/src/models"
	"meeting-center/src/utils"
	"testing"

	"reflect"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UtilsUtilsTestSuite struct {
	suite.Suite
}

func (suite *UtilsUtilsTestSuite) SetupTest() {
}

func (suite *UtilsUtilsTestSuite) TestIsEmptyValue() {
	// reflect.Int
	intA := int(0)
	assert.True(suite.T(), utils.IsEmptyValue(reflect.ValueOf(intA)))

	// reflect.Slice
	sliceA := []int{}
	assert.True(suite.T(), utils.IsEmptyValue(reflect.ValueOf(sliceA)))

	// None of the data types above
	errorA := error(nil)
	assert.False(suite.T(), utils.IsEmptyValue(reflect.ValueOf(errorA)))
}

func (suite *UtilsUtilsTestSuite) TestOverwriteValue() {
	// models.User
	userA := &models.User{}
	userB := &models.User{}
	utils.OverwriteValue(userA, userB)
	assert.Equal(suite.T(), userA, userB)

	// not pointer
	userC := models.User{}
	userD := models.User{}
	assert.Panics(suite.T(), func() { utils.OverwriteValue(userC, userD) })
}

func TestUtilsUtilsTestSuite(t *testing.T) {
	suite.Run(t, new(UtilsUtilsTestSuite))
}
