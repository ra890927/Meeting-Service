package db

import (
	"sync"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	dbInstance *gorm.DB
	once       sync.Once
)

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
