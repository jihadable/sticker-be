package main

import (
	"fmt"
	"os"

	"github.com/jihadable/sticker-be/config"
	"github.com/jihadable/sticker-be/models"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func userSeeder(db *gorm.DB) error {
	admin := models.User{
		Name:     "Stiker Admin",
		Email:    "stikeradmin@gmail.com",
		Password: os.Getenv("PRIVATE_PASSWORD"),
		Role:     "admin",
	}

	err := db.Create(&admin).Error
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := godotenv.Load(".env.local")
	if err != nil {
		panic("Failed to read .env file: " + err.Error())
	}

	db := config.DB()

	err = userSeeder(db)
	if err != nil {
		panic("Failed to seed users: " + err.Error())
	}

	fmt.Println("Seed successfully ✅")
}
