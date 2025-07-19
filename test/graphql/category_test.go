package graphql

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestCreateCategoryWithValidPayload1(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `mutation {
			post_category(id: "cat"){ id, products }
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

	postCategory, ok := data["post_category"].(map[string]any)
	assert.True(t, ok)

	assert.Equal(t, "cat", postCategory["id"])
	CategoryId = postCategory["id"].(string)
	assert.Empty(t, postCategory["products"])

	t.Log("✅")
}

func TestCreateCategoryWithInvalidPayload(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `mutation {
			post_category(){ id, products }
		}`,
	})
	request := httptest.NewRequest(fiber.MethodPost, "/graphql", requestBody)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+AdminJWT)

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

func TestGetCategories(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"qeury": `mutation {
			categories { id, products }
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

	categories, ok := data["categories"].([]map[string]any)
	assert.True(t, ok)

	category := categories[0]
	assert.NotEmpty(t, category["id"])
	assert.Empty(t, category["products"])

	t.Log("✅")
}

func TestGetCategoryWithValidId(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `query {
			category(id: "` + CategoryId + `"){ id, products }
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

	category, ok := data["category"].(map[string]any)
	assert.True(t, ok)

	assert.Equal(t, CategoryId, category["id"])
	assert.Empty(t, category["products"])

	t.Log("✅")
}

func TestGetCategoryWithInvalidId(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `query {
			category(id: "xxx"){ id, products }
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

func TestDeleteCategory(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `mutation {
			delete_category(id: "` + CategoryId + `")
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

	deleteCategory, ok := data["delete_category"].(bool)
	assert.True(t, ok)
	assert.True(t, deleteCategory)

	t.Log("✅")
}

func TestCreateCategoryWithValidPayload2(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `mutation {
			post_category(id: "cat"){ id, products }
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

	postCategory, ok := data["post_category"].(map[string]any)
	assert.True(t, ok)

	assert.Equal(t, "cat", postCategory["id"])
	CategoryId = postCategory["id"].(string)
	assert.Empty(t, postCategory["products"])

	t.Log("✅")
}
