package mapper

import (
	"github.com/jihadable/sticker-be/graph/model"
	"github.com/jihadable/sticker-be/models"
)

func DBCartProductToGraphQLCartProduct(cartProduct *models.CartProduct) *model.CartProduct {
	return &model.CartProduct{
		ID:            cartProduct.Id,
		Quantity:      int32(cartProduct.Quantity),
		Size:          model.Size(cartProduct.Size),
		Cart:          DBCartToGraphQLCart(cartProduct.Cart),
		Product:       DBProductToGraphQLProduct(cartProduct.Product),
		CustomProduct: DBCustomProductToGraphQLCustomProduct(cartProduct.CustomProduct),
	}
}
