package clients

import (
	"fmt"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"meeting-center/src/models"

	"github.com/spf13/viper"
)

var (
	dbInstance *gorm.DB
	dbOnce     sync.Once
)

func InitDB() {
	db := GetDBInstance()

	models := []interface{}{
		&models.User{},
		&models.Room{},
		&models.File{},
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

func getDSNFromConfig() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("mysql.username"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetString("mysql.port"),
		viper.GetString("mysql.database"),
	)
}

func GetDBInstance() *gorm.DB {
	if dbInstance == nil {
		dbOnce.Do(func() {
			db, err := gorm.Open(mysql.Open(getDSNFromConfig()), &gorm.Config{})
			if err != nil {
				panic("Connect db error" + err.Error())
			}

			sqlDB, err := db.DB()
			if err != nil {
				panic("Get DB instance error")
			}

			sqlDB.SetMaxIdleConns(viper.GetInt("mysql.maxIdleConns"))
			sqlDB.SetMaxOpenConns(viper.GetInt("mysql.maxOpenConns"))
			sqlDB.SetConnMaxLifetime(time.Duration(viper.GetInt("mysql.connMaxLifetime")) * time.Second)

			dbInstance = db
		})
	}
	return dbInstance
}
