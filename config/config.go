package config

import (
	"fmt"
	"fp-super-bootcamp-go/models"
	"fp-super-bootcamp-go/utils"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDataBase() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbProvider := utils.Getenv("DB_PROVIDER", "postgres")
	var db *gorm.DB

	if dbProvider == "postgres" {
		username := os.Getenv("DB_USERNAME")
		password := os.Getenv("DB_PASSWORD")
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		database := os.Getenv("DB_NAME")

		dsn := "host=" + host + " user=" + username + " password=" + password + " dbname=" + database + " port=" + port + " sslmode=require"
		dbGorm, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err.Error())
		}

		db = dbGorm

	} else {
		username := utils.Getenv("DB_USERNAME", "root")
		password := utils.Getenv("DB_PASSWORD", "")
		host := utils.Getenv("DB_HOST", "127.0.0.1")
		port := utils.Getenv("DB_PORT", "3306")
		database := utils.Getenv("DB_NAME", "review_book")

		dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)

		dbGorm, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			panic(err.Error())
		}

		db = dbGorm

	}

	db.AutoMigrate(&models.Users{}, &models.Books{}, &models.Comments{}, &models.Likes{}, &models.Profile{}, &models.Reviews{})

	return db
}
