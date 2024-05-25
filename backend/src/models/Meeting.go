package models

import (
	"gorm.io/datatypes"
	"time"
)

// Meeting represents a meeting scheduled in a meeting room
type Meeting struct {
	ID           int                       `json:"id" gorm:"primaryKey"`
	RoomID       int                       `json:"room_id"`
	OrganizerID  uint                      `json:"organizer"`
	Participants datatypes.JSONSlice[uint] `json:"participants" gorm:"type:jsonb"`
	Title        string                    `json:"title"`
	Description  string                    `json:"description"`
	StartTime    time.Time                 `json:"start_time"`
	EndTime      time.Time                 `json:"end_time"`
	StatusType   string                    `json:"status_type"`
	CreatedAt    time.Time                 `json:"created_at" gorm:"autoCreateTime" swaggerignore:"true"`
	UpdatedAt    time.Time                 `json:"updated_at" gorm:"autoUpdateTime" swaggerignore:"true"`
}
