package models

import (
	"meeting-center/src/utils"
	"reflect"
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (user *User) OverwriteValue(other *User) {
	va := reflect.ValueOf(user).Elem()
	vb := reflect.ValueOf(other).Elem()

	for i := 0; i < va.NumField(); i++ {
		fieldA := va.Field(i)
		fieldB := vb.Field(i)

		if !utils.IsEmptyValue(fieldB) {
			fieldA.Set(fieldB)
		}
	}
}
