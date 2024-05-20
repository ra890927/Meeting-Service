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

func NewCodePresentation(codeServiceArgs ...services.CodeService) CodePresentation {
	if len(codeServiceArgs) == 1 {
		return CodePresentation(&codePresentation{cs: codeServiceArgs[0]})
	} else if len(codeServiceArgs) == 0 {
		return CodePresentation(&codePresentation{cs: services.NewCodeService()})
	} else {
		panic("NewCodePresentation: too many arguments")
	}
}

type CreateCodeTypeParam struct {
	TypeName string `json:"typeName" binding:"required"`
	TypeDesc string `json:"typeDesc" binding:"required"`
}
type CreateCodeTypeResponse struct {
	ID       int    `json:"id"`
	TypeName string `json:"typeName"`
	TypeDesc string `json:"typeDesc"`
}

// @Summary Create a new code type
// @Description Create a new code type
// @Tags code
// @Accept json
// @Produce json
// @Param codeType body CreateCodeTypeParam true "CodeType"
// @Success 200 {object} models.CodeType
// @Router /code/type [post]
func (cp codePresentation) CreateCodeType(c *gin.Context) {
	var codeTypeParam CreateCodeTypeParam
	if err := c.ShouldBindJSON(&codeTypeParam); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// filter out TypeName, TypeDesc
	filteredCodeType := models.CodeType{
		TypeName: codeTypeParam.TypeName,
		TypeDesc: codeTypeParam.TypeDesc,
	}

	if err := cp.cs.CreateCodeType(&filteredCodeType); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, filteredCodeType)
}

type CreateCodeValueParam struct {
	CodeTypeID    int    `json:"codeTypeID" binding:"required"`
	CodeValue     string `json:"codeValue" binding:"required"`
	CodeValueDesc string `json:"codeValueDesc" binding:"required"`
}

// @Summary Create a new code value
// @Description Create a new code value
// @Tags code
// @Accept json
// @Produce json
// @Param codeValue body CreateCodeValueParam true "CodeValue"
// @Success 200 {object} models.CodeValue
// @Router /code/value [post]
func (cp codePresentation) CreateCodeValue(c *gin.Context) {
	var codeValueParam CreateCodeValueParam
	if err := c.ShouldBindJSON(&codeValueParam); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// filter out CodeTypeID, CodeValue, CodeValueDesc
	filteredCodeValue := models.CodeValue{
		CodeTypeID:    codeValueParam.CodeTypeID,
		CodeValue:     codeValueParam.CodeValue,
		CodeValueDesc: codeValueParam.CodeValueDesc,
	}

	if err := cp.cs.CreateCodeValue(&filteredCodeValue); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, filteredCodeValue)
}

// @Summary Get all code types
// @Description Get all code types
// @Tags code
// @Accept json
// @Produce json
// @Success 200 {object} []models.CodeType
// @Router /code/type/getAllCodeTypes [get]
func (cp codePresentation) GetAllCodeTypes(c *gin.Context) {
	codeTypes, err := cp.cs.GetAllCodeTypes()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, codeTypes)
}

func (cp codePresentation) GetAllCodeValuesByType(c *gin.Context) {
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

type UpdateCodeTypeParam struct {
	ID       int    `json:"id" binding:"required"`
	TypeName string `json:"typeName" binding:"required"`
	TypeDesc string `json:"typeDesc" binding:"required"`
}

// @Summary Update a code type
// @Description Update a code type
// @Tags code
// @Accept json
// @Produce json
// @Param codeType body UpdateCodeTypeParam true "CodeType"
// @Success 200 {object} models.CodeType
// @Router /code/type [put]
func (cp codePresentation) UpdateCodeType(c *gin.Context) {
	var codeTypeParam UpdateCodeTypeParam
	if err := c.ShouldBindJSON(&codeTypeParam); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// filter out TypeName, TypeDesc
	filteredCodeType := models.CodeType{
		ID:       codeTypeParam.ID,
		TypeName: codeTypeParam.TypeName,
		TypeDesc: codeTypeParam.TypeDesc,
	}

	if err := cp.cs.UpdateCodeType(&filteredCodeType); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, filteredCodeType)
}

type UpdateCodeValueParam struct {
	ID            int    `json:"id" binding:"required"`
	CodeTypeID    int    `json:"codeTypeID" binding:"required"`
	CodeValue     string `json:"codeValue" binding:"required"`
	CodeValueDesc string `json:"codeValueDesc" binding:"required"`
}

// @Summary Update a code value
// @Description Update a code value
// @Tags code
// @Accept json
// @Produce json
// @Param codeValue body UpdateCodeValueParam true "CodeValue"
// @Success 200 {object} models.CodeValue
// @Router /code/value [put]
func (cp codePresentation) UpdateCodeValue(c *gin.Context) {
	var codeValueParam UpdateCodeValueParam
	if err := c.ShouldBindJSON(&codeValueParam); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// filter out CodeTypeID, CodeValue, CodeValueDesc
	filteredCodeValue := models.CodeValue{
		ID:            codeValueParam.ID,
		CodeTypeID:    codeValueParam.CodeTypeID,
		CodeValue:     codeValueParam.CodeValue,
		CodeValueDesc: codeValueParam.CodeValueDesc,
	}

	if err := cp.cs.UpdateCodeValue(&filteredCodeValue); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, filteredCodeValue)
}

// @Summary Delete a code type
// @Description Delete a code type
// @Tags code
// @Accept json
// @Produce json
// @Param id query int true "CodeType ID"
// @Success 200 {object} string
// @Router /code/type [delete]
func (cp codePresentation) DeleteCodeType(c *gin.Context) {
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

// @Summary Delete a code value
// @Description Delete a code value
// @Tags code
// @Accept json
// @Produce json
// @Param id query int true "CodeValue ID"
// @Success 200 {object} string
// @Router /code/value [delete]
func (cp codePresentation) DeleteCodeValue(c *gin.Context) {
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
