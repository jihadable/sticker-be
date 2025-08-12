package mapper

import (
	"github.com/jihadable/sticker-be/graph/model"
	"github.com/jihadable/sticker-be/models"
	"github.com/jihadable/sticker-be/services"
)

func DBCategoryToGraphQLCategory(category *models.Category) *model.Category {
	products := make([]*model.Product, len(category.Products))
	for i, product := range category.Products {
		products[i] = DBProductToGraphQLProduct(&product)
	}

	storageService := services.NewStorageService()
	imageURL, _ := storageService.GetPublicFileURL(category.ImageURL)

	return &model.Category{
		ID:       category.Id,
		ImageURL: imageURL,
		Products: products,
	}
}
