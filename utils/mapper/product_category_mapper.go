package mapper

import (
	"github.com/jihadable/sticker-be/graph/model"
	"github.com/jihadable/sticker-be/models"
)

func DBProductCategoryToGraphQLProductCategory(productCategory *models.ProductCategory) *model.ProductCategory {
	return &model.ProductCategory{
		Product:  DBProductToGraphQLProduct(productCategory.Product),
		Category: DBCategoryToGraphQLCategory(productCategory.Category),
	}
}
