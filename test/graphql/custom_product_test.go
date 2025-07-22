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
		"query": "mutation($image: Upload!){ create_custom_product(name: \"custom product test\", image: $image){ id, name, image_url, customer { id, name, email, role, phone, address } } }",
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

	response, err := App.Test(request, -1)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)

	data, ok := responseBody["data"].(map[string]any)
	assert.True(t, ok)

	customProduct, ok := data["create_custom_product"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, customProduct["id"])
	CustomProductId = customProduct["id"].(string)
	assert.Equal(t, "custom product test", customProduct["name"])
	assert.NotEmpty(t, customProduct["image_url"])

	customer, ok := customProduct["customer"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, customer["id"])
	assert.NotEmpty(t, customer["name"])
	assert.NotEmpty(t, customer["email"])
	assert.NotEmpty(t, customer["role"])
	assert.NotEmpty(t, customer["phone"])
	assert.NotEmpty(t, customer["address"])

	t.Log("✅")
}

func TestCreateCustomProductWithInvalidPayload(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `mutation {
			create_custom_product(){
				id, name, image_url,
				customer { id, name, email, role, phone, address }
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

func TestGetCustomProductsByUser(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `query {
			get_custom_products_by_customer {
				id, name, image_url, 
				customer { id, name, email, role, phone, address }
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

	customProducts, ok := data["get_custom_products_by_customer"].([]any)
	assert.True(t, ok)

	customProduct := customProducts[0].(map[string]any)
	assert.NotEmpty(t, customProduct["id"])
	assert.NotEmpty(t, customProduct["name"])
	assert.NotEmpty(t, customProduct["image_url"])

	customer, ok := customProduct["customer"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, customer["id"])
	assert.NotEmpty(t, customer["name"])
	assert.NotEmpty(t, customer["email"])
	assert.NotEmpty(t, customer["role"])
	assert.NotEmpty(t, customer["phone"])
	assert.NotEmpty(t, customer["address"])

	t.Log("✅")
}

func TestGetCustomProductWithValidId(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `query {
			get_custom_product(id: "` + CustomProductId + `"){
				id, name, image_url,
				customer { id, name, email, role, phone, address }
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

	customProduct, ok := data["get_custom_product"].(map[string]any)
	assert.True(t, ok)

	assert.Equal(t, CustomProductId, customProduct["id"])
	assert.Equal(t, "custom product test", customProduct["name"])
	assert.NotEmpty(t, customProduct["image_url"])

	customer, ok := customProduct["customer"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, customer["id"])
	assert.NotEmpty(t, customer["name"])
	assert.NotEmpty(t, customer["email"])
	assert.NotEmpty(t, customer["role"])
	assert.NotEmpty(t, customer["phone"])
	assert.NotEmpty(t, customer["address"])

	t.Log("✅")
}

func TestGetCustomProductWithInvalidId(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `query {
			get_custom_product(id: "xxx"){
				id, name, image_url,
				customer { id, name, email, role, phone, address }
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

func TestUpdateCustomProduct(t *testing.T) {
	filePath := filepath.Join("..", "..", "static", "redis.png")
	file, err := os.Open(filePath)
	assert.Nil(t, err)
	defer file.Close()

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	operations := `
	{
		"query": "mutation($image: Upload!){ update_custom_product(id: \"` + CustomProductId + `\", name: \"update custom product test\", image: $image){ id, name, image_url, customer { id, name, email, role, phone, address } } }",
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

	response, err := App.Test(request, -1)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)

	data, ok := responseBody["data"].(map[string]any)
	assert.True(t, ok)

	customProduct, ok := data["update_custom_product"].(map[string]any)
	assert.True(t, ok)

	assert.Equal(t, CustomProductId, customProduct["id"])
	assert.Equal(t, "update custom product test", customProduct["name"])
	assert.NotEmpty(t, customProduct["image_url"])

	customer, ok := customProduct["customer"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, customer["id"])
	assert.NotEmpty(t, customer["name"])
	assert.NotEmpty(t, customer["email"])
	assert.NotEmpty(t, customer["role"])
	assert.NotEmpty(t, customer["phone"])
	assert.NotEmpty(t, customer["address"])

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

	response, err := App.Test(request, -1)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)

	data, ok := responseBody["data"].(map[string]any)
	assert.True(t, ok)

	deleteProduct, ok := data["delete_custom_product"].(bool)
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
		"query": "mutation($image: Upload!){ create_custom_product(name: \"custom product test\", image: $image){ id, name, image_url, customer { id, name, email, role, phone, address } } }",
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

	response, err := App.Test(request, -1)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)

	data, ok := responseBody["data"].(map[string]any)
	assert.True(t, ok)

	customProduct, ok := data["create_custom_product"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, customProduct["id"])
	CustomProductId = customProduct["id"].(string)
	assert.Equal(t, "custom product test", customProduct["name"])
	assert.NotEmpty(t, customProduct["image_url"])

	customer, ok := customProduct["customer"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, customer["id"])
	assert.NotEmpty(t, customer["name"])
	assert.NotEmpty(t, customer["email"])
	assert.NotEmpty(t, customer["role"])
	assert.NotEmpty(t, customer["phone"])
	assert.NotEmpty(t, customer["address"])

	t.Log("✅")
}
