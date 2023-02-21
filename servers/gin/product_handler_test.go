package ginserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sqlc-rest-api/helpers"
	"sqlc-rest-api/requests"
	"testing"

	"sqlc-rest-api/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestCreateProduct(t *testing.T) {
	user := helpers.NewUserTest()
	product := helpers.NewProductTest(user)

	testCases := []struct {
		name          string
		req           requests.CreateProductRequest
		mock          func(service *mocks.MockService)
		checkResponse func(t *testing.T, rec *httptest.ResponseRecorder)
	}{
		{
			name: "product created successfully",
			req:  helpers.NewCreateProductRequestTest(&user, &product),
			mock: func(service *mocks.MockService) {
				req := helpers.NewCreateProductRequestTest(&user, &product)
				service.EXPECT().
					CreateProduct(gomock.Any(), gomock.Eq(req)).
					Times(1).
					Return(&product, nil)
			},
			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, rec.Code)
				helpers.RequireProductMatchTest(t, rec.Body, product)
			},
		},
		{
			name: "validation error product name not given",
			req:  helpers.NewCreateProductRequestTest(nil, nil),
			mock: func(service *mocks.MockService) {
				req := helpers.NewCreateProductRequestTest(nil, nil)
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
			req:  helpers.NewCreateProductRequestTest(&user, &product),
			mock: func(service *mocks.MockService) {
				req := helpers.NewCreateProductRequestTest(&user, &product)
				service.EXPECT().
					CreateProduct(gomock.Any(), gomock.Eq(req)).
					Times(1).
					Return(nil, fmt.Errorf("internal server error"))
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
	user := helpers.NewUserTest()
	product := helpers.NewProductTest(user)

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
				req := helpers.NewBindUriIDRequestTest(product.ID)
				service.EXPECT().
					DeleteProduct(gomock.Any(), gomock.Eq(req)).
					Times(1).
					Return(helpers.NewProductDeletedTest(product.ID), nil)
			},
			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, rec.Code)
			},
		},
		{
			name:      "validation error because given id lower than one",
			productID: 0,
			mock: func(service *mocks.MockService) {
				req := helpers.NewBindUriIDRequestTest(0)
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
				req := helpers.NewBindUriIDRequestTest(product.ID)
				service.EXPECT().
					DeleteProduct(gomock.Any(), gomock.Eq(req)).
					Times(1).
					Return(nil, fmt.Errorf("product not found"))
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
				req := helpers.NewBindUriIDRequestTest(product.ID)
				service.EXPECT().
					DeleteProduct(gomock.Any(), gomock.Eq(req)).
					Times(1).
					Return(nil, fmt.Errorf("internal server error"))
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
	user := helpers.NewUserTest()
	product := helpers.NewProductTest(user)

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
				req := helpers.NewBindUriIDRequestTest(product.ID)
				service.EXPECT().
					GetProduct(gomock.Any(), gomock.Eq(req)).
					Times(1).
					Return(&product, nil)
			},
			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, rec.Code)
				helpers.RequireProductMatchTest(t, rec.Body, product)
			},
		},
		{
			name:      "validation error because given id lower than one",
			productID: 0,
			mock: func(service *mocks.MockService) {
				req := helpers.NewBindUriIDRequestTest(0)
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
				req := helpers.NewBindUriIDRequestTest(product.ID)
				service.EXPECT().
					GetProduct(gomock.Any(), gomock.Eq(req)).
					Times(1).
					Return(nil, fmt.Errorf("product not found"))
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
				req := helpers.NewBindUriIDRequestTest(product.ID)
				service.EXPECT().
					GetProduct(gomock.Any(), gomock.Eq(req)).
					Times(1).
					Return(nil, fmt.Errorf("internal server error"))
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

func TestUpdateProduct(t *testing.T) {
	user := helpers.NewUserTest()
	product := helpers.NewProductTest(user)

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
			req:       helpers.NewUpdateProductRequestTest(&product),
			mock: func(service *mocks.MockService) {
				req := helpers.NewUpdateProductRequestTest(&product)

				service.EXPECT().
					UpdateProduct(gomock.Any(), gomock.Eq(req)).
					Times(1).
					Return(&product, nil)
			},
			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, rec.Code)
				helpers.RequireProductMatchTest(t, rec.Body, product)
			},
		},
		{
			name:      "validation error product id given lower than 1",
			productID: 0,
			req:       helpers.NewUpdateProductRequestTest(nil),
			mock: func(service *mocks.MockService) {
				req := helpers.NewUpdateProductRequestTest(nil)

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
			req:       helpers.NewUpdateProductRequestTest(nil),
			mock: func(service *mocks.MockService) {
				req := helpers.NewUpdateProductRequestTest(nil)

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
			req:       helpers.NewUpdateProductRequestTest(&product),
			mock: func(service *mocks.MockService) {
				req := helpers.NewUpdateProductRequestTest(&product)

				service.EXPECT().
					UpdateProduct(gomock.Any(), gomock.Eq(req)).
					Times(1).
					Return(nil, fmt.Errorf("product not found"))
			},
			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder) {
				// for now when error not found just return error code 500
				require.Equal(t, http.StatusInternalServerError, rec.Code)
			},
		},
		{
			name:      "internal server error",
			productID: product.ID,
			req:       helpers.NewUpdateProductRequestTest(&product),
			mock: func(service *mocks.MockService) {
				req := helpers.NewUpdateProductRequestTest(&product)

				service.EXPECT().
					UpdateProduct(gomock.Any(), gomock.Eq(req)).
					Times(1).
					Return(nil, fmt.Errorf("internal server error"))
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
