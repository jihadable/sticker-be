package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CustomProduct struct {
	Id         string `gorm:"column:id;primaryKey"`
	Name       string `gorm:"column:name"`
	CustomerId string `gorm:"column:customer_id"`
	ImageURL   string `gorm:"column:image_url"`

	Customer *User `gorm:"foreignKey:CustomerId;references:Id"`
}

func (model *CustomProduct) BeforeCreate(tx *gorm.DB) error {
	if model.Id == "" {
		model.Id = uuid.NewString()
	}
	return nil
}
