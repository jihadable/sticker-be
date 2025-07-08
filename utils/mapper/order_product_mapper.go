package mapper

import (
	"github.com/jihadable/sticker-be/graph/model"
	"github.com/jihadable/sticker-be/models"
)

func DBOrderProductToGraphQLOrderProduct(orderProduct *models.OrderProduct) *model.OrderProduct {
	return &model.OrderProduct{
		ID:            orderProduct.Id,
		Quantity:      int32(orderProduct.Quantity),
		Size:          model.Size(orderProduct.Size),
		SubtotalPrice: int32(orderProduct.SubtotalPrice),
		Order:         DBOrderToGraphQLOrder(orderProduct.Order),
		Product:       DBProductToGraphQLProduct(orderProduct.Product),
		CustomProduct: DBCustomProductToGraphQLCustomProduct(orderProduct.CustomProduct),
	}
}
