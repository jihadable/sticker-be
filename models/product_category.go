package models

type ProductCategory struct {
	ProductId  string `gorm:"column:product_id;primaryKey"`
	CategoryId string `gorm:"column:category_id;primaryKey"`

	Product  *Product  `gorm:"foreignKey:ProductId;references:Id"`
	Category *Category `gorm:"foreignKey:CategoryId;references:Id"`
}
