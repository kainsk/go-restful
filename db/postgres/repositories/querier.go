// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package repositories

import (
	"context"
)

type Querier interface {
	CreateProduct(ctx context.Context, db DBTX, productName string) (Product, error)
	DeleteProduct(ctx context.Context, db DBTX, productID int64) error
	GetProduct(ctx context.Context, db DBTX, productID int64) (Product, error)
	ListProducts(ctx context.Context, db DBTX, arg ListProductsParams) ([]Product, error)
	UpdateProduct(ctx context.Context, db DBTX, arg UpdateProductParams) (Product, error)
}

var _ Querier = (*Queries)(nil)
