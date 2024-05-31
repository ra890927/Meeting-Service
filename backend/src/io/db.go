package io

import (
	"sync"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"meeting-center/src/models"
)

var (
	dbInstance *gorm.DB
	once       sync.Once
)

func init() {
	db := GetDBInstance()

	models := []interface{}{
		&models.User{},
		&models.Room{},
		&models.Meeting{},
		&models.CodeType{},
		&models.CodeValue{},
	}

	for _, model := range models {
		if err := db.AutoMigrate(model); err != nil {
			panic(err)
		}
	}
}

func GetDBInstance() *gorm.DB {
	if dbInstance == nil {
		once.Do(func() {
			dsn := "./sqlite.db"
			db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
			if err != nil {
				panic("Connect db error")
			}

			sqlDB, err := db.DB()
			if err != nil {
				panic("Get DB instance error")
			}

			sqlDB.SetMaxIdleConns(10)
			sqlDB.SetMaxOpenConns(100)
			sqlDB.SetConnMaxLifetime(time.Hour)

			dbInstance = db
		})
	}
	return dbInstance
}
