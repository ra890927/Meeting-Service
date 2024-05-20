package models

import (
	"time"
)

// Room defines the structure for a meeting room
type Room struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	RoomName  string    `json:"room_name"`
	Type      string    `json:"type"`
	Rules     []string  `json:"rules"`
	Capacity  int       `json:"capacity"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
