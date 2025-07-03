package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	Id          string `gorm:"column:id;primaryKey"`
	Name        string `grom:"column:name"`
	Price       int    `gorm:"column:price"`
	Stock       int    `gorm:"column:stock"`
	ImageURL    string `gorm:"column:image_url"`
	Description string `gorm:"column:description"`

	Categories []Category `gorm:"many2many:product_categories;joinForeignKey:ProductId;joinReferences:CategoryId"`
}

func (model *Product) BeforeCreate(tx *gorm.DB) error {
	if model.Id == "" {
		model.Id = uuid.NewString()
	}
	return nil
}
