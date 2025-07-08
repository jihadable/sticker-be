package mapper

import (
	"github.com/jihadable/sticker-be/graph/model"
	"github.com/jihadable/sticker-be/models"
)

func DBCartToGraphQLCart(cart *models.Cart) *model.Cart {
	products := make([]*model.Product, len(cart.Products))
	for i, product := range cart.Products {
		products[i] = DBProductToGraphQLProduct(&product)
	}

	return &model.Cart{
		ID:       cart.Id,
		Customer: DBUserToGraphQLUser(cart.Customer),
		Products: products,
	}
}
