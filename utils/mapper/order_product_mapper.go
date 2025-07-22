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
		Order: func() *model.Order {
			if orderProduct.Order != nil {
				return DBOrderToGraphQLOrder(orderProduct.Order)
			} else {
				return nil
			}
		}(),
		Product: func() *model.Product {
			if orderProduct.Product != nil {
				return DBProductToGraphQLProduct(orderProduct.Product)
			} else {
				return nil
			}
		}(),
		CustomProduct: func() *model.CustomProduct {
			if orderProduct.CustomProduct != nil {
				return DBCustomProductToGraphQLCustomProduct(orderProduct.CustomProduct)
			} else {
				return nil
			}
		}(),
	}
}
