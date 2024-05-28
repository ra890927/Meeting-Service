package models

import (
	"gorm.io/datatypes"
	"time"
)

// Room defines the structure for a meeting room
type Room struct {
	ID        int                      `json:"id" gorm:"primaryKey"`
	RoomName  string                   `json:"room_name"`
	Type      string                   `json:"type"`
	Rules     datatypes.JSONSlice[int] `json:"rules" gorm:"type:jsonb"`
	Capacity  int                      `json:"capacity"`
	CreatedAt time.Time                `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time                `json:"updated_at" gorm:"autoUpdateTime"`
}
