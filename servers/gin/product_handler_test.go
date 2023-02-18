package ginserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/http/httptest"
	"sqlc-rest-api/helpers"
	"sqlc-rest-api/requests"
	"sqlc-rest-api/responses"
	"testing"
	"time"

	"sqlc-rest-api/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
)

func TestCreateProduct(t *testing.T) {
	user := newUserTest()
	product := newProductTest(user)

	testCases := []struct {
		name          string
		req           requests.CreateProductRequest
		mock          func(service *mocks.MockService)
		checkResponse func(t *testing.T, rec *httptest.ResponseRecorder)
	}{
		{
			name: "product created successfully",
			req:  newCreateProductRequest(&user, &product),
			mock: func(service *mocks.MockService) {
				req := newCreateProductRequest(&user, &product)
				service.EXPECT().
					CreateProduct(gomock.Any(), gomock.Eq(req)).
					Times(1).
					Return(product, nil)
			},
			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, rec.Code)
				requireProductMatch(t, rec.Body, product)
			},
		},
		{
			name: "validation error product name not given",
			req:  newCreateProductRequest(nil, nil),
			mock: func(service *mocks.MockService) {
				req := newCreateProductRequest(nil, nil)
				service.EXPECT().
					CreateProduct(gomock.Any(), gomock.Eq(req)).
					Times(0)
			},
			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, rec.Code)
			},
		},
		{
			name: "internal server error",
			req:  newCreateProductRequest(&user, &product),
			mock: func(service *mocks.MockService) {
				req := newCreateProductRequest(&user, &product)
				service.EXPECT().
					CreateProduct(gomock.Any(), gomock.Eq(req)).
					Times(1).
					Return(responses.ProductResponse{}, fmt.Errorf("internal server error"))
			},
			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, rec.Code)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := mocks.NewMockService(ctrl)
			testCase.mock(service)

			data, err := json.Marshal(testCase.req)
			require.NoError(t, err)

			server := newGinTestServer(t, service)
			rec := httptest.NewRecorder()
			request, err := http.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(data))
			require.NoError(t, err)
			request.Header.Set("Content-Type", "application/json")

			server.Engine.ServeHTTP(rec, request)
			testCase.checkResponse(t, rec)
		})
	}
}

func TestDeleteProduct(t *testing.T) {
	user := newUserTest()
	product := newProductTest(user)

	testCases := []struct {
		name          string
		productID     int64
		mock          func(service *mocks.MockService)
		checkResponse func(t *testing.T, rec *httptest.ResponseRecorder)
	}{
		{
			name:      "product deleted successfully",
			productID: product.ID,
			mock: func(service *mocks.MockService) {
				req := newBindUriIDRequest(product.ID)
				service.EXPECT().
					DeleteProduct(gomock.Any(), gomock.Eq(req)).
					Times(1).
					Return(nil)
			},
			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, rec.Code)
			},
		},
		{
			name:      "validation error because given id lower than one",
			productID: 0,
			mock: func(service *mocks.MockService) {
				req := newBindUriIDRequest(0)
				service.EXPECT().
					DeleteProduct(gomock.Any(), gomock.Eq(req)).
					Times(0)
			},
			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, rec.Code)
			},
		},
		{
			name:      "product not found",
			productID: product.ID,
			mock: func(service *mocks.MockService) {
				req := newBindUriIDRequest(product.ID)
				service.EXPECT().
					DeleteProduct(gomock.Any(), gomock.Eq(req)).
					Times(1).
					Return(fmt.Errorf("product not found"))
			},
			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder) {
				// for now when product not found just return 500 error code
				require.Equal(t, http.StatusInternalServerError, rec.Code)
			},
		},
		{
			name:      "internal server error",
			productID: product.ID,
			mock: func(service *mocks.MockService) {
				req := newBindUriIDRequest(product.ID)
				service.EXPECT().
					DeleteProduct(gomock.Any(), gomock.Eq(req)).
					Times(1).
					Return(fmt.Errorf("internal server error"))
			},
			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, rec.Code)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := mocks.NewMockService(ctrl)
			server := newGinTestServer(t, service)

			testCase.mock(service)
			url := fmt.Sprintf("/products/%d", testCase.productID)
			rec := httptest.NewRecorder()
			request, err := http.NewRequest(http.MethodDelete, url, nil)
			require.NoError(t, err)

			server.Engine.ServeHTTP(rec, request)
			testCase.checkResponse(t, rec)
		})
	}
}

func TestGetProduct(t *testing.T) {
	user := newUserTest()
	product := newProductTest(user)

	testCases := []struct {
		name          string
		productID     int64
		mock          func(service *mocks.MockService)
		checkResponse func(t *testing.T, rec *httptest.ResponseRecorder)
	}{
		{
			name:      "get product successfully",
			productID: product.ID,
			mock: func(service *mocks.MockService) {
				req := newBindUriIDRequest(product.ID)
				service.EXPECT().
					GetProduct(gomock.Any(), gomock.Eq(req)).
					Times(1).
					Return(product, nil)
			},
			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, rec.Code)
				requireProductMatch(t, rec.Body, product)
			},
		},
		{
			name:      "validation error because given id lower than one",
			productID: 0,
			mock: func(service *mocks.MockService) {
				req := newBindUriIDRequest(0)
				service.EXPECT().
					GetProduct(gomock.Any(), gomock.Eq(req)).
					Times(0)
			},
			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, rec.Code)
			},
		},
		{
			name:      "product not found",
			productID: product.ID,
			mock: func(service *mocks.MockService) {
				req := newBindUriIDRequest(product.ID)
				service.EXPECT().
					GetProduct(gomock.Any(), gomock.Eq(req)).
					Times(1).
					Return(responses.ProductResponse{}, fmt.Errorf("product not found"))
			},
			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder) {
				// for now when product not found just return 500 error code
				require.Equal(t, http.StatusInternalServerError, rec.Code)
			},
		},
		{
			name:      "internal server error",
			productID: product.ID,
			mock: func(service *mocks.MockService) {
				req := newBindUriIDRequest(product.ID)
				service.EXPECT().
					GetProduct(gomock.Any(), gomock.Eq(req)).
					Times(1).
					Return(responses.ProductResponse{}, fmt.Errorf("internal server error"))
			},
			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, rec.Code)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := mocks.NewMockService(ctrl)
			server := newGinTestServer(t, service)

			testCase.mock(service)
			url := fmt.Sprintf("/products/%d", testCase.productID)
			rec := httptest.NewRecorder()
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.Engine.ServeHTTP(rec, request)
			testCase.checkResponse(t, rec)
		})
	}
}

func TestGetUserProducts(t *testing.T) {
	testCases := []struct {
		name          string
		page          int32
		perPage       int32
		userID        int32
		mock          func(service *mocks.MockService) ([]responses.ProductResponse, *responses.Pagination)
		checkResponse func(t *testing.T, rec *httptest.ResponseRecorder, products []responses.ProductResponse, pagination *responses.Pagination)
	}{
		{
			name:    "total products 100 per page 30 current page is 1",
			page:    1,
			perPage: 30,
			mock: func(service *mocks.MockService) ([]responses.ProductResponse, *responses.Pagination) {
				req := newListProductRequest(1, 30)
				products := newProductsTest(100, 1, 30)
				pagination := helpers.Paginate(100, int32(len(products)), 1, 30, "/products")

				service.EXPECT().
					GetUserProducts(gomock.Any(), gomock.Eq(req)).
					Times(1).
					Return(products, pagination, nil)

				return products, pagination
			},
			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder, products []responses.ProductResponse, pagination *responses.Pagination) {
				require.Equal(t, http.StatusOK, rec.Code)
				requireListProductsMatch(t, rec.Body, products, pagination)
			},
		},
		{
			name:    "total products 100 per page 30 current page is 2",
			page:    2,
			perPage: 30,
			mock: func(service *mocks.MockService) ([]responses.ProductResponse, *responses.Pagination) {
				req := newListProductRequest(2, 30)
				products := newProductsTest(100, 2, 30)
				pagination := helpers.Paginate(100, int32(len(products)), 2, 30, "/products")

				service.EXPECT().
					GetUserProducts(gomock.Any(), gomock.Eq(req)).
					Times(1).
					Return(products, pagination, nil)

				return products, pagination
			},
			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder, products []responses.ProductResponse, pagination *responses.Pagination) {
				require.Equal(t, http.StatusOK, rec.Code)
				requireListProductsMatch(t, rec.Body, products, pagination)
			},
		},
		{
			name:    "total products 100 per page 30 current page is 3",
			page:    3,
			perPage: 30,
			mock: func(service *mocks.MockService) ([]responses.ProductResponse, *responses.Pagination) {
				req := newListProductRequest(3, 30)
				products := newProductsTest(100, 3, 30)
				pagination := helpers.Paginate(100, int32(len(products)), 3, 30, "/products")

				service.EXPECT().
					GetUserProducts(gomock.Any(), gomock.Eq(req)).
					Times(1).
					Return(products, pagination, nil)

				return products, pagination
			},
			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder, products []responses.ProductResponse, pagination *responses.Pagination) {
				require.Equal(t, http.StatusOK, rec.Code)
				requireListProductsMatch(t, rec.Body, products, pagination)
			},
		},
		{
			name:    "total products 100 per page 30 current page is 4",
			page:    4,
			perPage: 30,
			mock: func(service *mocks.MockService) ([]responses.ProductResponse, *responses.Pagination) {
				req := newListProductRequest(4, 30)
				products := newProductsTest(100, 4, 30)
				pagination := helpers.Paginate(100, int32(len(products)), 4, 30, "/products")

				service.EXPECT().
					GetUserProducts(gomock.Any(), gomock.Eq(req)).
					Times(1).
					Return(products, pagination, nil)

				return products, pagination
			},
			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder, products []responses.ProductResponse, pagination *responses.Pagination) {
				require.Equal(t, http.StatusOK, rec.Code)
				requireListProductsMatch(t, rec.Body, products, pagination)
			},
		},
		{
			name:    "pagination over page",
			page:    2,
			perPage: 10,
			mock: func(service *mocks.MockService) ([]responses.ProductResponse, *responses.Pagination) {
				req := newListProductRequest(2, 10)
				products := newProductsTest(7, 2, 10)
				pagination := helpers.Paginate(7, int32(len(products)), 2, 10, "/products")

				service.EXPECT().
					GetUserProducts(gomock.Any(), gomock.Eq(req)).
					Times(1).
					Return(products, pagination, nil)

				return products, pagination
			},
			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder, products []responses.ProductResponse, pagination *responses.Pagination) {
				require.Equal(t, http.StatusOK, rec.Code)
				requireListProductsMatch(t, rec.Body, products, pagination)
			},
		},
		{
			name:    "validation error given page lower than one",
			page:    0,
			perPage: 10,
			mock: func(service *mocks.MockService) ([]responses.ProductResponse, *responses.Pagination) {
				req := newListProductRequest(0, 10)
				products := newProductsTest(10, 0, 10)
				pagination := helpers.Paginate(10, int32(len(products)), 0, 10, "/products")

				service.EXPECT().
					GetUserProducts(gomock.Any(), gomock.Eq(req)).
					Times(0)

				return products, pagination
			},
			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder, products []responses.ProductResponse, pagination *responses.Pagination) {
				require.Equal(t, http.StatusBadRequest, rec.Code)
			},
		},
		{
			name:    "validation error given per page lower than one",
			page:    1,
			perPage: 0,
			mock: func(service *mocks.MockService) ([]responses.ProductResponse, *responses.Pagination) {
				req := newListProductRequest(1, 0)
				products := newProductsTest(10, 1, 0)
				pagination := helpers.Paginate(10, int32(len(products)), 1, 0, "/products")

				service.EXPECT().
					GetUserProducts(gomock.Any(), gomock.Eq(req)).
					Times(0)

				return products, pagination
			},
			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder, products []responses.ProductResponse, pagination *responses.Pagination) {
				require.Equal(t, http.StatusBadRequest, rec.Code)
			},
		},
		{
			name:    "validation error given per page higher than 50",
			page:    1,
			perPage: 100,
			mock: func(service *mocks.MockService) ([]responses.ProductResponse, *responses.Pagination) {
				req := newListProductRequest(1, 100)
				products := newProductsTest(10, 1, 100)
				pagination := helpers.Paginate(10, int32(len(products)), 1, 100, "/products")

				service.EXPECT().
					GetUserProducts(gomock.Any(), gomock.Eq(req)).
					Times(0)

				return products, pagination
			},
			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder, products []responses.ProductResponse, pagination *responses.Pagination) {
				require.Equal(t, http.StatusBadRequest, rec.Code)
			},
		},
		{
			name:    "internal server error",
			page:    1,
			perPage: 10,
			mock: func(service *mocks.MockService) ([]responses.ProductResponse, *responses.Pagination) {
				req := newListProductRequest(1, 10)
				products := newProductsTest(10, 1, 10)
				pagination := helpers.Paginate(10, int32(len(products)), 1, 10, "/products")

				service.EXPECT().
					GetUserProducts(gomock.Any(), gomock.Eq(req)).
					Times(1).
					Return(nil, nil, fmt.Errorf("internal server error"))

				return products, pagination
			},
			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder, products []responses.ProductResponse, pagination *responses.Pagination) {
				require.Equal(t, http.StatusInternalServerError, rec.Code)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := mocks.NewMockService(ctrl)
			server := newGinTestServer(t, service)

			products, pagination := testCase.mock(service)
			url := fmt.Sprintf("/products?page=%d&per_page=%d", testCase.page, testCase.perPage)
			rec := httptest.NewRecorder()
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.Engine.ServeHTTP(rec, request)
			testCase.checkResponse(t, rec, products, pagination)
		})
	}
}

func TestUpdateProduct(t *testing.T) {
	user := newUserTest()
	product := newProductTest(user)

	testCases := []struct {
		name          string
		productID     int64
		req           requests.UpdateProductRequest
		mock          func(service *mocks.MockService)
		checkResponse func(t *testing.T, rec *httptest.ResponseRecorder)
	}{
		{
			name:      "product updated successfully",
			productID: product.ID,
			req:       newUpdateProductRequest(&product),
			mock: func(service *mocks.MockService) {
				req := newUpdateProductRequest(&product)

				service.EXPECT().
					UpdateProduct(gomock.Any(), gomock.Eq(req)).
					Times(1).
					Return(product, nil)
			},
			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, rec.Code)
				requireProductMatch(t, rec.Body, product)
			},
		},
		{
			name:      "validation error product id given lower than 1",
			productID: 0,
			req:       newUpdateProductRequest(nil),
			mock: func(service *mocks.MockService) {
				req := newUpdateProductRequest(nil)

				service.EXPECT().
					UpdateProduct(gomock.Any(), gomock.Eq(req)).
					Times(0)
			},
			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, rec.Code)
			},
		},
		{
			name:      "validation error product name not given",
			productID: product.ID,
			req:       newUpdateProductRequest(nil),
			mock: func(service *mocks.MockService) {
				req := newUpdateProductRequest(nil)

				service.EXPECT().
					UpdateProduct(gomock.Any(), gomock.Eq(req)).
					Times(0)
			},
			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, rec.Code)
			},
		},
		{
			name:      "product not found",
			productID: product.ID,
			req:       newUpdateProductRequest(&product),
			mock: func(service *mocks.MockService) {
				req := newUpdateProductRequest(&product)

				service.EXPECT().
					UpdateProduct(gomock.Any(), gomock.Eq(req)).
					Times(1).
					Return(responses.ProductResponse{}, fmt.Errorf("product not found"))
			},
			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder) {
				// for now when error not found just return error code 500
				require.Equal(t, http.StatusInternalServerError, rec.Code)
			},
		},
		{
			name:      "internal server error",
			productID: product.ID,
			req:       newUpdateProductRequest(&product),
			mock: func(service *mocks.MockService) {
				req := newUpdateProductRequest(&product)

				service.EXPECT().
					UpdateProduct(gomock.Any(), gomock.Eq(req)).
					Times(1).
					Return(responses.ProductResponse{}, fmt.Errorf("internal server error"))
			},
			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, rec.Code)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := mocks.NewMockService(ctrl)
			server := newGinTestServer(t, service)

			testCase.mock(service)
			url := fmt.Sprintf("/products/%d", testCase.productID)
			rec := httptest.NewRecorder()

			data, err := json.Marshal(testCase.req)
			require.NoError(t, err)

			request, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(data))
			request.Header.Set("Content-Type", "application/json")
			require.NoError(t, err)

			server.Engine.ServeHTTP(rec, request)
			testCase.checkResponse(t, rec)
		})
	}
}

func newProductTest(user responses.User) responses.ProductResponse {
	product := responses.ProductResponse{
		ID:        1,
		Name:      "Test Product",
		Price:     100,
		UserID:    user.ID,
		CreatedAt: time.Now(),
	}

	return product
}

func newUserTest() responses.User {
	return responses.User{
		ID:        1,
		Name:      "royyan",
		Email:     "roy@gmail.com",
		CreatedAt: time.Now(),
	}
}

func newProductsTest(total, page, perPage int) []responses.ProductResponse {
	lastPage := int(math.Ceil(float64(total) - float64(perPage)))
	var count int
	if page < lastPage {
		count = perPage
	} else {
		count = total - ((page - 1) * perPage)
	}

	var products []responses.ProductResponse
	for i := 0; i < count; i++ {
		product := responses.ProductResponse{
			Name: fmt.Sprintf("test product %d", count+1),
		}

		products = append(products, product)
	}

	return products
}

func newCreateProductRequest(user *responses.User, product *responses.ProductResponse) requests.CreateProductRequest {
	if user == nil && product == nil {
		return requests.CreateProductRequest{}
	}

	return requests.CreateProductRequest{
		UserID: user.ID,
		Name:   product.Name,
		Price:  product.Price,
	}
}

func newBindUriIDRequest(id int64) requests.BindUriID {
	return requests.BindUriID{
		ID: id,
	}
}

func newListProductRequest(page, perPage int32) requests.GetUserProductsRequest {
	return requests.GetUserProductsRequest{
		Page:    page,
		PerPage: perPage,
	}
}

func newUpdateProductRequest(product *responses.ProductResponse) requests.UpdateProductRequest {
	if product == nil {
		return requests.UpdateProductRequest{}
	}

	return requests.UpdateProductRequest{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}
}

func requireProductMatch(t *testing.T, body *bytes.Buffer, expectedProduct responses.ProductResponse) {
	jsonData, err := io.ReadAll(body)
	require.NoError(t, err)

	var product responses.ProductResponse
	parseJson(t, jsonData, "data.product", &product)

	// require.Equal(t, expectedProduct, product)
	require.Equal(t, expectedProduct.ID, product.ID)
	require.Equal(t, expectedProduct.Name, product.Name)
	require.Equal(t, expectedProduct.Price, product.Price)
	require.Equal(t, expectedProduct.UserID, product.UserID)
}

func requireListProductsMatch(t *testing.T, body *bytes.Buffer, products []responses.ProductResponse, pagination *responses.Pagination) {
	jsonData, err := io.ReadAll(body)
	require.NoError(t, err)

	var actualProducts []responses.ProductResponse
	parseJson(t, jsonData, "data.products", &actualProducts)

	var actualPagination *responses.Pagination
	parseJson(t, jsonData, "data.pagination", &actualPagination)

	require.Equal(t, products, actualProducts)
	require.Equal(t, pagination, actualPagination)
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
