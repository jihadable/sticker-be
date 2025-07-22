package mapper

import (
	"github.com/jihadable/sticker-be/graph/model"
	"github.com/jihadable/sticker-be/models"
)

func DBCustomProductToGraphQLCustomProduct(customProduct *models.CustomProduct) *model.CustomProduct {
	return &model.CustomProduct{
		ID:       customProduct.Id,
		Name:     customProduct.Name,
		ImageURL: customProduct.ImageURL,
		Customer: func() *model.User {
			if customProduct.Customer != nil {
				return DBUserToGraphQLUser(customProduct.Customer)
			} else {
				return nil
			}
		}(),
	}
}
