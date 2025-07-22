package graphql

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetCartByCustomer(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `query {
			get_cart_by_customer {
				id,
				customer { id, name, email, role, phone, address },
				products { id, name, price, stock, image_url, description }
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

	cart, ok := data["get_cart_by_customer"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, cart["id"])
	CartId = cart["id"].(string)

	customer, ok := cart["customer"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, customer["id"])
	assert.NotEmpty(t, customer["name"])
	assert.NotEmpty(t, customer["email"])
	assert.NotEmpty(t, customer["role"])
	assert.NotEmpty(t, customer["phone"])
	assert.NotEmpty(t, customer["address"])

	assert.Empty(t, cart["products"])

	t.Log("✅")
}
