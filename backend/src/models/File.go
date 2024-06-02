package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type File struct {
	ID         string    `json:"id" gorm:"primaryKey"`
	MeetingID  string    `json:"meeting_id"`
	UploaderID uint      `json:"uploader_id"`
	FileName   string    `json:"file_name"`
	FileExt    string    `json:"file_ext"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime" swaggerignore:"true"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime" swaggerignore:"true"`
}

func (file *File) BeforeCreate(db *gorm.DB) (err error) {
	file.ID = uuid.New().String()
	return nil
}
