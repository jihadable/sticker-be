package graphql

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestCreateProductCategoryWithValidPayload(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `mutation {
			create_product_category(product_id: "` + ProductId + `", category_id: "` + CategoryId + `"){
				product { id, name, price, stock, image_url, description },
				category { id }
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

	productCategory, ok := data["create_product_category"].(map[string]any)
	assert.True(t, ok)

	product, ok := productCategory["product"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, product["id"])
	assert.NotEmpty(t, product["name"])
	assert.NotEmpty(t, product["price"])
	assert.NotEmpty(t, product["stock"])
	assert.NotEmpty(t, product["image_url"])
	assert.NotEmpty(t, product["description"])

	category, ok := productCategory["category"].(map[string]any)
	assert.True(t, ok)

	assert.NotEmpty(t, category["id"])

	fmt.Println("TestCreateProductCategoryWithValidPayload: ✅")
}

func TestCreateProductCategoryWithInvalidPayload(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `mutation {
			create_product_category(){
				product { id, name, price, stock, image_url, description },
				category { id }
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

	fmt.Println("TestCreateProductCategoryWithInvalidPayload: ✅")
}

func TestDeleteProductCategory(t *testing.T) {
	requestBody := RequestBodyParser(map[string]string{
		"query": `mutation {
			delete_product_category(product_id: "` + ProductId + `", category_id: "` + CategoryId + `")
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

	deleteProductCategory, ok := data["delete_product_category"].(bool)
	assert.True(t, ok)
	assert.True(t, deleteProductCategory)

	fmt.Println("TestDeleteProductCategory: ✅")
}
