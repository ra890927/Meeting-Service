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
	// Get the reflect value of the user(user) and other(expected)
	va := reflect.ValueOf(user).Elem()
	vb := reflect.ValueOf(other).Elem()

	for i := 0; i < va.NumField(); i++ {
		fieldA := va.Field(i)
		fieldB := vb.Field(i)

		if !utils.IsEmptyValue(fieldB) && utils.IsEmptyValue(fieldA) {
			// only if that expected's is not empty and the input's field is empty
			// then overwrite the input's field with the expected's field
			fieldA.Set(fieldB)
		}
	}
}
