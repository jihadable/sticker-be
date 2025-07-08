package graph

import "github.com/jihadable/sticker-be/services"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	services.UserService
	services.ProductService
	services.CustomProductService
	services.CategoryService
	services.ProductCategoryService
	services.CartService
	services.CartProductService
	services.OrderService
	services.OrderProductService
	services.ConversationService
	services.MessageService
}
