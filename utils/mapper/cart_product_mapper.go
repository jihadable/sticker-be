package mapper

import (
	"github.com/jihadable/sticker-be/graph/model"
	"github.com/jihadable/sticker-be/models"
)

func DBCartProductToGraphQLCartProduct(cartProduct *models.CartProduct) *model.CartProduct {
	return &model.CartProduct{
		ID:       cartProduct.Id,
		Quantity: int32(cartProduct.Quantity),
		Size:     model.Size(cartProduct.Size),
		Cart: func() *model.Cart {
			if cartProduct.Cart != nil {
				return DBCartToGraphQLCart(cartProduct.Cart)
			} else {
				return nil
			}
		}(),
		Product: func() *model.Product {
			if cartProduct.Product != nil {
				return DBProductToGraphQLProduct(cartProduct.Product)
			} else {
				return nil
			}
		}(),
		CustomProduct: func() *model.CustomProduct {
			if cartProduct.CustomProduct != nil {
				return DBCustomProductToGraphQLCustomProduct(cartProduct.CustomProduct)
			} else {
				return nil
			}
		}(),
	}
}
