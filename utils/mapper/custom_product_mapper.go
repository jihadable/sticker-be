package mapper

import (
	"github.com/jihadable/sticker-be/graph/model"
	"github.com/jihadable/sticker-be/models"
	"github.com/jihadable/sticker-be/services"
)

func DBCustomProductToGraphQLCustomProduct(customProduct *models.CustomProduct) *model.CustomProduct {
	storageService := services.NewStorageService()
	imageURL, _ := storageService.GetPublicFileURL(customProduct.ImageURL)

	return &model.CustomProduct{
		ID:       customProduct.Id,
		Name:     customProduct.Name,
		ImageURL: imageURL,
		Customer: func() *model.User {
			if customProduct.Customer != nil {
				return DBUserToGraphQLUser(customProduct.Customer)
			} else {
				return nil
			}
		}(),
	}
}
