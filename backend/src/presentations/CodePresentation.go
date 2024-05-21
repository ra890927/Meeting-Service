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
	GetCodeTypeByID(c *gin.Context)
	GetCodeValueByID(c *gin.Context)

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

type CreateCodeTypeParam struct {
	TypeName string `json:"type_name" binding:"required"`
	TypeDesc string `json:"type_desc" binding:"required"`
}
type CreateCodeTypeResponse struct {
	Status string `json:"status"`
	Data   struct {
		CodeType struct {
			ID         int                `json:"id"`
			TypeName   string             `json:"type_name"`
			TypeDesc   string             `json:"type_desc"`
			CodeValues []models.CodeValue `json:"code_values"`
		} `json:"code_type"`
	} `json:"data"`
	Message string `json:"message"`
}
type UpdateCodeValueParam struct {
	ID            int    `json:"id" binding:"required"`
	CodeTypeID    int    `json:"code_type_id" binding:"required"`
	CodeValue     string `json:"code_value" binding:"required"`
	CodeValueDesc string `json:"code_value_desc" binding:"required"`
}

type UpdateCodeValueResponse struct {
	Status string `json:"status"`
	Data   struct {
		CodeValue struct {
			ID            int    `json:"id"`
			CodeTypeID    int    `json:"code_type_id"`
			CodeValue     string `json:"code_value"`
			CodeValueDesc string `json:"code_value_desc"`
		} `json:"code_value"`
	} `json:"data"`
	Message string `json:"message"`
}

type DeleteCodeTypeResponse struct {
	Status  string   `json:"status"`
	Data    struct{} `json:"data"`
	Message string   `json:"message"`
}

type DeleteCodeValueResponse struct {
	Status  string   `json:"status"`
	Data    struct{} `json:"data"`
	Message string   `json:"message"`
}

type CreateCodeValueParam struct {
	CodeTypeID    int    `json:"code_type_id" binding:"required"`
	CodeValue     string `json:"code_value" binding:"required"`
	CodeValueDesc string `json:"code_value_desc" binding:"required"`
}

type CreateCodeValueResponse struct {
	Status string `json:"status"`
	Data   struct {
		CodeValue struct {
			ID            int    `json:"id"`
			CodeTypeID    int    `json:"code_type_id"`
			CodeValue     string `json:"code_value"`
			CodeValueDesc string `json:"code_value_desc"`
		} `json:"code_value"`
	} `json:"data"`
	Message string `json:"message"`
}

type UpdateCodeTypeParam struct {
	ID       int    `json:"id" binding:"required"`
	TypeName string `json:"type_name" binding:"required"`
	TypeDesc string `json:"type_desc" binding:"required"`
}

type UpdateCodeTypeResponse struct {
	Status string `json:"status"`
	Data   struct {
		CodeType models.CodeType `json:"code_type"`
	} `json:"data"`
	Message string `json:"message"`
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

// @Summary Create a new code type
// @Description Create a new code type
// @Tags code
// @Accept json
// @Produce json
// @Param codeType body CreateCodeTypeParam true "CodeType"
// @Success 200 {object} CreateCodeTypeResponse
// @Router /code/type [post]
func (cp codePresentation) CreateCodeType(c *gin.Context) {
	var codeTypeParam CreateCodeTypeParam
	var response CreateCodeTypeResponse
	if err := c.ShouldBindJSON(&codeTypeParam); err != nil {
		response.Status = "fail"
		response.Message = err.Error()
		c.JSON(400, response)
		return
	}
	// filter out TypeName, TypeDesc
	filteredCodeType := models.CodeType{
		TypeName: codeTypeParam.TypeName,
		TypeDesc: codeTypeParam.TypeDesc,
	}

	if err := cp.cs.CreateCodeType(&filteredCodeType); err != nil {
		response.Status = "fail"
		response.Message = err.Error()
		c.JSON(500, response)
		return
	}

	response.Status = "success"
	response.Data.CodeType.ID = filteredCodeType.ID
	response.Data.CodeType.TypeName = filteredCodeType.TypeName
	response.Data.CodeType.TypeDesc = filteredCodeType.TypeDesc
	response.Data.CodeType.CodeValues = []models.CodeValue{}
	response.Message = "CodeType created"

	c.JSON(200, response)
}

// @Summary Create a new code value
// @Description Create a new code value
// @Tags code
// @Accept json
// @Produce json
// @Param codeValue body CreateCodeValueParam true "CodeValue"
// @Success 200 {object} CreateCodeValueResponse
// @Router /code/value [post]
func (cp codePresentation) CreateCodeValue(c *gin.Context) {
	var codeValueParam CreateCodeValueParam
	var response CreateCodeValueResponse
	if err := c.ShouldBindJSON(&codeValueParam); err != nil {
		response.Status = "fail"
		response.Message = err.Error()
		c.JSON(400, response)
		return
	}
	// filter out CodeTypeID, CodeValue, CodeValueDesc
	filteredCodeValue := models.CodeValue{
		CodeTypeID:    codeValueParam.CodeTypeID,
		CodeValue:     codeValueParam.CodeValue,
		CodeValueDesc: codeValueParam.CodeValueDesc,
	}

	if err := cp.cs.CreateCodeValue(&filteredCodeValue); err != nil {
		response.Status = "fail"
		response.Message = err.Error()
		c.JSON(500, response)
		return
	}

	response.Status = "success"
	response.Data.CodeValue.ID = filteredCodeValue.ID
	response.Data.CodeValue.CodeTypeID = filteredCodeValue.CodeTypeID
	response.Data.CodeValue.CodeValue = filteredCodeValue.CodeValue
	response.Data.CodeValue.CodeValueDesc = filteredCodeValue.CodeValueDesc
	response.Message = "CodeValue created"
	c.JSON(200, response)
}

type GetAllCodeTypesResponse struct {
	Status string `json:"status"`
	Data   struct {
		CodeTypes []models.CodeType `json:"code_types"`
	} `json:"data"`
	Message string `json:"message"`
}

// @Summary Get all code types
// @Description Get all code types
// @Tags code
// @Accept json
// @Produce json
// @Success 200 {object} GetAllCodeTypesResponse
// @Router /code/type/getAllCodeTypes [get]
func (cp codePresentation) GetAllCodeTypes(c *gin.Context) {
	var response GetAllCodeTypesResponse
	codeTypes, err := cp.cs.GetAllCodeTypes()
	if err != nil {
		response.Status = "fail"
		response.Message = err.Error()
		c.JSON(500, response)
		return
	}
	response.Status = "success"
	response.Data.CodeTypes = codeTypes
	response.Message = "CodeTypes fetched"
	c.JSON(200, response)
}

type GetCodeTypeByIDResponse struct {
	Status string `json:"status"`
	Data   struct {
		CodeType models.CodeType `json:"code_type"`
	} `json:"data"`
	Message string `json:"message"`
}

// @Summary Get a code type by ID
// @Description Get a code type by ID
// @Tags code
// @Accept json
// @Produce json
// @Param id query int true "CodeType ID"
// @Success 200 {object} CreateCodeTypeResponse
// @Router /code/type/getCodeTypeByID [get]
func (cp codePresentation) GetCodeTypeByID(c *gin.Context) {
	codeTypeID, err := strconv.Atoi(c.Query("id"))
	var response GetCodeTypeByIDResponse
	if err != nil {
		response.Status = "fail"
		response.Message = err.Error()
		c.JSON(400, response)
		return
	}
	codeType, err := cp.cs.GetCodeTypeByID(codeTypeID)
	if err != nil {
		response.Status = "fail"
		response.Message = err.Error()
		c.JSON(500, response)
		return
	}

	response.Status = "success"
	response.Data.CodeType = *codeType
	response.Message = "CodeType fetched"
	c.JSON(200, response)
}

type GetCodeValueByIDResponse struct {
	Status string `json:"status"`
	Data   struct {
		CodeValue models.CodeValue `json:"code_value"`
	} `json:"data"`
	Message string `json:"message"`
}

// @Summary Get a code value by ID
// @Description Get a code value by ID
// @Tags code
// @Accept json
// @Produce json
// @Param id query int true "CodeValue ID"
// @Success 200 {object} GetCodeValueByIDResponse
// @Router /code/value/getCodeValueByID [get]
func (cp codePresentation) GetCodeValueByID(c *gin.Context) {
	codeValueID, err := strconv.Atoi(c.Query("id"))
	var response GetCodeValueByIDResponse
	if err != nil {
		response.Status = "fail"
		response.Message = err.Error()
		c.JSON(400, response)
		return
	}
	codeValue, err := cp.cs.GetCodeValueByID(codeValueID)
	if err != nil {
		response.Status = "fail"
		response.Message = err.Error()
		c.JSON(500, response)
		return
	}

	response.Status = "success"
	response.Data.CodeValue = *codeValue
	response.Message = "CodeValue fetched"
	c.JSON(200, response)
}

// @Summary Update a code type
// @Description Update a code type
// @Tags code
// @Accept json
// @Produce json
// @Param codeType body UpdateCodeTypeParam true "CodeType"
// @Success 200 {object} UpdateCodeTypeResponse
// @Router /code/type [put]
func (cp codePresentation) UpdateCodeType(c *gin.Context) {
	var codeTypeParam UpdateCodeTypeParam
	var response UpdateCodeTypeResponse
	if err := c.ShouldBindJSON(&codeTypeParam); err != nil {
		response.Status = "fail"
		response.Message = err.Error()
		c.JSON(400, response)
		return
	}
	// filter out TypeName, TypeDesc
	filteredCodeType := models.CodeType{
		ID:       codeTypeParam.ID,
		TypeName: codeTypeParam.TypeName,
		TypeDesc: codeTypeParam.TypeDesc,
	}

	if err := cp.cs.UpdateCodeType(&filteredCodeType); err != nil {
		response.Status = "fail"
		response.Message = err.Error()
		c.JSON(500, response)
		return
	}

	response.Status = "success"
	response.Data.CodeType = filteredCodeType
	response.Message = "CodeType updated"

	c.JSON(200, response)
}

// @Summary Update a code value
// @Description Update a code value
// @Tags code
// @Accept json
// @Produce json
// @Param codeValue body UpdateCodeValueParam true "CodeValue"
// @Success 200 {object} UpdateCodeValueResponse
// @Router /code/value [put]
func (cp codePresentation) UpdateCodeValue(c *gin.Context) {
	var codeValueParam UpdateCodeValueParam
	var response UpdateCodeValueResponse
	if err := c.ShouldBindJSON(&codeValueParam); err != nil {
		response.Status = "fail"
		response.Message = err.Error()
		c.JSON(400, response)
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
		response.Status = "fail"
		response.Message = err.Error()
		c.JSON(500, response)
		return
	}

	response.Status = "success"
	response.Data.CodeValue.ID = filteredCodeValue.ID
	response.Data.CodeValue.CodeTypeID = filteredCodeValue.CodeTypeID
	response.Data.CodeValue.CodeValue = filteredCodeValue.CodeValue
	response.Data.CodeValue.CodeValueDesc = filteredCodeValue.CodeValueDesc
	response.Message = "CodeValue updated"

	c.JSON(200, response)
}

// @Summary Delete a code type
// @Description Delete a code type
// @Tags code
// @Accept json
// @Produce json
// @Param id query int true "CodeType ID"
// @Success 200 {object} DeleteCodeTypeResponse
// @Router /code/type [delete]
func (cp codePresentation) DeleteCodeType(c *gin.Context) {
	codeTypeID, err := strconv.Atoi(c.Query("id"))
	var response DeleteCodeTypeResponse
	if err != nil {
		response.Status = "fail"
		response.Message = err.Error()
		c.JSON(400, response)
		return
	}
	if err := cp.cs.DeleteCodeType(codeTypeID); err != nil {
		response.Status = "fail"
		response.Message = err.Error()
		c.JSON(500, response)
		return
	}
	response.Status = "success"
	response.Message = fmt.Sprintf("CodeType with ID %d deleted", codeTypeID)
	c.JSON(200, response)
}

// @Summary Delete a code value
// @Description Delete a code value
// @Tags code
// @Accept json
// @Produce json
// @Param id query int true "CodeValue ID"
// @Success 200 {object} DeleteCodeValueResponse
// @Router /code/value [delete]
func (cp codePresentation) DeleteCodeValue(c *gin.Context) {
	codeValueID, err := strconv.Atoi(c.Query("id"))
	var response DeleteCodeValueResponse
	if err != nil {
		response.Status = "fail"
		response.Message = err.Error()
		c.JSON(400, response)
		return
	}
	if err := cp.cs.DeleteCodeValue(codeValueID); err != nil {
		response.Status = "fail"
		response.Message = err.Error()
		c.JSON(500, response)
		return
	}
	response.Status = "success"
	response.Message = fmt.Sprintf("CodeValue with ID %d deleted", codeValueID)
	c.JSON(200, response)
}
