package config

import (
	"fmt"
	"game_suit/core/entity"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dbUsername = "DB_USERNAME"
	dbPassword = "DB_PASSWORD"
	dbHost     = "DB_HOST"
	dbName     = "DB_NAME"
	dbPort     = "DB_PORT"
)

func GetConfig() string {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true", os.Getenv(dbUsername), os.Getenv(dbPassword),
		os.Getenv(dbHost), os.Getenv(dbPort), os.Getenv(dbName))
}

func InitDatabaseConnection(conf string) *gorm.DB {
	db := GetConfig()
	fmt.Println("ini :", db)
	DB, err := gorm.Open(mysql.Open(GetConfig()), &gorm.Config{})
	if err != nil {
		panic("Connection Failed!")
	}
	return DB
}

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		entity.User{},
	)
	if err != nil {
		log.Fatalf("Migration Fail. Error : %v", err)
	}
	log.Println("Migration Succes")
}
