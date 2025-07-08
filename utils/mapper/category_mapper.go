package mapper

import (
	"github.com/jihadable/sticker-be/graph/model"
	"github.com/jihadable/sticker-be/models"
)

func DBCategoryToGraphQLCategory(category *models.Category) *model.Category {
	products := make([]*model.Product, len(category.Products))
	for i, product := range category.Products {
		products[i] = DBProductToGraphQLProduct(&product)
	}

	return &model.Category{
		ID:       category.Id,
		Products: products,
	}
}
