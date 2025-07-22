package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	Id         string `gorm:"column:id;primaryKey"`
	CustomerId string `gorm:"column:customer_id"`
	Status     string `gorm:"column:status;default:'Pending confirmation'"`
	TotalPrice int    `gorm:"column:total_price"`

	Customer *User     `gorm:"foreignKey:CustomerId;references:Id"`
	Products []Product `gorm:"many2many:order_products;joinForeignKey:OrderId;joinReferences:ProductId"`
}

func (model *Order) BeforeCreate(tx *gorm.DB) error {
	if model.Id == "" {
		model.Id = uuid.NewString()
	}
	return nil
}
