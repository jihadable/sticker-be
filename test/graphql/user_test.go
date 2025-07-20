package graphql

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestRegisterUserWithValidPayload(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `mutation {
			register(
				name: "user test"
				email: "usertest@gmail.com"
				password: "test"
				phone: "081234567890",
				address: "Jl. Langsat"
			){
				token
				user { id, name, email, role, phone, address }
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

	t.Log("✅")
}

func TestPostUserWithInvalidPayload(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `mutation {
			register(){
				token
				user { id, name, email, role, phone, address }
			}
		}`,
	})
	request := httptest.NewRequest(fiber.MethodPost, "/graphql", requestBody)
	request.Header.Set("Content-Type", "application/json")

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

func TestGetUserWithToken(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `query {
			get_user { id, name, email, role, phone, address }
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

	user, ok := data["user"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, user["id"])
	assert.Equal(t, "user test", user["name"])
	assert.Equal(t, "usertest@gmail.com", user["email"])
	assert.Equal(t, "customer", user["role"])
	assert.Equal(t, "081234567890", user["phone"])
	assert.Equal(t, "Jl. Langsat", user["address"])

	t.Log("✅")
}

func TestGetUserWithoutToken(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `query {
			get_user { id, name, email, role, phone, address }
		}`,
	})
	request := httptest.NewRequest(fiber.MethodPost, "/graphql", requestBody)
	request.Header.Set("Content-Type", "application/json")

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

func TestUpadateUser(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `mutation {
			update_user(
				phone: "081122334455",
				address: "Jl. Durian"
			){ id, name, email, role, phone, address }
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

	t.Log("✅")
}

func TestLoginAsCustomer(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `mutation {
			login(
				email: "usertest@gmail.com"
				password: "test"
			){
				token
				user { id, name, email, role, phone, address }
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

	t.Log("✅")
}

func TestLoginAsAdmin(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `mutation {
			login(
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

func TestLoginWithInvalidPayload(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `mutation {
			login(){
				token
				user { id, name, email, role, phone, address }
			}
		}`,
	})
	request := httptest.NewRequest(fiber.MethodPost, "/graphql", requestBody)
	request.Header.Set("Content-Type", "application/json")

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
