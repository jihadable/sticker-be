package mapper

import (
	"github.com/jihadable/sticker-be/graph/model"
	"github.com/jihadable/sticker-be/models"
)

func DBOrderToGraphQLOrder(order *models.Order) *model.Order {
	products := make([]*model.Product, len(order.Products))
	for i, product := range order.Products {
		products[i] = DBProductToGraphQLProduct(&product)
	}

	return &model.Order{
		ID:         order.Id,
		Status:     order.Status,
		TotalPrice: int32(order.TotalPrice),
		Customer:   DBUserToGraphQLUser(order.Customer),
		Products:   products,
	}
}
