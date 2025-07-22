package graphql

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetNotificationsByRecipient(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `query {
			get_notifications_by_recipient {
				id, type, title, message, is_read,
				recipient { id, name, email, role }
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

	notifications, ok := data["get_notifications_by_recipient"].([]any)
	assert.True(t, ok)

	notification, ok := notifications[0].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, notification["id"])
	NotificationId = notification["id"].(string)
	assert.NotEmpty(t, notification["type"])
	assert.NotEmpty(t, notification["title"])
	assert.NotEmpty(t, notification["message"])
	assert.Equal(t, false, notification["is_read"])

	recipient, ok := notification["recipient"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, recipient["id"])
	assert.NotEmpty(t, recipient["name"])
	assert.NotEmpty(t, recipient["email"])
	assert.NotEmpty(t, recipient["role"])

	fmt.Println("TestGetNotificationsByRecipient: ✅")
}

func TestReadNotificationByValidId(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `mutation {
			read_notification(id: "` + NotificationId + `")
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

	readNotification, ok := data["read_notification"].(bool)
	assert.True(t, ok)
	assert.True(t, readNotification)

	fmt.Println("TestReadNotificationByValidId: ✅")
}

func TestReadNotificationByInvalidId(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `mutation {
			read_notification(id: "xxx")
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

	fmt.Println("TestReadNotificationByInvalidId: ✅")
}

func TestReadAllNotifications(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `mutation {
			read_all_notifications
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

	readNotifications, ok := data["read_all_notifications"].(bool)
	assert.True(t, ok)
	assert.True(t, readNotifications)

	fmt.Println("TestReadAllNotifications: ✅")
}
