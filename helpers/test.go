package helpers

import (
	"bytes"
	"encoding/json"
	"io"
	"sqlc-rest-api/requests"
	"sqlc-rest-api/responses"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
)

func NewProductTest(user responses.User) responses.Product {
	product := responses.Product{
		ID:        1,
		Name:      "Test Product",
		Price:     100,
		UserID:    user.ID,
		CreatedAt: time.Now(),
	}

	return product
}

func NewProductDeletedTest(productID int64) *responses.DeletedProduct {
	return &responses.DeletedProduct{
		Deleted:   true,
		ProductID: productID,
	}
}

func NewUserTest() responses.User {
	return responses.User{
		ID:        1,
		Name:      "royyan",
		Email:     "roy@gmail.com",
		CreatedAt: time.Now(),
	}
}

func NewCreateProductRequestTest(user *responses.User, product *responses.Product) requests.CreateProductRequest {
	if user == nil && product == nil {
		return requests.CreateProductRequest{}
	}

	return requests.CreateProductRequest{
		UserID: user.ID,
		Name:   product.Name,
		Price:  product.Price,
	}
}

func NewBindUriIDRequestTest(id int64) requests.BindUriID {
	return requests.BindUriID{
		ID: id,
	}
}

func NewUpdateProductRequestTest(product *responses.Product) requests.UpdateProductRequest {
	if product == nil {
		return requests.UpdateProductRequest{}
	}

	return requests.UpdateProductRequest{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}
}

func RequireProductMatchTest(t *testing.T, body *bytes.Buffer, expectedProduct responses.Product) {
	jsonData, err := io.ReadAll(body)
	require.NoError(t, err)

	var product responses.Product
	parseJson(t, jsonData, "data.product", &product)

	// require.Equal(t, expectedProduct, product)
	require.Equal(t, expectedProduct.ID, product.ID)
	require.Equal(t, expectedProduct.Name, product.Name)
	require.Equal(t, expectedProduct.Price, product.Price)
	require.Equal(t, expectedProduct.UserID, product.UserID)
}

func parseJson(t *testing.T, data []byte, path string, placeholder any) {
	valid := gjson.ValidBytes(data)
	require.True(t, valid)

	result := gjson.GetBytes(data, path)
	exists := result.Exists()
	require.True(t, exists)

	raw := []byte(result.Raw)
	err := json.Unmarshal(raw, placeholder)
	require.NoError(t, err)
}
