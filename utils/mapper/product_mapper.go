package mapper

import (
	"github.com/jihadable/sticker-be/graph/model"
	"github.com/jihadable/sticker-be/models"
	"github.com/jihadable/sticker-be/services"
)

func DBProductToGraphQLProduct(product *models.Product) *model.Product {
	categories := make([]*model.Category, len(product.Categories))
	for i, category := range product.Categories {
		categories[i] = DBCategoryToGraphQLCategory(&category)
	}

	storageService := services.NewStorageService()
	imageURL, _ := storageService.GetPublicFileURL(product.ImageURL)

	return &model.Product{
		ID:          product.Id,
		Name:        product.Name,
		Price:       int32(product.Price),
		Stock:       int32(product.Stock),
		ImageURL:    imageURL,
		Description: product.Description,
		Categories:  categories,
	}
}
