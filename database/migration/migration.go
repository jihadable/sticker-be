package main

import (
	"fmt"

	"github.com/jihadable/sticker-be/config"
	"github.com/jihadable/sticker-be/models"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func migration(db *gorm.DB) error {
	err := db.Migrator().DropTable(
		&models.User{},
		&models.Product{},
		&models.CustomProduct{},
		&models.Category{},
		&models.ProductCategory{},
		&models.Cart{},
		&models.CartProduct{},
		&models.Order{},
		&models.OrderProduct{},
		&models.Conversation{},
		&models.Message{},
		&models.Notification{},
	)
	if err != nil {
		return err
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.CustomProduct{},
		&models.Category{},
		&models.ProductCategory{},
		&models.Cart{},
		&models.CartProduct{},
		&models.Order{},
		&models.OrderProduct{},
		&models.Conversation{},
		&models.Message{},
		&models.Notification{},
	)

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

	err = migration(db)
	if err != nil {
		panic("Failed to migrate: " + err.Error())
	}

	fmt.Println("Migration successfully ✅")
}
