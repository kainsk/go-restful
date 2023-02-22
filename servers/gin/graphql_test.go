package ginserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sqlc-rest-api/helpers"
	"sqlc-rest-api/mocks"
	"sqlc-rest-api/responses"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestQueryGetProduct(t *testing.T) {
	user := helpers.NewUserTest()
	product := helpers.NewProductTest(user)

	testCases := []struct {
		name          string
		query         string
		operationName string
		variables     map[string]any
		mock          func(service *mocks.MockService)
		checkResponse func(t *testing.T, rec httptest.ResponseRecorder)
	}{
		{
			name: "get product successfully",
			query: `
				query GetProduct($getProductReq: UriID!) {
					GetProduct(input: $getProductReq) {
						id
						name
						price
						user_id
						created_at
						user {
							id
							name
							email
							created_at
						}
					}
				}
			`,
			operationName: "GetProduct",
			variables: gin.H{
				"getProductReq": gin.H{
					"id": product.ID,
				},
			},
			mock: func(service *mocks.MockService) {
				getProductArg := helpers.NewBindUriIDRequestTest(product.ID)
				getUserArg := helpers.NewBindUriIDRequestTest(user.ID)

				service.EXPECT().
					GetProduct(gomock.Any(), gomock.Eq(getProductArg)).
					Times(1).
					Return(&product, nil)

				service.EXPECT().
					GetUser(gomock.Any(), gomock.Eq(getUserArg)).
					Times(1).
					Return(&user, nil)
			},
			checkResponse: func(t *testing.T, rec httptest.ResponseRecorder) {
				helpers.GraphProductMatchTest(t, "data.GetProduct", *rec.Body, product)
				helpers.GraphUserMatchTest(t, "data.GetProduct.user", *rec.Body, user)
			},
		},
		{
			name: "complexity limit product user more than 4",
			query: `
				query GetProduct($getProductReq: UriID!) {
					GetProduct(input: $getProductReq) {
						id
						name
						price
						user_id
						created_at
						user {
							id
							name
							email
							created_at
							products {
								page_info {
								  start_cursor
								}
							  }
						}
					}
				}
			`,
			operationName: "GetProduct",
			variables: gin.H{
				"getProductReq": gin.H{
					"id": product.ID,
				},
			},
			mock: func(service *mocks.MockService) {
				getProductArg := helpers.NewBindUriIDRequestTest(product.ID)
				getUserArg := helpers.NewBindUriIDRequestTest(user.ID)

				service.EXPECT().
					GetProduct(gomock.Any(), gomock.Eq(getProductArg)).
					Times(0)

				service.EXPECT().
					GetUser(gomock.Any(), gomock.Eq(getUserArg)).
					Times(0)
			},
			checkResponse: func(t *testing.T, rec httptest.ResponseRecorder) {
				helpers.GraphExpectComplexityLimit(t, "errors.0.extensions.code", *rec.Body)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := mocks.NewMockService(ctrl)
			testCase.mock(service)

			req := helpers.NewGraphQLRequestTest(testCase.operationName, testCase.query, testCase.variables)
			data, err := json.Marshal(req)
			require.NoError(t, err)

			server := newGinTestServer(t, service)
			rec := httptest.NewRecorder()
			request, err := http.NewRequest(http.MethodPost, "/graph", bytes.NewBuffer(data))
			require.NoError(t, err)
			request.Header.Set("Content-Type", "application/json")

			server.Engine.ServeHTTP(rec, request)
			testCase.checkResponse(t, *rec)
		})
	}
}

func TestQueryGetUser(t *testing.T) {
	user := helpers.NewUserTest()

	testCases := []struct {
		name          string
		query         string
		operationName string
		variables     map[string]any
		mock          func(service *mocks.MockService)
		checkResponse func(t *testing.T, rec httptest.ResponseRecorder)
	}{
		{
			name: "get user successfully",
			query: `
				query GetUser($getUserReq: UriID!) {
					GetUser(input: $getUserReq) {
						id
						name
						email
						created_at
					}
				}
			`,
			operationName: "GetUser",
			variables: gin.H{
				"getUserReq": gin.H{
					"id": user.ID,
				},
			},
			mock: func(service *mocks.MockService) {
				getUserArg := helpers.NewBindUriIDRequestTest(user.ID)

				service.EXPECT().
					GetUser(gomock.Any(), gomock.Eq(getUserArg)).
					Times(1).
					Return(&user, nil)
			},
			checkResponse: func(t *testing.T, rec httptest.ResponseRecorder) {
				helpers.GraphUserMatchTest(t, "data.GetUser", *rec.Body, user)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := mocks.NewMockService(ctrl)
			testCase.mock(service)

			req := helpers.NewGraphQLRequestTest(testCase.operationName, testCase.query, testCase.variables)
			data, err := json.Marshal(req)
			require.NoError(t, err)

			server := newGinTestServer(t, service)
			rec := httptest.NewRecorder()
			request, err := http.NewRequest(http.MethodPost, "/graph", bytes.NewBuffer(data))
			require.NoError(t, err)
			request.Header.Set("Content-Type", "application/json")

			server.Engine.ServeHTTP(rec, request)
			testCase.checkResponse(t, *rec)
		})
	}
}

func TestQueryGetUserProducts(t *testing.T) {
	user := helpers.NewUserTest()

	testCases := []struct {
		name          string
		query         string
		operationName string
		variables     map[string]any
		mock          func(service *mocks.MockService) *responses.Products
		checkResponse func(t *testing.T, rec httptest.ResponseRecorder, expectedProducts responses.Products)
	}{
		{
			name: "get user products successfully",
			query: `
				query GetUser($getUserReq: UriID!) {
					GetUser(input: $getUserReq) {
						id
						name
						email
						created_at
						products(input: {first: 5}) {
							edges {
								cursor
								node {
								  id
								  name
								  price
								  user_id
								  created_at
								}
							}
							page_info {
								start_cursor
								end_cursor
								has_next_page
							}
						}
					}
				}
			`,
			operationName: "GetUser",
			variables: gin.H{
				"getUserReq": gin.H{
					"id": user.ID,
				},
			},
			mock: func(service *mocks.MockService) *responses.Products {
				first := 5
				getUserArg := helpers.NewBindUriIDRequestTest(user.ID)
				getUserProductsArg := helpers.NewGetUserProductsRequestTest(user.ID, &first, nil)
				products := helpers.NewProductsTest(5, user.ID)

				service.EXPECT().
					GetUser(gomock.Any(), gomock.Eq(getUserArg)).
					Times(1).
					Return(&user, nil)

				service.EXPECT().
					GetUserProducts(gomock.Any(), gomock.Eq(getUserProductsArg)).
					Times(1).
					Return(products, nil)

				return products
			},
			checkResponse: func(t *testing.T, rec httptest.ResponseRecorder, expectedProducts responses.Products) {
				helpers.GraphUserMatchTest(t, "data.GetUser", *rec.Body, user)
				helpers.GraphProductsMatchTest(t, "data.GetUser.products", *rec.Body, expectedProducts)
			},
		},
		{
			name: "get user products complexity limit more than 100",
			query: `
				query GetUser($getUserReq: UriID!) {
					GetUser(input: $getUserReq) {
						id
						name
						email
						created_at
						products(input: {first: 100}) {
							edges {
								cursor
								node {
								  id
								  name
								  price
								  user_id
								  created_at
								}
							}
							page_info {
								start_cursor
								end_cursor
								has_next_page
							}
						}
					}
				}
			`,
			operationName: "GetUser",
			variables: gin.H{
				"getUserReq": gin.H{
					"id": user.ID,
				},
			},
			mock: func(service *mocks.MockService) *responses.Products {
				first := 10
				getUserArg := helpers.NewBindUriIDRequestTest(user.ID)
				getUserProductsArg := helpers.NewGetUserProductsRequestTest(user.ID, &first, nil)

				service.EXPECT().
					GetUser(gomock.Any(), gomock.Eq(getUserArg)).
					Times(0)

				service.EXPECT().
					GetUserProducts(gomock.Any(), gomock.Eq(getUserProductsArg)).
					Times(0)

				return &responses.Products{}
			},
			checkResponse: func(t *testing.T, rec httptest.ResponseRecorder, expectedProducts responses.Products) {
				helpers.GraphExpectComplexityLimit(t, "errors.0.extensions.code", *rec.Body)
			},
		},
		{
			name: "user products child complexity limit more than 12",
			query: `
				query GetUser($getUserReq: UriID!) {
					GetUser(input: $getUserReq) {
						id
						name
						email
						created_at
						products(input: {first: 5}) {
							edges {
								cursor
								node {
									id
									name
									price
									user_id
									created_at
									user {
										name
									}
								}
							}
							page_info {
								start_cursor
								end_cursor
								has_next_page
							}
						}
					}
				}
			`,
			operationName: "GetUser",
			variables: gin.H{
				"getUserReq": gin.H{
					"id": user.ID,
				},
			},
			mock: func(service *mocks.MockService) *responses.Products {
				first := 5
				getUserArg := helpers.NewBindUriIDRequestTest(user.ID)
				getUserProductsArg := helpers.NewGetUserProductsRequestTest(user.ID, &first, nil)

				service.EXPECT().
					GetUser(gomock.Any(), gomock.Eq(getUserArg)).
					Times(0)

				service.EXPECT().
					GetUserProducts(gomock.Any(), gomock.Eq(getUserProductsArg)).
					Times(0)

				return &responses.Products{}
			},
			checkResponse: func(t *testing.T, rec httptest.ResponseRecorder, expectedProducts responses.Products) {
				helpers.GraphExpectComplexityLimit(t, "errors.0.extensions.code", *rec.Body)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := mocks.NewMockService(ctrl)
			products := testCase.mock(service)

			req := helpers.NewGraphQLRequestTest(testCase.operationName, testCase.query, testCase.variables)
			data, err := json.Marshal(req)
			require.NoError(t, err)

			server := newGinTestServer(t, service)
			rec := httptest.NewRecorder()
			request, err := http.NewRequest(http.MethodPost, "/graph", bytes.NewBuffer(data))
			require.NoError(t, err)
			request.Header.Set("Content-Type", "application/json")

			server.Engine.ServeHTTP(rec, request)
			testCase.checkResponse(t, *rec, *products)
		})
	}
}
