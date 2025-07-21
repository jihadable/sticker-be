package config

import (
	"fmt"
	"os"

	"github.com/jihadable/sticker-be/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func DB() *gorm.DB {
	dbHost := os.Getenv("PGHOST")
	dbUser := os.Getenv("PGUSER")
	dbPassword := os.Getenv("PGPASSWORD")
	dbName := os.Getenv("PGDATABASE")
	dbPort := os.Getenv("PGPORT")
	sslMode := os.Getenv("PGSSL")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbHost, dbUser, dbPassword, dbName, dbPort, sslMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
		// SkipDefaultTransaction: true,
		PrepareStmt: true,
	})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	db.AutoMigrate(
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
		&models.NotificationRecipient{},
	)

	return db
}
