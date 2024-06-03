package presentations

import (
	"meeting-center/src/models"
	"meeting-center/src/services"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type FilePresentation interface {
	UploadFile(c *gin.Context)
	GetFileURLsByMeetingID(c *gin.Context)
	GetFile(c *gin.Context)
	DeleteFile(c *gin.Context)
}

type filePresentation struct {
	fileService services.FileService
}

type UploadFileForm struct {
	MeetingID string                `form:"meeting_id"`
	FormFile  *multipart.FileHeader `form:"file"`
}

type FileResponse struct {
	Url        string `json:"url"`
	UploaderID uint   `json:"uploader_id"`
	FileName   string `json:"file_name"`
	FileID     string `json:"file_id"`
}

type UploadFileResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Url        string `json:"url"`
		UploaderID uint   `json:"uploader_id"`
		FileName   string `json:"file_name"`
		FileID     string `json:"file_id"`
	} `json:"data"`
}

type GetFileURLsByMeetingIDResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    []struct {
		Url        string `json:"url"`
		UploaderID uint   `json:"uploader_id"`
		FileName   string `json:"file_name"`
		FileID     string `json:"file_id"`
	} `json:"data"`
}

type DeleteFileResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func NewFilePresentation(fileServiceArgs ...services.FileService) FilePresentation {
	if len(fileServiceArgs) == 0 {
		return filePresentation{fileService: services.NewFileService()}
	} else if len(fileServiceArgs) == 1 {
		return filePresentation{fileService: fileServiceArgs[0]}
	} else {
		panic("too many arguments")
	}
}

// @Summary Upload a file
// @Description Upload a file
// @Tags File
// @Accept json
// @Produce json
// @Param file form UploadFileForm true "File details"
// @Success 200 {object} UploadFileResponse
// @Router /file [post]
func (fp filePresentation) UploadFile(c *gin.Context) {
	var form UploadFileForm
	var response UploadFileResponse

	if err := c.ShouldBind(&form); err != nil {
		response = UploadFileResponse{
			Status:  "error",
			Message: err.Error(),
		}

		c.JSON(http.StatusBadRequest, response)
		return
	}

	if form.FormFile == nil {
		response = UploadFileResponse{
			Status:  "error",
			Message: "No file uploaded",
		}

		c.JSON(http.StatusBadRequest, response)
		return
	}

	src, err := form.FormFile.Open()
	if err != nil {
		response = UploadFileResponse{
			Status:  "error",
			Message: err.Error(),
		}

		c.JSON(http.StatusBadRequest, response)
		return
	}
	defer src.Close()

	ctx := c.Request.Context()
	operator := c.MustGet("validate_user").(models.User)
	file := &models.File{
		MeetingID:  form.MeetingID,
		UploaderID: operator.ID,
		FileName:   form.FormFile.Filename,
		FileExt:    filepath.Ext(form.FormFile.Filename),
	}

	url, err := fp.fileService.UploadFile(ctx, operator, file, src)
	if err != nil {
		response = UploadFileResponse{
			Status:  "error",
			Message: err.Error(),
		}

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response = UploadFileResponse{
		Status:  "success",
		Message: "File uploaded",
		Data: FileResponse{
			Url:        url,
			UploaderID: operator.ID,
			FileName:   form.FormFile.Filename,
			FileID:     file.ID,
		},
	}

	c.JSON(http.StatusOK, response)
}

// @Summary Get files by meeting ID
// @Description Get files by meeting ID
// @Tags File
// @Param meeting_id path string true "Meeting ID"
// @Success 200 {object} GetFileURLsByMeetingIDResponse
// @Router /file/getFileURLsByMeetingID [get]
func (fp filePresentation) GetFileURLsByMeetingID(c *gin.Context) {
	var response GetFileURLsByMeetingIDResponse

	meetingID := c.Param("meeting_id")
	if len(meetingID) == 0 {
		response = GetFileURLsByMeetingIDResponse{
			Status:  "error",
			Message: "Invalid request",
		}

		c.JSON(http.StatusBadRequest, response)
	}

	ctx := c.Request.Context()
	operator := c.MustGet("validate_user").(models.User)
	files, urls, err := fp.fileService.GetAllSignedURLsByMeeting(ctx, operator, meetingID)
	if err != nil {
		response = GetFileURLsByMeetingIDResponse{
			Status:  "error",
			Message: err.Error(),
		}

		c.JSON(http.StatusBadRequest, response)
	}

	response.Status = "success"
	response.Message = "Files retrieved"
	for i := 0; i < len(files); i++ {
		response.Data = append(response.Data, FileResponse{
			Url:        urls[i],
			UploaderID: files[i].UploaderID,
			FileName:   files[i].FileName,
			FileID:     files[i].ID,
		})
	}

	c.JSON(http.StatusOK, response)
}

// @Summary Get a file
// @Description Get a file
// @Tags File
// @Param id path string true "File ID"
// @Success 200 {object} DeleteFileResponse
// @Router /file/{id} [delete]
func (fp filePresentation) GetFile(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	ctx := c.Request.Context()
	operator := c.MustGet("validate_user").(models.User)
	signedUrl, err := fp.fileService.GetSignedURL(ctx, operator, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Redirect(http.StatusMovedPermanently, signedUrl)
}

// @Summary Delete a file
// @Description Delete a file
// @Tags File
// @Param id path string true "File ID"
// @Success 200 {object} DeleteFileResponse
// @Router /file/{id} [delete]
func (fp filePresentation) DeleteFile(c *gin.Context) {
	var response DeleteFileResponse

	id := c.Param("id")
	if len(id) == 0 {
		response = DeleteFileResponse{
			Status:  "error",
			Message: "Invalid request",
		}

		c.JSON(http.StatusBadRequest, response)
	}

	ctx := c.Request.Context()
	operator := c.MustGet("validate_user").(models.User)
	if err := fp.fileService.DeleteFile(ctx, operator, id); err != nil {
		response = DeleteFileResponse{
			Status:  "error",
			Message: err.Error(),
		}

		c.JSON(http.StatusBadRequest, response)
	}

	response = DeleteFileResponse{
		Status:  "success",
		Message: "File deleted",
	}

	c.JSON(http.StatusOK, response)
}
