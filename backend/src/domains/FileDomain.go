package domains

import (
	"meeting-center/src/models"
	"meeting-center/src/repos"
)

type fileDomain struct {
	fileRepo repos.FileRepo
}

type FileDomain interface {
	UploadFile(file *models.File) error
	DeleteFile(id string) error
	GetFile(id string) (models.File, error)
	GetAllFilesByMeetingID(meetingID string) ([]models.File, error)
}

func NewFileDomain(fileRepoArgs ...repos.FileRepo) FileDomain {
	if len(fileRepoArgs) == 0 {
		return fileDomain{fileRepo: repos.NewFileRepo()}
	} else if len(fileRepoArgs) == 1 {
		return fileDomain{fileRepo: fileRepoArgs[0]}
	} else {
		panic("Too many arguments")
	}
}

func (fd fileDomain) UploadFile(file *models.File) error {
	return fd.fileRepo.UploadFile(file)
}

func (fd fileDomain) DeleteFile(id string) error {
	return fd.fileRepo.DeleteFile(id)
}

func (fd fileDomain) GetFile(id string) (models.File, error) {
	return fd.fileRepo.GetFile(id)
}

func (fd fileDomain) GetAllFilesByMeetingID(meetingID string) ([]models.File, error) {
	return fd.fileRepo.GetFilesByMeetingID(meetingID)
}
