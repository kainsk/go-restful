package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"fmt"
	"sqlc-rest-api/graph/generated"
	"sqlc-rest-api/requests"
	"sqlc-rest-api/responses"
)

// CreateProduct is the resolver for the CreateProduct field.
func (r *mutationResolver) CreateProduct(ctx context.Context, input requests.CreateProductRequest) (*responses.Product, error) {
	panic(fmt.Errorf("not implemented: CreateProduct - CreateProduct"))
}

// UpdateProduct is the resolver for the UpdateProduct field.
func (r *mutationResolver) UpdateProduct(ctx context.Context, input requests.UpdateProductRequest) (*responses.Product, error) {
	panic(fmt.Errorf("not implemented: UpdateProduct - UpdateProduct"))
}

// DeleteProduct is the resolver for the DeleteProduct field.
func (r *mutationResolver) DeleteProduct(ctx context.Context, input requests.BindUriID) (*bool, error) {
	panic(fmt.Errorf("not implemented: DeleteProduct - DeleteProduct"))
}

// User is the resolver for the user field.
func (r *productResolver) User(ctx context.Context, obj *responses.Product, input requests.BindUriID) (*responses.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// GetProduct is the resolver for the GetProduct field.
func (r *queryResolver) GetProduct(ctx context.Context, input requests.BindUriID) (*responses.Product, error) {
	panic(fmt.Errorf("not implemented: GetProduct - GetProduct"))
}

// Product returns generated.ProductResolver implementation.
func (r *Resolver) Product() generated.ProductResolver { return &productResolver{r} }

type productResolver struct{ *Resolver }
