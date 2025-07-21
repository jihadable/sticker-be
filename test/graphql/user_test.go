package graphql

import (
	"fmt"
	"net/http/httptest"
	"os"
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
				password: "test password"
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

	response, err := App.Test(request, -1)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)
	fmt.Println(responseBody)

	data, ok := responseBody["data"].(map[string]any)
	assert.True(t, ok)

	register, ok := data["register"].(map[string]any)
	assert.True(t, ok)

	token, ok := register["token"].(string)
	assert.True(t, ok)
	assert.NotEmpty(t, token)
	CustomerJWT = token

	user, ok := register["user"].(map[string]any)
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

	user, ok := data["get_user"].(map[string]any)
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

func TestUpdateUser(t *testing.T) {
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

	updateUser, ok := data["update_user"].(map[string]any)
	assert.True(t, ok)

	token, ok := updateUser["token"].(string)
	assert.True(t, ok)
	assert.NotEmpty(t, token)
	CustomerJWT = token

	user, ok := updateUser["user"].(map[string]any)
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
				password: "test password"
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

	login, ok := data["login"].(map[string]any)
	assert.True(t, ok)

	token, ok := login["token"].(string)
	assert.True(t, ok)
	assert.NotEmpty(t, token)
	CustomerJWT = token

	user, ok := login["user"].(map[string]any)
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
				email: "stikeradmin@gmail.com"
				password: "` + os.Getenv("PRIVATE_PASSWORD") + `"
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

	login, ok := data["login"].(map[string]any)
	assert.True(t, ok)

	token, ok := login["token"].(string)
	assert.True(t, ok)
	assert.NotEmpty(t, token)
	AdminJWT = token

	user, ok := login["user"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, user["id"])
	assert.Equal(t, "Stiker Admin", user["name"])
	assert.Equal(t, "stikeradmin@gmail.com", user["email"])
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
