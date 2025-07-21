package mapper

import (
	"github.com/jihadable/sticker-be/graph/model"
	"github.com/jihadable/sticker-be/models"
)

func DBMessageToGraphQLMessage(message *models.Message) *model.Message {
	return &model.Message{
		ID:      message.Id,
		Message: message.Message,
		Product: func() *model.Product {
			if message.Product != nil {
				return DBProductToGraphQLProduct(message.Product)
			} else {
				return nil
			}
		}(),
		CustomProduct: func() *model.CustomProduct {
			if message.CustomProduct != nil {
				return DBCustomProductToGraphQLCustomProduct(message.CustomProduct)
			} else {
				return nil
			}
		}(),
		Sender: DBUserToGraphQLUser(message.Sender),
	}
}
