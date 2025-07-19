package graphql

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestPostOrderWithValidPayload(t *testing.T) {

}

func TestPostOrderWithInvalidPayload(t *testing.T) {

}

func TestGetOrdersByUser(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `query {
			orders_by_user {
				id, status, total_price, 
				customer { id, name, email, role, phone, address }, 
				products { id, name, price, stock, image_url, description }
			}
		}`,
	})
	request := httptest.NewRequest(fiber.MethodPost, "/graphql", requestBody)

	response, err := App.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)

	data, ok := responseBody["data"].(map[string]any)
	assert.True(t, ok)

	orders, ok := data["orders_by_user"].([]map[string]any)
	assert.True(t, ok)

	order := orders[0]

	assert.NotEmpty(t, order["id"])
	assert.NotEmpty(t, order["status"])
	assert.NotEmpty(t, order["total_price"])

	customer, ok := order["customer"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, customer["id"])
	assert.NotEmpty(t, customer["name"])
	assert.NotEmpty(t, customer["email"])
	assert.NotEmpty(t, customer["role"])
	assert.NotEmpty(t, customer["phone"])
	assert.NotEmpty(t, customer["address"])

	products, ok := order["products"].([]map[string]any)
	assert.True(t, ok)

	product := products[0]
	assert.NotEmpty(t, product["id"])
	assert.NotEmpty(t, product["name"])
	assert.NotEmpty(t, product["price"])
	assert.NotEmpty(t, product["stock"])
	assert.NotEmpty(t, product["image_url"])
	assert.NotEmpty(t, product["description"])

	t.Log("✅")
}

func TestGetOrderById(t *testing.T) {

}

func TestUpdateOrder(t *testing.T) {

}
