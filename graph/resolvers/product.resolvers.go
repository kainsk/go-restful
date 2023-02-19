package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"sqlc-rest-api/graph/generated"
	"sqlc-rest-api/requests"
	"sqlc-rest-api/responses"
)

// CreateProduct is the resolver for the CreateProduct field.
func (r *mutationResolver) CreateProduct(ctx context.Context, input requests.CreateProductRequest) (responses.Product, error) {
	return r.Service.CreateProduct(ctx, input)
}

// UpdateProduct is the resolver for the UpdateProduct field.
func (r *mutationResolver) UpdateProduct(ctx context.Context, input requests.UpdateProductRequest) (responses.Product, error) {
	return r.Service.UpdateProduct(ctx, input)
}

// DeleteProduct is the resolver for the DeleteProduct field.
func (r *mutationResolver) DeleteProduct(ctx context.Context, input requests.BindUriID) (*bool, error) {
	err := r.Service.DeleteProduct(ctx, input)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// User is the resolver for the user field.
func (r *productResolver) User(ctx context.Context, obj *responses.Product, input *requests.BindUriID) (responses.User, error) {
	arg := requests.BindUriID{
		ID: obj.UserID,
	}

	return r.Service.GetUser(ctx, arg)
}

// GetProduct is the resolver for the GetProduct field.
func (r *queryResolver) GetProduct(ctx context.Context, input requests.BindUriID) (responses.Product, error) {
	return r.Service.GetProduct(ctx, input)
}

// Product returns generated.ProductResolver implementation.
func (r *Resolver) Product() generated.ProductResolver { return &productResolver{r} }

type productResolver struct{ *Resolver }