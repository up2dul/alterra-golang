package config

import (
	"alterra2/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	HOST := "containers-us-west-11.railway.app"
	PASSWORD := "oMsPp1mwDRFPl20lhLU5"
	PORT := "6393"

	dsn := "root:" + PASSWORD + "@tcp(" + HOST + ":" + PORT +")/railway?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("connection DB")
	}
	
	migration()
}

func migration()  {
	DB.AutoMigrate(&model.Book{})
}
