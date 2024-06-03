package test

import (
	"context"
	"meeting-center/src/domains"
	// "meeting-center/src/models"
	"mime/multipart"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockGcsRepo struct {
	mock.Mock
}

type GcsDomainTestSuite struct {
	suite.Suite
	gd domains.GcsDomain
	gr *MockGcsRepo
}

func (m *MockGcsRepo) UploadFile(ctx context.Context, file multipart.File, filename string) error {
	args := m.Called(ctx, file, filename)
	return args.Error(0)
}

func (m *MockGcsRepo) GetSignedURL(ctx context.Context, objectName string) (string, error) {
	args := m.Called(ctx, objectName)
	return args.String(0), args.Error(1)
}

func (m *MockGcsRepo) DeleteFile(ctx context.Context, filename string) error {
	args := m.Called(ctx, filename)
	return args.Error(0)
}

func (m *MockGcsRepo) NewGcsRepo() *MockGcsRepo {
	return &MockGcsRepo{}
}

func (suite *GcsDomainTestSuite) SetupTest() {
	suite.gr = suite.gr.NewGcsRepo()
	suite.gd = domains.NewGcsDomain(suite.gr)
}

func (suite *GcsDomainTestSuite) TestNewGcsDomain() {
	// 0 arguments
	// TODO: find a way to test when 0 arguments

	// 1 argument
	suite.gr.On("NewGcsDomain", suite.gr).Return(nil)
	gd := domains.NewGcsDomain(suite.gr)
	assert.NotNil(suite.T(), gd)

	// // 2 arguments
	assert.Panics(suite.T(), func() { domains.NewGcsDomain(suite.gr, suite.gr) })
}

func (suite *GcsDomainTestSuite) TestUploadFile() {
	ctx := context.Background()
	file := new(multipart.File)
	objID := "test"
	objExt := ".txt"
	suite.gr.On("UploadFile", ctx, *file, objID+objExt).Return(nil)
	err := suite.gd.UploadFile(ctx, *file, objID, objExt)
	assert.NoError(suite.T(), err)
}

func (suite *GcsDomainTestSuite) TestGetSignedURL() {
	ctx := context.Background()
	objID := "test"
	objExt := ".txt"
	suite.gr.On("GetSignedURL", ctx, objID+objExt).Return("url", nil)
	url, err := suite.gd.GetSignedURL(ctx, objID, objExt)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "url", url)
}

func (suite *GcsDomainTestSuite) TestDeleteFile() {
	ctx := context.Background()
	objID := "test"
	objExt := ".txt"
	suite.gr.On("DeleteFile", ctx, objID+objExt).Return(nil)
	err := suite.gd.DeleteFile(ctx, objID, objExt)
	assert.NoError(suite.T(), err)
}

func TestGcsDomainTestSuite(t *testing.T) {
	suite.Run(t, new(GcsDomainTestSuite))
}
