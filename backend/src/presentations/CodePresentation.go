package presentations

import (
	"fmt"
	"meeting-center/src/models"
	"meeting-center/src/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CodePresentation interface {
	// C
	CreateCodeType(c *gin.Context)
	CreateCodeValue(c *gin.Context)

	// R
	GetAllCodeTypes(c *gin.Context)
	GetAllCodeValuesByType(c *gin.Context)

	// U
	UpdateCodeType(c *gin.Context)
	UpdateCodeValue(c *gin.Context)

	// D
	DeleteCodeType(c *gin.Context)
	DeleteCodeValue(c *gin.Context)
}

type codePresentation struct {
	cs services.CodeService
}

func NewCodePresentation(opt ...services.CodeService) CodePresentation {
	cs := services.NewCodeService()
	if len(opt) == 1 {
		cs = opt[0]
	}
	return &codePresentation{cs: cs}
}

func (cp *codePresentation) CreateCodeType(c *gin.Context) {
	var codeType models.CodeType
	if err := c.ShouldBindJSON(&codeType); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// filter out TypeName, TypeDesc
	filteredCodeType := models.CodeType{
		TypeName: codeType.TypeName,
		TypeDesc: codeType.TypeDesc,
	}

	if err := cp.cs.CreateCodeType(&filteredCodeType); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, filteredCodeType)
}

func (cp *codePresentation) CreateCodeValue(c *gin.Context) {
	var codeValue models.CodeValue
	if err := c.ShouldBindJSON(&codeValue); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// filter out CodeTypeID, CodeValue, CodeValueDesc
	filteredCodeValue := models.CodeValue{
		CodeTypeID:    codeValue.CodeTypeID,
		CodeValue:     codeValue.CodeValue,
		CodeValueDesc: codeValue.CodeValueDesc,
	}

	if err := cp.cs.CreateCodeValue(&filteredCodeValue); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, filteredCodeValue)
}

func (cp *codePresentation) GetAllCodeTypes(c *gin.Context) {
	codeTypes, err := cp.cs.GetAllCodeTypes()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, codeTypes)
}

func (cp *codePresentation) GetAllCodeValuesByType(c *gin.Context) {
	codeTypeID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	codeValues, err := cp.cs.GetAllCodeValuesByType(codeTypeID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, codeValues)
}

func (cp *codePresentation) UpdateCodeType(c *gin.Context) {
	var codeType models.CodeType
	if err := c.ShouldBindJSON(&codeType); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// filter out TypeName, TypeDesc
	filteredCodeType := models.CodeType{
		ID:       codeType.ID,
		TypeName: codeType.TypeName,
		TypeDesc: codeType.TypeDesc,
	}

	if err := cp.cs.UpdateCodeType(&filteredCodeType); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, codeType)
}

func (cp *codePresentation) UpdateCodeValue(c *gin.Context) {
	var codeValue models.CodeValue
	if err := c.ShouldBindJSON(&codeValue); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(codeValue)
	// filter out CodeTypeID, CodeValue, CodeValueDesc
	filteredCodeValue := models.CodeValue{
		ID:            codeValue.ID,
		CodeTypeID:    codeValue.CodeTypeID,
		CodeValue:     codeValue.CodeValue,
		CodeValueDesc: codeValue.CodeValueDesc,
	}
	if err := cp.cs.UpdateCodeValue(&filteredCodeValue); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, filteredCodeValue)
}

func (cp *codePresentation) DeleteCodeType(c *gin.Context) {
	codeTypeID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(codeTypeID)
	if err := cp.cs.DeleteCodeType(codeTypeID); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "CodeType deleted"})
}

func (cp *codePresentation) DeleteCodeValue(c *gin.Context) {
	codeValueID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := cp.cs.DeleteCodeValue(codeValueID); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "CodeValue deleted"})
}
