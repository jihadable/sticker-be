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
	requestBody := RequestBodyParser(map[string]string{
		"query": `mutation {
			post_order(){
				id, status, total_price, 
				customer { id, name, email, role, phone, address }, 
				products { id, name, price, stock, image_url, description }
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
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+CustomerJWT)

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

func TestGetOrderWithValidId(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `query {
			order(id: "` + OrderId + `"){
				id, status, total_price, 
				customer { id, name, email, role, phone, address }, 
				products { id, name, price, stock, image_url, description }
			}
		}`,
	})
	request := httptest.NewRequest(fiber.MethodPost, "/graphql", requestBody)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+CustomerJWT)

	response, err := App.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)

	data, ok := responseBody["data"].(map[string]any)
	assert.True(t, ok)

	order, ok := data["order"].(map[string]any)
	assert.True(t, ok)

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

func TestGetOrderWithInvalidId(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `query {
			order(id: "xxx"){
				id, status, total_price, 
				customer { id, name, email, role, phone, address }, 
				products { id, name, price, stock, image_url, description }
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

func TestUpdateOrder(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `mutation {
			update_order(id: "` + OrderId + `", status: "Completed"){
				id, status, total_price, 
				customer { id, name, email, role, phone, address }, 
				products { id, name, price, stock, image_url, description }	
			}
		}`,
	})
	request := httptest.NewRequest(fiber.MethodPost, "/graphql", requestBody)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+AdminJWT)

	response, err := App.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)

	data, ok := responseBody["data"].(map[string]any)
	assert.True(t, ok)

	order, ok := data["update_order"].(map[string]any)
	assert.True(t, ok)

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
