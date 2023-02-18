package services

import (
	"context"
	"sqlc-rest-api/requests"
	"sqlc-rest-api/responses"
)

type Service interface {
	CreateProduct(ctx context.Context, req requests.CreateProductRequest) (responses.Product, error)
	DeleteProduct(ctx context.Context, req requests.BindUriID) error
	GetProduct(ctx context.Context, req requests.BindUriID) (responses.Product, error)
	GetUserProducts(ctx context.Context, req requests.GetUserProductsRequest) ([]responses.Product, *responses.Pagination, error)
	UpdateProduct(ctx context.Context, req requests.UpdateProductRequest) (responses.Product, error)
	CreateUser(ctx context.Context, req requests.CreateUserRequest) (responses.User, error)
	GetUser(ctx context.Context, req requests.BindUriID) (responses.User, error)
}
