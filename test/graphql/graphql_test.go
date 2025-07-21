package graphql

import "testing"

func TestAll(t *testing.T) {
	t.Run("User test", func(t *testing.T) {
		TestRegisterUserWithValidPayload(t)
		TestPostUserWithInvalidPayload(t)
		TestGetUserWithToken(t)
		TestGetUserWithoutToken(t)
		TestUpdateUser(t)
		TestLoginAsCustomer(t)
		TestLoginAsAdmin(t)
		TestLoginWithInvalidPayload(t)
	})

	t.Run("Product test", func(t *testing.T) {
		TestCreateProductWithValidPayload1(t)
		TestCreateProductWithInvalidPayload(t)
		TestGetProducts(t)
		TestGetProductWithValidId(t)
		TestGetProductWithInvalidId(t)
		TestUpdateProduct(t)
		TestDeleteProduct(t)
		TestCreateProductWithValidPayload2(t)
	})

	t.Run("Custom product test", func(t *testing.T) {
		TestCreateCustomProductWithValidPayload1(t)
		TestCreateCustomProductWithInvalidPayload(t)
		TestGetCustomProductsByUser(t)
		TestGetCustomProductWithValidId(t)
		TestGetCustomProductWithInvalidId(t)
		TestUpdateCustomProduct(t)
		TestDeleteCustomProduct(t)
		TestCreateCustomProductWithValidPayload2(t)
	})

	t.Run("Category test", func(t *testing.T) {
		TestCreateCategoryWithValidPayload1(t)
		TestCreateCategoryWithInvalidPayload(t)
		TestGetCategories(t)
		TestGetCategoryWithValidId(t)
		TestGetCategoryWithInvalidId(t)
		TestDeleteCategory(t)
		TestCreateCategoryWithValidPayload2(t)
	})

	t.Run("Product category test", func(t *testing.T) {
		TestCreateProductCategoryWithValidPayload(t)
		TestCreateProductCategoryWithInvalidPayload(t)
		TestDeleteProductCategory(t)
	})

	t.Run("Cart test", func(t *testing.T) {
		TestGetCartByCustomer(t)
	})

	t.Run("Cart product test", func(t *testing.T) {
		TestCreateCartProductWithProduct(t)
		TestCreateCartProductWithCustomProduct(t)
		TestCreateCartProductWithInvalidPayload(t)
		TestUpdateCartProduct(t)
		TestDeleteCartProduct(t)
	})

	t.Run("Conversation test", func(t *testing.T) {
		TestGetConversationByUser(t)
	})

	t.Run("Message test", func(t *testing.T) {
		TestCreateMessageWithoutReply(t)
		TestCreateMessageWithInvalidPayload(t)
		TestCreateMessageWithProductReply(t)
		TestCreateMessageWithCustomProductReply(t)
	})

	t.Run("Order test", func(t *testing.T) {
		TestPostOrderWithValidPayload(t)
		TestPostOrderWithInvalidPayload(t)
		TestGetOrdersByUser(t)
		TestGetOrderWithValidId(t)
		TestGetOrderWithInvalidId(t)
		TestUpdateOrder(t)
	})
}
