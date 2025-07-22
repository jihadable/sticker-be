package graphql

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrderWithValidPayload(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `mutation {
			create_order(
				order_items: [
					{
						product_id: "` + ProductId + `"
						quantity: 2
						size: M
						subtotal_price: 5000
					},
					{
						custom_product_id: "` + CustomProductId + `"
						quantity: 3
						size: L
						subtotal_price: 7500
					}
				],
				total_price: 12500
			){
				id, status, total_price, 
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

	order, ok := data["create_order"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, order["id"])
	OrderId = order["id"].(string)
	assert.Equal(t, "Pending confirmation", order["status"])
	assert.Equal(t, float64(12500), order["total_price"])

	customer, ok := order["customer"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, customer["id"])
	assert.NotEmpty(t, customer["name"])
	assert.NotEmpty(t, customer["email"])
	assert.NotEmpty(t, customer["role"])
	assert.NotEmpty(t, customer["phone"])
	assert.NotEmpty(t, customer["address"])

	products, ok := order["products"].([]any)
	assert.True(t, ok)

	product := products[0].(map[string]any)
	assert.NotEmpty(t, product["id"])
	assert.NotEmpty(t, product["name"])
	assert.NotEmpty(t, product["price"])
	assert.NotEmpty(t, product["stock"])
	assert.NotEmpty(t, product["image_url"])
	assert.NotEmpty(t, product["description"])

	fmt.Println("TestCreateOrderWithValidPayload ✅")
}

func TestCreateOrderWithInvalidPayload(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `mutation {
			create_order(){
				id, status, total_price, 
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

	responseBody := ResponseBodyParser(response.Body)

	assert.Nil(t, responseBody["data"])

	errors, ok := responseBody["errors"].([]any)
	assert.True(t, ok)
	assert.NotEmpty(t, errors)

	fmt.Println("TestCreateOrderWithInvalidPayload: ✅")
}

func TestGetOrdersByCustomer(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `query {
			get_orders_by_customer {
				id, status, total_price, 
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

	orders, ok := data["get_orders_by_customer"].([]any)
	assert.True(t, ok)

	order := orders[0].(map[string]any)

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

	products, ok := order["products"].([]any)
	assert.True(t, ok)

	product := products[0].(map[string]any)
	assert.NotEmpty(t, product["id"])
	assert.NotEmpty(t, product["name"])
	assert.NotEmpty(t, product["price"])
	assert.NotEmpty(t, product["stock"])
	assert.NotEmpty(t, product["image_url"])
	assert.NotEmpty(t, product["description"])

	fmt.Println("TestGetOrdersByCustomer: ✅")
}

func TestGetOrderWithValidId(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `query {
			get_order(id: "` + OrderId + `"){
				id, status, total_price, 
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

	order, ok := data["get_order"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, order["id"])
	assert.Equal(t, "Pending confirmation", order["status"])
	assert.Equal(t, float64(12500), order["total_price"])

	customer, ok := order["customer"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, customer["id"])
	assert.NotEmpty(t, customer["name"])
	assert.NotEmpty(t, customer["email"])
	assert.NotEmpty(t, customer["role"])
	assert.NotEmpty(t, customer["phone"])
	assert.NotEmpty(t, customer["address"])

	products, ok := order["products"].([]any)
	assert.True(t, ok)

	product := products[0].(map[string]any)
	assert.NotEmpty(t, product["id"])
	assert.NotEmpty(t, product["name"])
	assert.NotEmpty(t, product["price"])
	assert.NotEmpty(t, product["stock"])
	assert.NotEmpty(t, product["image_url"])
	assert.NotEmpty(t, product["description"])

	fmt.Println("TestGetOrderWithValidId: ✅")
}

func TestGetOrderWithInvalidId(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `query {
			get_order(id: "xxx"){
				id, status, total_price, 
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

	responseBody := ResponseBodyParser(response.Body)

	assert.Nil(t, responseBody["data"])

	errors, ok := responseBody["errors"].([]any)
	assert.True(t, ok)
	assert.NotEmpty(t, errors)

	fmt.Println("TestGetOrderWithInvalidId: ✅")
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

	response, err := App.Test(request, -1)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)

	data, ok := responseBody["data"].(map[string]any)
	assert.True(t, ok)

	order, ok := data["update_order"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, order["id"])
	assert.Equal(t, "Completed", order["status"])
	assert.Equal(t, float64(12500), order["total_price"])

	customer, ok := order["customer"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, customer["id"])
	assert.NotEmpty(t, customer["name"])
	assert.NotEmpty(t, customer["email"])
	assert.NotEmpty(t, customer["role"])
	assert.NotEmpty(t, customer["phone"])
	assert.NotEmpty(t, customer["address"])

	products, ok := order["products"].([]any)
	assert.True(t, ok)

	product := products[0].(map[string]any)
	assert.NotEmpty(t, product["id"])
	assert.NotEmpty(t, product["name"])
	assert.NotEmpty(t, product["price"])
	assert.NotEmpty(t, product["stock"])
	assert.NotEmpty(t, product["image_url"])
	assert.NotEmpty(t, product["description"])

	fmt.Println("TestUpdateOrder: ✅")
}
