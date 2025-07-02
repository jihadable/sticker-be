package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CustomProduct struct {
	Id         uuid.UUID `gorm:"column:id;primaryKey"`
	CustomerId string    `gorm:"column:customer_id"`
	ImageURL   string    `gorm:"column:image_url"`

	Customer *User `gorm:"foreignKey:CustomerId;references:Id"`
}

func (model *CustomProduct) BeforeCreate(tx *gorm.DB) error {
	if model.Id == uuid.Nil {
		model.Id = uuid.New()
	}
	return nil
}
