package config

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dbUsername = "DB_USERNAME"
	dbPassword = "DB_PASSWORD"
	dbHost     = "DB_HOST"
	dbName     = "DB_HOST"
	dbPort     = "DB_PORT"
)

func GetConfig() string {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parsTime=true", os.Getenv(dbName), os.Getenv(dbPassword),
		os.Getenv(dbHost), os.Getenv(dbPort), os.Getenv(dbName))
}

func InitDatabaseConnection(conf string) *gorm.DB {
	DB, err := gorm.Open(mysql.Open(GetConfig()), &gorm.Config{})
	if err != nil {
		panic("Connection Failed!")
	}
	return DB
}
