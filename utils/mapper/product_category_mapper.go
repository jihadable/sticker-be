package mapper

import (
	"github.com/jihadable/sticker-be/graph/model"
	"github.com/jihadable/sticker-be/models"
)

func DBProductCategoryToGraphQLProductCategory(productCategory *models.ProductCategory) *model.ProductCategory {
	return &model.ProductCategory{
		Product: func() *model.Product {
			if productCategory.Product != nil {
				return DBProductToGraphQLProduct(productCategory.Product)
			} else {
				return nil
			}
		}(),
		Category: func() *model.Category {
			if productCategory.Category != nil {
				return DBCategoryToGraphQLCategory(productCategory.Category)
			} else {
				return nil
			}
		}(),
	}
}
