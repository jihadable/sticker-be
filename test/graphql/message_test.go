package graphql

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestCreateMessageWithoutReply(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `mutation {
			create_message(conversation_id: "` + ConversationId + `", message: "test message"){
				id, message,
				sender { id, name, email, role, phone, address }
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

	message, ok := data["create_message"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, message["id"])
	assert.Equal(t, "test message", message["message"])

	sender, ok := message["sender"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, sender["id"])
	assert.NotEmpty(t, sender["name"])
	assert.NotEmpty(t, sender["email"])
	assert.NotEmpty(t, sender["role"])
	assert.NotEmpty(t, sender["phone"])
	assert.NotEmpty(t, sender["address"])

	t.Log("✅")
}

func TestCreateMessageWithInvalidPayload(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `mutation {
			create_message(){
				id, message,
				sender { id, name, email, role, phone, address }
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

	t.Log("✅")
}

func TestCreateMessageWithProductReply(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `mutation {
			create_message(conversation_id: "` + ConversationId + `", product_id: "` + ProductId + `", message: "test message"){
				id, message,
				product { id, name, price, stock, image_url, description }
				sender { id, name, email, role, phone, address }
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

	message, ok := data["create_message"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, message["id"])
	assert.Equal(t, "test message", message["message"])

	product, ok := message["product"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, product["id"])
	assert.NotEmpty(t, product["name"])
	assert.NotEmpty(t, product["price"])
	assert.NotEmpty(t, product["stock"])
	assert.NotEmpty(t, product["image_url"])
	assert.NotEmpty(t, product["description"])

	sender, ok := message["sender"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, sender["id"])
	assert.NotEmpty(t, sender["name"])
	assert.NotEmpty(t, sender["email"])
	assert.NotEmpty(t, sender["role"])
	assert.NotEmpty(t, sender["phone"])
	assert.NotEmpty(t, sender["address"])

	t.Log("✅")
}

func TestCreateMessageWithCustomProductReply(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `mutation {
			create_message(conversation_id: "` + ConversationId + `", custom_product_id: "` + CustomProductId + `", message: "test message"){
				id, message, 
				custom_product { id, name, image_url }
				sender { id, name, email, role, phone, address }
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

	message, ok := data["create_message"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, message["id"])
	assert.Equal(t, "test message", message["message"])

	customProduct, ok := message["custom_product"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, customProduct["id"])
	assert.NotEmpty(t, customProduct["name"])
	assert.NotEmpty(t, customProduct["image_url"])

	sender, ok := message["sender"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, sender["id"])
	assert.NotEmpty(t, sender["name"])
	assert.NotEmpty(t, sender["email"])
	assert.NotEmpty(t, sender["role"])
	assert.NotEmpty(t, sender["phone"])
	assert.NotEmpty(t, sender["address"])

	t.Log("✅")
}
