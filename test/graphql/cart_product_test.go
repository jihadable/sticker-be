package graphql

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestCreateCartProductWithProduct(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `mutation {
			create_cart_product(
				cart_id: "` + CartId + `"
				product_id: "` + ProductId + `"
				quantity: 3
				size: L
			){
				id, quantity, size,
				product { id, name, price, stock, image_url, description }	
			}
		}`,
	})
	request := httptest.NewRequest(fiber.MethodPost, "/graphql", requestBody)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+CustomerJWT)

	response, err := App.Test(request, -1)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)

	data, ok := responseBody["data"].(map[string]any)
	assert.True(t, ok)

	cartProduct, ok := data["create_cart_product"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, cartProduct["id"])
	CartProductId = cartProduct["id"].(string)
	assert.Equal(t, float64(3), cartProduct["quantity"])
	assert.Equal(t, "L", cartProduct["size"])

	product, ok := cartProduct["product"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, product["id"])
	assert.NotEmpty(t, product["name"])
	assert.NotEmpty(t, product["price"])
	assert.NotEmpty(t, product["stock"])
	assert.NotEmpty(t, product["image_url"])
	assert.NotEmpty(t, product["description"])

	fmt.Println("TestCreateCartProductWithProduct: ✅")
}

func TestCreateCartProductWithCustomProduct(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `mutation {
			create_cart_product(
				cart_id: "` + CartId + `"
				custom_product_id: "` + CustomProductId + `"
				quantity: 2
				size: XL
			){
				id, quantity, size,
				custom_product { id, name, image_url }	
			}
		}`,
	})
	request := httptest.NewRequest(fiber.MethodPost, "/graphql", requestBody)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+CustomerJWT)

	response, err := App.Test(request, -1)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)

	data, ok := responseBody["data"].(map[string]any)
	assert.True(t, ok)

	cartProduct, ok := data["create_cart_product"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, cartProduct["id"])
	assert.Equal(t, float64(2), cartProduct["quantity"])
	assert.Equal(t, "XL", cartProduct["size"])

	customProduct, ok := cartProduct["custom_product"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, customProduct["id"])
	assert.NotEmpty(t, customProduct["name"])
	assert.NotEmpty(t, customProduct["image_url"])

	fmt.Println("TestCreateCartProductWithCustomProduct: ✅")
}

func TestCreateCartProductWithInvalidPayload(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `mutation {
			create_cart_product(){
				id, quantity, size,
				product { id, name, price, stock, image_url, description }	
			}
		}`,
	})
	request := httptest.NewRequest(fiber.MethodPost, "/graphql", requestBody)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+CustomerJWT)

	response, err := App.Test(request, -1)

	assert.Nil(t, err)

	responseBody := ResponseBodyParser(response.Body)

	assert.Nil(t, responseBody["data"])

	errors, ok := responseBody["errors"].([]any)
	assert.True(t, ok)
	assert.NotEmpty(t, errors)

	fmt.Println("TestCreateCartProductWithInvalidPayload: ✅")
}

func TestUpdateCartProduct(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `mutation {
			update_cart_product(
				id: "` + CartProductId + `"
				quantity: 2
				size: M
			){
				id, quantity, size,
				product { id, name, price, stock, image_url, description }	
			}
		}`,
	})
	request := httptest.NewRequest(fiber.MethodPost, "/graphql", requestBody)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+CustomerJWT)

	response, err := App.Test(request, -1)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)

	data, ok := responseBody["data"].(map[string]any)
	assert.True(t, ok)

	cartProduct, ok := data["update_cart_product"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, cartProduct["id"])
	assert.Equal(t, float64(2), cartProduct["quantity"])
	assert.Equal(t, "M", cartProduct["size"])

	product, ok := cartProduct["product"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, product["id"])
	assert.NotEmpty(t, product["name"])
	assert.NotEmpty(t, product["price"])
	assert.NotEmpty(t, product["stock"])
	assert.NotEmpty(t, product["image_url"])
	assert.NotEmpty(t, product["description"])

	fmt.Println("TestUpdateCartProduct: ✅")
}

func TestDeleteCartProduct(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `mutation {
			delete_cart_product(id: "` + CartProductId + `")	
		}`,
	})
	request := httptest.NewRequest(fiber.MethodPost, "/graphql", requestBody)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+CustomerJWT)

	response, err := App.Test(request, -1)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)

	data, ok := responseBody["data"].(map[string]any)
	assert.True(t, ok)

	deleteCartProduct, ok := data["delete_cart_product"].(bool)
	assert.True(t, ok)
	assert.True(t, deleteCartProduct)

	fmt.Println("TestDeleteCartProduct: ✅")
}
