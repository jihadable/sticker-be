package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderProduct struct {
	Id              string  `gorm:"column:id;primaryKey"`
	OrderId         string  `gorm:"column:order_id"`
	ProductId       *string `gorm:"column:product_id"`
	CustomProductId *string `gorm:"column:custom_product_id"`
	Quantity        int     `gorm:"column:quantity"`
	Size            string  `gorm:"column:size"`
	SubtotalPrice   int     `gorm:"column:subtotal_price"`

	Order         *Order         `gorm:"foreignKey:OrderId;references:Id"`
	Product       *Product       `gorm:"foreignKey:ProductId;references:Id"`
	CustomProduct *CustomProduct `gorm:"foreignKey:CustomProductId;references:Id"`
}

func (model *OrderProduct) BeforeCreate(tx *gorm.DB) error {
	if model.Id == "" {
		model.Id = uuid.NewString()
	}
	return nil
}
