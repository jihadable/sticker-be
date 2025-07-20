package graphql

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetConversationByUser(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `query {
			get_conversation_by_user { id }		
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

	conversation, ok := data["get_conversation_by_user"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, conversation["id"])
	ConversationId = conversation["id"].(string)

	t.Log("✅")
}
