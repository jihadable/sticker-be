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

func TestPostProductWithValidPayload(t *testing.T) {
	filePath := filepath.Join("..", "..", "static", "redis.png")
	file, err := os.Open(filePath)
	assert.Nil(t, err)
	defer file.Close()

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	operations := `
	{
		"query": "mutation($image: Upload!){ post_product(name: \"product test\", price: 1000, stock: 100, description: \"desc test\", image: $image){ id, name, price, stock, image_url, description, categories }}",
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
	request.Header.Set("Authorization", "Bearer "+AdminJWT)

	response, err := App.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)

	data, ok := responseBody["data"].(map[string]any)
	assert.True(t, ok)

	postProduct, ok := data["post_product"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, postProduct["id"])
	ProductId = postProduct["id"].(string)
	assert.Equal(t, "product test", postProduct["name"])
	assert.Equal(t, 1000, postProduct["price"])
	assert.Equal(t, 100, postProduct["stock"])
	assert.NotEmpty(t, postProduct["image_url"])
	assert.Equal(t, "desc test", postProduct["description"])
	assert.Empty(t, postProduct["categories"])

	t.Log("✅")
}

func TestPostProductWithInvalidPayload(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `mutation {
			post_product(){ id, name, price, stock, image_url, description, categories }	
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

func TestGetProducts(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `query {
			products { id, name, price, stock, image_url, description, categories }	
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

	products, ok := data["products"].([]map[string]any)
	assert.True(t, ok)

	product := products[0]
	assert.NotEmpty(t, product["id"])
	assert.NotEmpty(t, product["name"])
	assert.NotEmpty(t, product["price"])
	assert.NotEmpty(t, product["stock"])
	assert.NotEmpty(t, product["image_url"])
	assert.NotEmpty(t, product["description"])
	assert.Empty(t, product["categories"])

	t.Log("✅")
}

func TestGetProductWithValidId(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `query {
			product(id: "` + ProductId + `"){
				id, name, price, stock, image_url, description, categories
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

	product, ok := data["product"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, product["id"])
	assert.NotEmpty(t, product["name"])
	assert.NotEmpty(t, product["price"])
	assert.NotEmpty(t, product["stock"])
	assert.NotEmpty(t, product["image_url"])
	assert.NotEmpty(t, product["description"])
	assert.Empty(t, product["categories"])

	t.Log("✅")
}

func TestGetProductWithInvalidId(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `query {
			product(id: "xxx"){ id, name, price, stock, image_url, description, categories }
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

func TestUpdateProduct(t *testing.T) {
	filePath := filepath.Join("..", "..", "static", "redis.png")
	file, err := os.Open(filePath)
	assert.Nil(t, err)
	defer file.Close()

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	operations := `
	{
		"query": "mutation($id: ID!, $image: Upload!){ update_product(id: $id, name: \"update product test\", price: 1500, stock: 10, description: \"update desc test\", image: $image){ id, name, price, stock, image_url, description, categories }}",
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
	request.Header.Set("Authorization", "Bearer "+AdminJWT)

	response, err := App.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)

	data, ok := responseBody["data"].(map[string]any)
	assert.True(t, ok)

	updateProduct, ok := data["update_product"].(map[string]any)
	assert.True(t, ok)

	assert.Equal(t, ProductId, updateProduct["id"])
	assert.Equal(t, "update product test", updateProduct["name"])
	assert.Equal(t, 1500, updateProduct["price"])
	assert.Equal(t, 10, updateProduct["stock"])
	assert.NotEmpty(t, updateProduct["image_url"])
	assert.Equal(t, "update desc test", updateProduct["description"])
	assert.Empty(t, updateProduct["categories"])

	t.Log("✅")
}

func TestDeleteProduct(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `mutation {
			delete_product(id: "` + ProductId + `")
		}`,
	})
	request := httptest.NewRequest(fiber.MethodPost, "/graphql", requestBody)

	response, err := App.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)

	data, ok := responseBody["data"].(map[string]any)
	assert.True(t, ok)

	deleteProduct, ok := data["delete_product"].(bool)
	assert.True(t, ok)
	assert.True(t, deleteProduct)

	t.Log("✅")
}
