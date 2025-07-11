package graphql

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestRegisterUserWithValidPayload(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `mutation {
			post_user(
				name: "user test"
				email: "usertest@gmail.com"
				password: "test"
				phone: "081234567890",
				address: "Jl. Langsat"
			){
				token
				user { id, name, email, role, phone, address, custom_products, orders }
			}
		}`,
	})
	request := httptest.NewRequest(fiber.MethodPost, "/graphql", requestBody)
	request.Header.Set("Content-Type", "application/json")

	response, err := App.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)

	data, ok := responseBody["data"].(map[string]any)
	assert.True(t, ok)

	postUser, ok := data["post_user"].(map[string]any)
	assert.True(t, ok)

	token, ok := postUser["token"].(string)
	assert.True(t, ok)
	assert.NotEmpty(t, token)
	CustomerJWT = token

	user, ok := postUser["user"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, user["id"])
	assert.Equal(t, "user test", user["name"])
	assert.Equal(t, "usertest@gmail.com", user["email"])
	assert.Equal(t, "customer", user["role"])
	assert.Equal(t, "081234567890", user["phone"])
	assert.Equal(t, "Jl. Langsat", user["address"])

	customProducts, ok := user["custom_products"].([]any)
	assert.True(t, ok)
	assert.Empty(t, customProducts)

	orders, ok := user["orders"].([]any)
	assert.True(t, ok)
	assert.Empty(t, orders)

	t.Log("✅")
}

func TestPostUserWithInvalidPayload(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `mutation {
			post_user(){
				token
				user { id, name, email, role, phone, address, custom_products, orders }
			}
		}`,
	})
	request := httptest.NewRequest(fiber.MethodPost, "/graphql", requestBody)
	request.Header.Set("Content-Type", "application/json")

	response, err := App.Test(request)

	assert.Nil(t, err)
	assert.NotEqual(t, fiber.StatusOK, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)
	fmt.Println(responseBody)

	assert.Nil(t, responseBody["data"])

	errors, ok := responseBody["errors"].([]any)
	assert.True(t, ok)
	assert.NotEmpty(t, errors)

	t.Log("✅")
}

func TestGetUserWithToken(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `query {
			user { id, name, email, role, phone, address, custom_products, orders }
		}`,
	})
	request := httptest.NewRequest(fiber.MethodPost, "/graphql", requestBody)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", CustomerJWT)

	response, err := App.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)

	data, ok := responseBody["data"].(map[string]any)
	assert.True(t, ok)

	user, ok := data["user"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, user["id"])
	assert.Equal(t, "user test", user["name"])
	assert.Equal(t, "usertest@gmail.com", user["email"])
	assert.Equal(t, "customer", user["role"])
	assert.Equal(t, "081234567890", user["phone"])
	assert.Equal(t, "Jl. Langsat", user["address"])

	customProducts, ok := user["custom_products"].([]any)
	assert.True(t, ok)
	assert.Empty(t, customProducts)

	orders, ok := user["orders"].([]any)
	assert.True(t, ok)
	assert.Empty(t, orders)

	t.Log("✅")
}

func TestGetUserWithoutToken(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `query {
			user { id, name, email, role, phone, address, custom_products, orders }
		}`,
	})
	request := httptest.NewRequest(fiber.MethodPost, "/graphql", requestBody)
	request.Header.Set("Content-Type", "application/json")

	response, err := App.Test(request)

	assert.Nil(t, err)
	assert.NotEqual(t, fiber.StatusOK, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)
	fmt.Println(responseBody)

	assert.Nil(t, responseBody["data"])

	errors, ok := responseBody["errors"].([]any)
	assert.True(t, ok)
	assert.NotEmpty(t, errors)

	t.Log("✅")
}

func TestUpadateUser(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `mutation {
			update_user(
				phone: "081122334455",
				address: "Jl. Durian"
			){ id, name, email, role, phone, address, custom_products, orders }
		}`,
	})
	request := httptest.NewRequest(fiber.MethodPost, "/graphql", requestBody)
	request.Header.Set("Content-Type", "application/json")

	response, err := App.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)

	data, ok := responseBody["data"].(map[string]any)
	assert.True(t, ok)

	postUser, ok := data["post_user"].(map[string]any)
	assert.True(t, ok)

	token, ok := postUser["token"].(string)
	assert.True(t, ok)
	assert.NotEmpty(t, token)
	CustomerJWT = token

	user, ok := postUser["user"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, user["id"])
	assert.Equal(t, "user test", user["name"])
	assert.Equal(t, "usertest@gmail.com", user["email"])
	assert.Equal(t, "customer", user["role"])
	assert.Equal(t, "081122334455", user["phone"])
	assert.Equal(t, "Jl. Durian", user["address"])

	customProducts, ok := user["custom_products"].([]any)
	assert.True(t, ok)
	assert.Empty(t, customProducts)

	orders, ok := user["orders"].([]any)
	assert.True(t, ok)
	assert.Empty(t, orders)

	t.Log("✅")
}

func TestLoginAsCustomer(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `mutation {
			verify_user(
				email: "usertest@gmail.com"
				password: "test"
			){
				token
				user { id, name, email, role, phone, address, custom_products, orders }
			}
		}`,
	})
	request := httptest.NewRequest(fiber.MethodPost, "/graphql", requestBody)
	request.Header.Set("Content-Type", "application/json")

	response, err := App.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)

	data, ok := responseBody["data"].(map[string]any)
	assert.True(t, ok)

	postUser, ok := data["post_user"].(map[string]any)
	assert.True(t, ok)

	token, ok := postUser["token"].(string)
	assert.True(t, ok)
	assert.NotEmpty(t, token)
	CustomerJWT = token

	user, ok := postUser["user"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, user["id"])
	assert.Equal(t, "user test", user["name"])
	assert.Equal(t, "usertest@gmail.com", user["email"])
	assert.Equal(t, "customer", user["role"])
	assert.Equal(t, "081122334455", user["phone"])
	assert.Equal(t, "Jl. Durian", user["address"])

	customProducts, ok := user["custom_products"].([]any)
	assert.True(t, ok)
	assert.Empty(t, customProducts)

	orders, ok := user["orders"].([]any)
	assert.True(t, ok)
	assert.Empty(t, orders)

	t.Log("✅")
}

func TestLoginAsAdmin(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `mutation {
			verify_user(
				email: "stickeradmin@gmail.com"
				password: "test"
			){
				token
				user { id, name, email, role }
			}
		}`,
	})
	request := httptest.NewRequest(fiber.MethodPost, "/graphql", requestBody)
	request.Header.Set("Content-Type", "application/json")

	response, err := App.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)

	data, ok := responseBody["data"].(map[string]any)
	assert.True(t, ok)

	postUser, ok := data["post_user"].(map[string]any)
	assert.True(t, ok)

	token, ok := postUser["token"].(string)
	assert.True(t, ok)
	assert.NotEmpty(t, token)
	AdminJWT = token

	user, ok := postUser["user"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, user["id"])
	assert.Equal(t, "sticker admin test", user["name"])
	assert.Equal(t, "stickeradmin@gmail.com", user["email"])
	assert.Equal(t, "admin", user["role"])

	t.Log("✅")
}
