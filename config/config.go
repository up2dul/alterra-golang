package config

import (
	"alterra2/model"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	errEnv := godotenv.Load()
  if errEnv != nil {
    log.Fatal("Error loading .env file")
  }

	HOST := os.Getenv("HOST")
	PASSWORD := os.Getenv("PASSWORD")
	PORT := os.Getenv("PORT")

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
