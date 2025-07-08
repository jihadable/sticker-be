package mapper

import (
	"github.com/jihadable/sticker-be/graph/model"
	"github.com/jihadable/sticker-be/models"
)

func DBUserToGraphQLUser(user *models.User) *model.User {
	customProducts := make([]*model.CustomProduct, len(user.CustomProducts))
	for i, customProduct := range user.CustomProducts {
		customProducts[i] = DBCustomProductToGraphQLCustomProduct(&customProduct)
	}

	orders := make([]*model.Order, len(user.Orders))
	for i, order := range user.Orders {
		orders[i] = DBOrderToGraphQLOrder(&order)
	}

	return &model.User{
		ID:             user.Id,
		Name:           user.Name,
		Email:          user.Email,
		Role:           model.Role(user.Role),
		Phone:          user.Phone,
		Address:        user.Address,
		CustomProducts: customProducts,
		Cart:           DBCartToGraphQLCart(&user.Cart),
		Orders:         orders,
		Conversation:   DBConversationTOGraphQLConversation(&user.Conversation),
	}
}
