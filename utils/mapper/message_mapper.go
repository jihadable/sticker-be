package mapper

import (
	"github.com/jihadable/sticker-be/graph/model"
	"github.com/jihadable/sticker-be/models"
)

func DBMessageToGraphQLMessage(message *models.Message) *model.Message {
	return &model.Message{
		ID:            message.Id,
		Message:       message.Message,
		Product:       DBProductToGraphQLProduct(message.Product),
		CustomProduct: DBCustomProductToGraphQLCustomProduct(message.CustomProduct),
		Sender:        DBUserToGraphQLUser(message.Sender),
	}
}
