package graphql

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestCreateCustomProductWithValidPayload1(t *testing.T) {
	filePath := filepath.Join("..", "..", "static", "redis.png")
	file, err := os.Open(filePath)
	assert.Nil(t, err)
	defer file.Close()

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	operations := `
	{
		"query": "mutation($image: Upload!){ post_custom_product(name: \"custom product test\", image: $image){ id, name, image_url }}",
		"variables": {
			"image": null
		}
	}`
	_ = writer.WriteField("operations", operations)
	_ = writer.WriteField("map", `{ "0": ["variables.image"] }`)

	part, err := writer.CreateFormFile("0", filepath.Base(filePath))
	assert.Nil(t, err)
	_, err = io.Copy(part, file)
	assert.Nil(t, err)

	writer.Close()

	request := httptest.NewRequest(fiber.MethodPost, "/graphql", &requestBody)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	request.Header.Set("Authorization", "Bearer "+CustomerJWT)

	response, err := App.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)

	data, ok := responseBody["data"].(map[string]any)
	assert.True(t, ok)

	postCustomProduct, ok := data["post_custom_product"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, postCustomProduct["id"])
	CustomProductId = postCustomProduct["id"].(string)
	assert.Equal(t, "custom product test", postCustomProduct["name"])
	assert.NotEmpty(t, postCustomProduct["image_url"])

	t.Log("✅")
}

func TestCreateCustomProductWithInvalidPayload(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `mutation {
			post_custom_product(){ id, name, image_url }
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

func TestGetCustomProductsByUser(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `query {
			custom_products_by_user { id, name, image_url }
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

	customProducts, ok := data["custom_products_by_user"].([]map[string]any)
	assert.True(t, ok)

	customProduct := customProducts[0]
	assert.NotEmpty(t, customProduct["id"])
	assert.NotEmpty(t, customProduct["name"])
	assert.NotEmpty(t, customProduct["image_url"])

	t.Log("✅")
}

func TestGetCustomProductWithValidId(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `query {
			custom_product(id: "` + CustomProductId + `"){ id, name, image_url }
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

	customProduct, ok := data["custom_product"].(map[string]any)
	assert.True(t, ok)

	assert.Equal(t, CustomProductId, customProduct["id"])
	assert.Equal(t, "custom product test", customProduct["name"])
	assert.NotEmpty(t, customProduct["image_url"])

	t.Log("✅")
}

func TestGetCustomProductWithInvalidId(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `query {
			custom_product(id: "xxx"){ id, name, image_url }
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

func TestUpdateCustomProduct(t *testing.T) {
	filePath := filepath.Join("..", "..", "static", "redis.png")
	file, err := os.Open(filePath)
	assert.Nil(t, err)
	defer file.Close()

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	operations := `
	{
		"query": "mutation($id: ID!, $image: Upload!){ update_custom_product(id: $id, name: \"update custom product test\", image: $image){ id, name, image_url }}",
		"variables": {
			"id": null,
			"image": null
		}
	}`
	_ = writer.WriteField("operations", operations)
	_ = writer.WriteField("map", `{ "0": ["variables.id"], "1": ["variables.image"] }`)
	_ = writer.WriteField("0", ProductId)

	part, err := writer.CreateFormFile("1", filepath.Base(filePath))
	assert.Nil(t, err)
	_, err = io.Copy(part, file)
	assert.Nil(t, err)

	writer.Close()

	request := httptest.NewRequest(fiber.MethodPost, "/graphql", &requestBody)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	request.Header.Set("Authorization", "Bearer "+CustomerJWT)

	response, err := App.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)

	data, ok := responseBody["data"].(map[string]any)
	assert.True(t, ok)

	updateCustomProduct, ok := data["update_custom_product"].(map[string]any)
	assert.True(t, ok)

	assert.Equal(t, ProductId, updateCustomProduct["id"])
	assert.Equal(t, "update custom product test", updateCustomProduct["name"])
	assert.NotEmpty(t, updateCustomProduct["image_url"])

	t.Log("✅")
}

func TestDeleteCustomProduct(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `mutation {
			delete_custom_product(id: "` + CustomProductId + `")
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

	deleteProduct, ok := data["delete_category"].(bool)
	assert.True(t, ok)
	assert.True(t, deleteProduct)

	t.Log("✅")
}

func TestCreateCustomProductWithValidPayload2(t *testing.T) {
	filePath := filepath.Join("..", "..", "static", "redis.png")
	file, err := os.Open(filePath)
	assert.Nil(t, err)
	defer file.Close()

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	operations := `
	{
		"query": "mutation($image: Upload!){ post_custom_product(name: \"custom product test\", image: $image){ id, name, image_url }}",
		"variables": {
			"image": null
		}
	}`
	_ = writer.WriteField("operations", operations)
	_ = writer.WriteField("map", `{ "0": ["variables.image"] }`)

	part, err := writer.CreateFormFile("0", filepath.Base(filePath))
	assert.Nil(t, err)
	_, err = io.Copy(part, file)
	assert.Nil(t, err)

	writer.Close()

	request := httptest.NewRequest(fiber.MethodPost, "/graphql", &requestBody)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	request.Header.Set("Authorization", "Bearer "+CustomerJWT)

	response, err := App.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)

	data, ok := responseBody["data"].(map[string]any)
	assert.True(t, ok)

	postCustomProduct, ok := data["post_custom_product"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, postCustomProduct["id"])
	CustomProductId = postCustomProduct["id"].(string)
	assert.Equal(t, "custom product test", postCustomProduct["name"])
	assert.NotEmpty(t, postCustomProduct["image_url"])

	t.Log("✅")
}
