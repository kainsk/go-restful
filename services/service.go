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
	ListProducts(ctx context.Context, req requests.ListProductsRequest) ([]responses.ProductResponse, *responses.Pagination, error)
	UpdateProduct(ctx context.Context, req requests.UpdateProductRequest) (responses.ProductResponse, error)
}
