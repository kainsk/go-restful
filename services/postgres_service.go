package services

import (
	"context"
	"database/sql"
	"fmt"
	"sqlc-rest-api/db/postgres/repositories"
	"sqlc-rest-api/helpers"
	"sqlc-rest-api/requests"
	"sqlc-rest-api/responses"
)

type PostgresService struct {
	Repo repositories.Querier
	DB   *sql.DB
}

func NewPostgresService(db *sql.DB, pqrepo repositories.Querier) *PostgresService {
	return &PostgresService{
		Repo: pqrepo,
		DB:   db,
	}
}

func (pq *PostgresService) CreateProduct(ctx context.Context, req requests.CreateProductRequest) (responses.Product, error) {
	arg := repositories.CreateProductParams{
		UserID: req.UserID,
		Name:   req.Name,
		Price:  req.Price,
	}

	prod, err := pq.Repo.CreateProduct(ctx, pq.DB, arg)
	if err != nil {
		return responses.Product{}, err
	}

	return responses.Product{
		ID:        prod.ID,
		Name:      prod.Name,
		Price:     prod.Price,
		UserID:    prod.UserID,
		CreatedAt: prod.CreatedAt.Time,
	}, nil
}

func (pq *PostgresService) DeleteProduct(ctx context.Context, req requests.BindUriID) error {
	prod, err := pq.Repo.GetProduct(ctx, pq.DB, req.ID)
	if err != nil {
		return fmt.Errorf("product with id %d not found", req.ID)
	}

	return pq.Repo.DeleteProduct(ctx, pq.DB, prod.ID)
}

func (pq *PostgresService) GetProduct(ctx context.Context, req requests.BindUriID) (responses.Product, error) {
	prod, err := pq.Repo.GetProduct(ctx, pq.DB, req.ID)
	if err != nil {
		return responses.Product{}, fmt.Errorf("product with id %d not found", req.ID)
	}

	return responses.Product{
		ID:        prod.ID,
		Name:      prod.Name,
		Price:     prod.Price,
		UserID:    prod.UserID,
		CreatedAt: prod.CreatedAt.Time,
	}, nil
}

func (pq *PostgresService) GetUserProducts(ctx context.Context, req requests.GetUserProductsRequest) ([]responses.Product, *responses.Pagination, error) {
	countProds, err := pq.Repo.CountProductsByUserID(ctx, pq.DB, req.UserID)
	if err != nil {
		return nil, nil, err
	}

	arg := repositories.ListProductsParams{
		Offset: (req.Page - 1) * req.PerPage,
		Limit:  req.PerPage,
		UserID: req.UserID,
	}

	prods, err := pq.Repo.ListProducts(ctx, pq.DB, arg)
	if err != nil {
		return nil, nil, err
	}

	var products []responses.Product
	for _, prod := range prods {
		p := responses.Product{
			ID:        prod.ID,
			Name:      prod.Name,
			Price:     prod.Price,
			UserID:    prod.UserID,
			CreatedAt: prod.CreatedAt.Time,
		}

		products = append(products, p)
	}

	// TODO: using cursor based pagination
	pagination := helpers.Paginate(
		int32(countProds),
		int32(len(products)),
		req.Page,
		req.PerPage,
		"/products",
	)

	return products, pagination, nil
}

func (pq *PostgresService) UpdateProduct(ctx context.Context, req requests.UpdateProductRequest) (responses.Product, error) {
	prod, err := pq.Repo.GetProduct(ctx, pq.DB, req.ID)
	if err != nil {
		return responses.Product{}, fmt.Errorf("product with id %d not found", req.ID)
	}

	arg := repositories.UpdateProductParams{
		ID:    prod.ID,
		Name:  req.Name,
		Price: req.Price,
	}

	updated, err := pq.Repo.UpdateProduct(ctx, pq.DB, arg)
	if err != nil {
		return responses.Product{}, err
	}

	return responses.Product{
		ID:        updated.ID,
		Name:      updated.Name,
		Price:     updated.Price,
		UserID:    updated.UserID,
		CreatedAt: updated.CreatedAt.Time,
	}, nil
}

func (pq *PostgresService) CreateUser(ctx context.Context, req requests.CreateUserRequest) (responses.User, error) {
	arg := repositories.CreateUserParams{
		Name:  req.Name,
		Email: req.Email,
	}

	user, err := pq.Repo.CreateUser(ctx, pq.DB, arg)
	if err != nil {
		return responses.User{}, err
	}

	return responses.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Time,
	}, nil
}

func (pq *PostgresService) GetUser(ctx context.Context, req requests.BindUriID) (responses.User, error) {
	user, err := pq.Repo.GetUser(ctx, pq.DB, req.ID)
	if err != nil {
		return responses.User{}, err
	}

	return responses.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Time,
	}, nil
}
