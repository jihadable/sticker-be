package graphql

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestPostMessage(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `mutation {
			post_message(conversation_id: "` + ConversationId + `", message: "test message"){
				id, message, sender
			}
		}`,
	})
	request := httptest.NewRequest(fiber.MethodPost, "/graphql", requestBody)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer"+CustomerJWT)

	response, err := App.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)

	data, ok := responseBody["data"].(map[string]any)
	assert.True(t, ok)

	postMessage, ok := data["post_message"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, postMessage["id"])
	assert.Equal(t, "test message", postMessage["message"])

	sender, ok := postMessage["sender"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, sender["id"])
	assert.NotEmpty(t, sender["name"])
	assert.NotEmpty(t, sender["email"])
	assert.NotEmpty(t, sender["role"])
	assert.NotEmpty(t, sender["phone"])
	assert.NotEmpty(t, sender["address"])

	t.Log("✅")
}

func TestPostMessageWithProductReply(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `mutation {
			post_message(conversation_id: "` + ConversationId + `", product_id: "` + ProductId + `", message: "test message"){
				id, product, message, sender
			}
		}`,
	})
	request := httptest.NewRequest(fiber.MethodPost, "/graphql", requestBody)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer"+CustomerJWT)

	response, err := App.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)

	data, ok := responseBody["data"].(map[string]any)
	assert.True(t, ok)

	postMessage, ok := data["post_message"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, postMessage["id"])
	assert.Equal(t, "test message", postMessage["message"])

	product, ok := postMessage["product"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, product["id"])
	assert.NotEmpty(t, product["name"])
	assert.NotEmpty(t, product["price"])
	assert.NotEmpty(t, product["stock"])
	assert.NotEmpty(t, product["image_url"])
	assert.NotEmpty(t, product["description"])

	sender, ok := postMessage["sender"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, sender["id"])
	assert.NotEmpty(t, sender["name"])
	assert.NotEmpty(t, sender["email"])
	assert.NotEmpty(t, sender["role"])
	assert.NotEmpty(t, sender["phone"])
	assert.NotEmpty(t, sender["address"])

	t.Log("✅")
}

func TestPostMessageWithCustomProductReply(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `mutation {
			post_message(conversation_id: "` + ConversationId + `", custom_product_id: "` + CustomProductId + `", message: "test message"){
				id, custom_product, message, sender
			}
		}`,
	})
	request := httptest.NewRequest(fiber.MethodPost, "/graphql", requestBody)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer"+CustomerJWT)

	response, err := App.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)

	data, ok := responseBody["data"].(map[string]any)
	assert.True(t, ok)

	postMessage, ok := data["post_message"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, postMessage["id"])
	assert.Equal(t, "test message", postMessage["message"])

	customProduct, ok := postMessage["custom_product"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, customProduct["id"])
	assert.NotEmpty(t, customProduct["name"])
	assert.NotEmpty(t, customProduct["image_url"])

	sender, ok := postMessage["sender"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, sender["id"])
	assert.NotEmpty(t, sender["name"])
	assert.NotEmpty(t, sender["email"])
	assert.NotEmpty(t, sender["role"])
	assert.NotEmpty(t, sender["phone"])
	assert.NotEmpty(t, sender["address"])

	t.Log("✅")
}
