package models

type Category struct {
	Id string `gorm:"column:id;primaryKey"`

	Products []Product `gorm:"many2many:product_categories;joinForeignKey:CategoryId;joinReferences:ProductId"`
}
