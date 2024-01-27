package utils

import (
	"log"
	"user_api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func ConnectDB() {

	dsn := "root@tcp(localhost:3306)/demo"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Can Not Connect Database", err.Error())
	}
	DB.AutoMigrate(&models.User{})

}

func GetDB() *gorm.DB {
	return DB
}
