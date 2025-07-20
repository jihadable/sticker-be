package graphql

import (
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
				size: "L"
			){
				id, quantity, size,
				product { id, name, price, stock, image_url, description }	
			}
		}`,
	})
	request := httptest.NewRequest(fiber.MethodPost, "/graphql", requestBody)

	response, err := App.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)
}

func TestCreateCartProductWithCustomProduct(t *testing.T) {

}

func TestCreateCartProductWithInvalidPayload(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `mutation {
			create_cart_product(
				cart_id: "` + CartId + `"
				product_id: "` + ProductId + `"
				quantity: 3
				size: "L"
			){
				id, quantity, size,
				product { id, name, price, stock, image_url, description }	
			}
		}`,
	})
	request := httptest.NewRequest(fiber.MethodPost, "/graphql", requestBody)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+CustomerJWT)

	response, err := App.Test(request)

	assert.Nil(t, err)
	assert.NotEqual(t, fiber.StatusOK, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)

	assert.Nil(t, responseBody["data"])

	errors, ok := responseBody["errors"].([]any)
	assert.True(t, ok)
	assert.NotEmpty(t, errors)

	t.Log("✅")
}

func TestUpdateCartProduct(t *testing.T) {

}
