package graphql

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestCreateProductCategory(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `mutation {
			post_product_category(product_id: "` + ProductId + `", category_id: "` + CategoryId + `"){
				product { id, name, price, stock, image_url, description }, 
				category { id }
			}
		}`,
	})
	request := httptest.NewRequest(fiber.MethodPost, "/graphql", requestBody)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer"+AdminJWT)

	response, err := App.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)

	data, ok := responseBody["data"].(map[string]any)
	assert.True(t, ok)

	postProductCategory, ok := data["post_product_category"].(map[string]any)
	assert.True(t, ok)

	product, ok := postProductCategory["product"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, product["id"])
	assert.NotEmpty(t, product["name"])
	assert.NotEmpty(t, product["price"])
	assert.NotEmpty(t, product["stock"])
	assert.NotEmpty(t, product["image_url"])
	assert.NotEmpty(t, product["description"])

	category, ok := postProductCategory["product"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, category["id"])

	t.Log("✅")
}

func TestDeleteProductCategory(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `mutation {
			delete_product_category(product_id: "` + ProductId + `", category_id: "` + CategoryId + `")
		}`,
	})
	request := httptest.NewRequest(fiber.MethodPost, "/graphql", requestBody)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer"+AdminJWT)

	response, err := App.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	responseBody := ResponseBodyParser(response.Body)

	data, ok := responseBody["data"].(map[string]any)
	assert.True(t, ok)

	deleteProductCategory, ok := data["delete_product_category"].(bool)
	assert.True(t, ok)
	assert.True(t, deleteProductCategory)

	t.Log("✅")
}
