package models

import (
	"time"
)

// Room defines the structure for a meeting room
type Room struct {
	ID        int       `json:"id" gorm:"primaryKey" swaggerignore:"true"`
	RoomName  string    `json:"room_name"`
	Type      string    `json:"type"`
	Rules     []string  `json:"rules"`
	Capacity  int       `json:"capacity"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime" swaggerignore:"true"`
}
