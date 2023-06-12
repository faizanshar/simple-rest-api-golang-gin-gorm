package models

import (
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:8111)/programmer"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Programmer{})

	DB = database
}
