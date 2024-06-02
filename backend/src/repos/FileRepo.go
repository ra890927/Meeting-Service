package repos

import (
	"meeting-center/src/clients"
	"meeting-center/src/models"

	"gorm.io/gorm"
)

type fileRepo struct {
	db *gorm.DB
}

type FileRepo interface {
	UploadFile(file *models.File) error
	DeleteFile(id string) error
	GetFile(id string) (models.File, error)
	GetFilesByMeetingID(meetingID string) ([]models.File, error)
}

func NewFileRepo(dbArgs ...*gorm.DB) FileRepo {
	if len(dbArgs) == 0 {
		return fileRepo{db: clients.GetDBInstance()}
	} else if len(dbArgs) == 1 {
		return fileRepo{db: dbArgs[0]}
	} else {
		panic("Too many arguments")
	}
}

func (fr fileRepo) UploadFile(file *models.File) error {
	result := fr.db.Create(file)
	return result.Error
}

func (fr fileRepo) DeleteFile(id string) error {
	result := fr.db.Where("id = ?", id).Delete(&models.File{})
	return result.Error
}

func (fr fileRepo) GetFile(id string) (models.File, error) {
	var file models.File
	result := fr.db.Where("id = ?", id).First(&file)
	if result.Error != nil {
		return models.File{}, result.Error
	}
	return file, nil
}

func (fr fileRepo) GetFilesByMeetingID(meetingID string) ([]models.File, error) {
	var files []models.File
	result := fr.db.Where("meeting_id = ?", meetingID).Find(&files)
	if result.Error != nil {
		return []models.File{}, result.Error
	}
	return files, nil
}
