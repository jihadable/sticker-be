package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CartProduct struct {
	Id              string  `gorm:"column:id;primaryKey"`
	CartId          string  `gorm:"column:cart_id"`
	ProductId       *string `gorm:"column:product_id"`
	CustomProductId *string `gorm:"column:custom_product_id"`
	Quantity        int     `gorm:"column:quantity"`
	Size            string  `gorm:"column:size"`

	Cart          *Cart          `gorm:"foreignKey:CartId;references:Id"`
	Product       *Product       `gorm:"foreignKey:ProductId;references:Id"`
	CustomProduct *CustomProduct `gorm:"foreignKey:CustomProductId;references:Id"`
}

func (model *CartProduct) BeforeCreate(tx *gorm.DB) error {
	if model.Id == "" {
		model.Id = uuid.NewString()
	}
	return nil
}
