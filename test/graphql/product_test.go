package graphql

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestCreateProductWithValidPayload1(t *testing.T) {
	filePath := filepath.Join("..", "..", "static", "redis.png")
	file, err := os.Open(filePath)
	assert.Nil(t, err)
	defer file.Close()

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	operations := `
	{
		"query": "mutation($image: Upload!){ create_product(name: \"product test\", price: 1000, stock: 100, description: \"desc test\", image: $image){ id, name, price, stock, image_url, description, categories { id } } }",
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

	response, err := App.Test(request, -1)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)

	data, ok := responseBody["data"].(map[string]any)
	assert.True(t, ok)

	product, ok := data["create_product"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, product["id"])
	ProductId = product["id"].(string)
	assert.Equal(t, "product test", product["name"])
	assert.Equal(t, float64(1000), product["price"])
	assert.Equal(t, float64(100), product["stock"])
	assert.NotEmpty(t, product["image_url"])
	assert.Equal(t, "desc test", product["description"])
	assert.Empty(t, product["categories"])

	fmt.Println("TestCreateProductWithValidPayload1: ✅")
}

func TestCreateProductWithInvalidPayload(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `mutation {
			create_product(){
				id, name, price, stock, image_url, description,
				categories { id }
			}	
		}`,
	})
	request := httptest.NewRequest(fiber.MethodPost, "/graphql", requestBody)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+AdminJWT)

	response, err := App.Test(request, -1)

	assert.Nil(t, err)

	responseBody := ResponseBodyParser(response.Body)

	assert.Nil(t, responseBody["data"])

	errors, ok := responseBody["errors"].([]any)
	assert.True(t, ok)
	assert.NotEmpty(t, errors)

	fmt.Println("TestCreateProductWithInvalidPayload: ✅")
}

func TestGetProducts(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `query {
			get_products {
				id, name, price, stock, image_url, description, 
				categories { id }
			}	
		}`,
	})
	request := httptest.NewRequest(fiber.MethodPost, "/graphql", requestBody)
	request.Header.Set("Content-Type", "application/json")

	response, err := App.Test(request, -1)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)

	data, ok := responseBody["data"].(map[string]any)
	assert.True(t, ok)

	products, ok := data["get_products"].([]any)
	assert.True(t, ok)

	product := products[0].(map[string]any)
	assert.NotEmpty(t, product["id"])
	assert.NotEmpty(t, product["name"])
	assert.NotEmpty(t, product["price"])
	assert.NotEmpty(t, product["stock"])
	assert.NotEmpty(t, product["image_url"])
	assert.NotEmpty(t, product["description"])
	assert.Empty(t, product["categories"])

	fmt.Println("TestGetProducts: ✅")
}

func TestGetProductWithValidId(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `query {
			get_product(id: "` + ProductId + `"){
				id, name, price, stock, image_url, description,
				categories { id }
			}
		}`,
	})
	request := httptest.NewRequest(fiber.MethodPost, "/graphql", requestBody)
	request.Header.Set("Content-Type", "application/json")

	response, err := App.Test(request, -1)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)

	data, ok := responseBody["data"].(map[string]any)
	assert.True(t, ok)

	product, ok := data["get_product"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, product["id"])
	assert.NotEmpty(t, product["name"])
	assert.NotEmpty(t, product["price"])
	assert.NotEmpty(t, product["stock"])
	assert.NotEmpty(t, product["image_url"])
	assert.NotEmpty(t, product["description"])
	assert.Empty(t, product["categories"])

	fmt.Println("TestGetProductWithValidId: ✅")
}

func TestGetProductWithInvalidId(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `query {
			get_product(id: "xxx"){
				id, name, price, stock, image_url, description,
				categories { id }
			}
		}`,
	})
	request := httptest.NewRequest(fiber.MethodPost, "/graphql", requestBody)
	request.Header.Set("Content-Type", "application/json")

	response, err := App.Test(request, -1)

	assert.Nil(t, err)

	responseBody := ResponseBodyParser(response.Body)

	assert.Nil(t, responseBody["data"])

	errors, ok := responseBody["errors"].([]any)
	assert.True(t, ok)
	assert.NotEmpty(t, errors)

	fmt.Println("TestGetProductWithInvalidId: ✅")
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
		"query": "mutation($image: Upload!){ update_product(id: \"` + ProductId + `\", name: \"update product test\", price: 1500, stock: 10, description: \"update desc test\", image: $image){ id, name, price, stock, image_url, description, categories { id } } }",
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

	response, err := App.Test(request, -1)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)

	data, ok := responseBody["data"].(map[string]any)
	assert.True(t, ok)

	product, ok := data["update_product"].(map[string]any)
	assert.True(t, ok)

	assert.Equal(t, ProductId, product["id"])
	assert.Equal(t, "update product test", product["name"])
	assert.Equal(t, float64(1500), product["price"])
	assert.Equal(t, float64(10), product["stock"])
	assert.NotEmpty(t, product["image_url"])
	assert.Equal(t, "update desc test", product["description"])
	assert.Empty(t, product["categories"])

	fmt.Println("TestUpdateProduct: ✅")
}

func TestDeleteProduct(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `mutation {
			delete_product(id: "` + ProductId + `")
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

	deleteProduct, ok := data["delete_product"].(bool)
	assert.True(t, ok)
	assert.True(t, deleteProduct)

	fmt.Println("TestDeleteProduct: ✅")
}

func TestCreateProductWithValidPayload2(t *testing.T) {
	filePath := filepath.Join("..", "..", "static", "redis.png")
	file, err := os.Open(filePath)
	assert.Nil(t, err)
	defer file.Close()

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	operations := `
	{
		"query": "mutation($image: Upload!){ create_product(name: \"product test\", price: 1000, stock: 100, description: \"desc test\", image: $image){ id, name, price, stock, image_url, description, categories { id } } }",
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

	response, err := App.Test(request, -1)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)

	data, ok := responseBody["data"].(map[string]any)
	assert.True(t, ok)

	product, ok := data["create_product"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, product["id"])
	ProductId = product["id"].(string)
	assert.Equal(t, "product test", product["name"])
	assert.Equal(t, float64(1000), product["price"])
	assert.Equal(t, float64(100), product["stock"])
	assert.NotEmpty(t, product["image_url"])
	assert.Equal(t, "desc test", product["description"])
	assert.Empty(t, product["categories"])

	fmt.Println("TestCreateProductWithValidPayload2: ✅")
}
