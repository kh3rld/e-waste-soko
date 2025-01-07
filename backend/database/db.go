package database

import (
	"log"

	"github.com/kh3rld/e-waste-soko/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=postgres dbname=soko_portal port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	DB = database

	// AutoMigrate Models
	err = DB.AutoMigrate(&models.User{}, &models.Waste{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	log.Println("Database connection and migration successful!")
}
