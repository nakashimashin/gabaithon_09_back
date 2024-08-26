package database

import (
	"fmt"
	"gabaithon-09-back/models"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	DB = Connect()
}

func Connect() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(gabaithon-09-db:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func Migrate() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("failed to migrate: %v", err)
	} else {
		log.Println("migrate success")
	}
}
