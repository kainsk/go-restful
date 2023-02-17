package services

import (
	"context"
	"sqlc-rest-api/requests"
	"sqlc-rest-api/responses"
)

type Service interface {
	CreateProduct(ctx context.Context, req requests.CreateProductRequest) (responses.ProductResponse, error)
	DeleteProduct(ctx context.Context, req requests.BindUriID) error
	GetProduct(ctx context.Context, req requests.BindUriID) (responses.ProductResponse, error)
	GetUserProducts(ctx context.Context, req requests.GetUserProductsRequest) ([]responses.ProductResponse, *responses.Pagination, error)
	UpdateProduct(ctx context.Context, req requests.UpdateProductRequest) (responses.ProductResponse, error)
	CreateUser(ctx context.Context, req requests.CreateUserRequest) (responses.User, error)
	GetUser(ctx context.Context, req requests.BindUriID) (responses.User, error)
}
