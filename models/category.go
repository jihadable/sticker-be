package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	Id   uuid.UUID `gorm:"column:id;primaryKey"`
	Name string    `gorm:"column:name"`

	Products []Product `gorm:"many2many:product_categories;joinForeignKey:CategoryId;joinReferences:ProductId"`
}

func (model *Category) BeforeCreate(tx *gorm.DB) error {
	if model.Id == uuid.Nil {
		model.Id = uuid.New()
	}
	return nil
}
