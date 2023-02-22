package helpers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"sqlc-rest-api/requests"
	"sqlc-rest-api/responses"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
)

type GraphQLRequest struct {
	Query         string         `json:"query"`
	OperationName string         `json:"operation_name"`
	Variables     map[string]any `json:"variables"`
}

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

func NewProductsTest(n int, userID int64) *responses.Products {
	var productEdges []*responses.ProductEdge
	for i := 0; i < n; i++ {
		tt := time.Now()
		edge := responses.ProductEdge{
			Cursor: base64.StdEncoding.EncodeToString([]byte(tt.String())),
			Node: &responses.Product{
				ID:        int64(i + 1),
				Name:      fmt.Sprintf("Product %d", i+1),
				Price:     100,
				UserID:    userID,
				CreatedAt: tt,
			},
		}

		productEdges = append(productEdges, &edge)
	}

	if len(productEdges) < 1 {
		return &responses.Products{
			Edges:    []*responses.ProductEdge{},
			PageInfo: NewPageInfo("", "", false),
		}
	}

	sc := base64.StdEncoding.EncodeToString([]byte(productEdges[0].Node.CreatedAt.String()))
	ec := base64.StdEncoding.EncodeToString([]byte(productEdges[len(productEdges)-1].Node.CreatedAt.String()))
	pageInfo := NewPageInfo(sc, ec, true)

	return &responses.Products{
		Edges:    productEdges,
		PageInfo: pageInfo,
	}
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

func NewGraphQLRequestTest(name, query string, vars map[string]any) GraphQLRequest {
	return GraphQLRequest{
		Query:         query,
		OperationName: name,
		Variables:     vars,
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

func NewGetUserProductsRequestTest(userID int64, first *int, after *string) requests.GetUserProductsRequest {
	return requests.GetUserProductsRequest{
		UserID: userID,
		First:  first,
		After:  after,
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

func GraphProductMatchTest(t *testing.T, jsonPath string, body bytes.Buffer, expectedProduct responses.Product) {
	jsonData, err := io.ReadAll(&body)
	require.NoError(t, err)

	var product responses.Product
	parseJson(t, jsonData, jsonPath, &product)

	require.Equal(t, expectedProduct.ID, product.ID)
	require.Equal(t, expectedProduct.Name, product.Name)
	require.Equal(t, expectedProduct.Price, product.Price)
	require.Equal(t, expectedProduct.UserID, product.UserID)
}

func GraphProductsMatchTest(t *testing.T, jsonPath string, body bytes.Buffer, expectedProducts responses.Products) {
	jsonData, err := io.ReadAll(&body)
	require.NoError(t, err)

	var products responses.Products
	parseJson(t, jsonData, jsonPath, &products)

	require.Equal(t, len(expectedProducts.Edges), len(products.Edges))
}

func GraphUserMatchTest(t *testing.T, jsonPath string, body bytes.Buffer, expectedUser responses.User) {
	jsonData, err := io.ReadAll(&body)
	require.NoError(t, err)

	var user responses.User
	parseJson(t, jsonData, jsonPath, &user)

	require.Equal(t, expectedUser.ID, user.ID)
	require.Equal(t, expectedUser.Name, user.Name)
	require.Equal(t, expectedUser.Email, user.Email)
}

func GraphExpectComplexityLimit(t *testing.T, jsonPath string, body bytes.Buffer) {
	jsonData, err := io.ReadAll(&body)
	require.NoError(t, err)
	fmt.Println("jsondata", string(jsonData))
	var complexity string
	parseJson(t, jsonData, jsonPath, &complexity)

	require.Equal(t, "COMPLEXITY_LIMIT_EXCEEDED", complexity)
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
