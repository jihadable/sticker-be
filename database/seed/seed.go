package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jihadable/sticker-be/config"
	"github.com/jihadable/sticker-be/models"
	"github.com/jihadable/sticker-be/utils"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func truncateAllTable(db *gorm.DB) error {
	tables := []string{
		"users",
		"products",
		"custom_products",
		"categories",
		"product_categories",
		"carts",
		"cart_products",
		"conversations",
		"messages",
		"orders",
		"order_products",
		"notifications",
	}

	query := fmt.Sprintf("TRUNCATE TABLE %s CASCADE", strings.Join(tables, ", "))

	if err := db.Exec(query).Error; err != nil {
		return fmt.Errorf("failed to truncate tables: %w", err)
	}

	return nil
}

func userSeeder(db *gorm.DB) error {
	hashedPassword, err := utils.HashPassword(os.Getenv("PRIVATE_PASSWORD"))
	if err != nil {
		return err
	}

	admin := models.User{
		Name:     "Stiker Admin",
		Email:    "stikeradmin@gmail.com",
		Password: hashedPassword,
		Role:     "admin",
	}

	err = db.Create(&admin).Error
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Failed to read .env file: " + err.Error())
	}

	db := config.DB()

	err = truncateAllTable(db)
	if err != nil {
		panic("Failed to truncate all table: " + err.Error())
	}

	err = userSeeder(db)
	if err != nil {
		panic("Failed to seed users: " + err.Error())
	}

	fmt.Println("Seed successfully ✅")
}
