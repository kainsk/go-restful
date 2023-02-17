// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0

package repositories

import (
	"context"
)

type Querier interface {
	CountProductsByUserID(ctx context.Context, db DBTX, userID int64) (int64, error)
	CreateProduct(ctx context.Context, db DBTX, arg CreateProductParams) (Product, error)
	CreateUser(ctx context.Context, db DBTX, arg CreateUserParams) (User, error)
	DeleteProduct(ctx context.Context, db DBTX, id int64) error
	GetBatchUsers(ctx context.Context, db DBTX, ids []int64) ([]User, error)
	GetProduct(ctx context.Context, db DBTX, id int64) (Product, error)
	GetUser(ctx context.Context, db DBTX, id int64) (User, error)
	ListProducts(ctx context.Context, db DBTX, arg ListProductsParams) ([]Product, error)
	UpdateProduct(ctx context.Context, db DBTX, arg UpdateProductParams) (Product, error)
}

var _ Querier = (*Queries)(nil)
