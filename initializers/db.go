package initializers

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"main.go/models"
)

var DB *gorm.DB

func LoadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func DBconnection() {
	var err error
	dsn := os.Getenv("DB")
	if dsn == "" {
		log.Fatal("Error: Database DSN is missing in .env file")
		return
	}

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
		return
	}

	fmt.Println("Connected to the database successfully")
	
	if err := DB.AutoMigrate(&models.UserModel{}, &models.AdminModel{}); err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}
}
