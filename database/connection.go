package database

import (
	"latihan/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root@(localhost)/latihan?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}
