package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Cart struct {
	Id         uuid.UUID `gorm:"column:id;primaryKey"`
	CustomerId string    `gorm:"column:customer_id"`

	Customer *User     `gorm:"foreignKey:CustomerId;references:Id"`
	Products []Product `gorm:"many2many:cart_products;joinForeignKey:CartId;joinReferences:ProductId"`
}

func (model *Cart) BeforeCreate(tx *gorm.DB) error {
	if model.Id == uuid.Nil {
		model.Id = uuid.New()
	}
	return nil
}
