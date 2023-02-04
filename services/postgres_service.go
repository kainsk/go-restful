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

func NewPostgresService(db *sql.DB) *PostgresService {
	return &PostgresService{
		Repo: repositories.New(),
		DB:   db,
	}
}

func (pq *PostgresService) CreateProduct(ctx context.Context, req requests.CreateProductRequest) (responses.ProductResponse, error) {
	prod, err := pq.Repo.CreateProduct(ctx, pq.DB, req.Name)
	if err != nil {
		return responses.ProductResponse{}, err
	}

	return responses.ProductResponse{
		ID:   prod.ID,
		Name: prod.Name,
	}, nil
}

func (pq *PostgresService) DeleteProduct(ctx context.Context, req requests.BindUriID) error {
	prod, err := pq.Repo.GetProduct(ctx, pq.DB, req.ID)
	if err != nil {
		return fmt.Errorf("product with id %d not found", req.ID)
	}

	return pq.Repo.DeleteProduct(ctx, pq.DB, prod.ID)
}

func (pq *PostgresService) GetProduct(ctx context.Context, req requests.BindUriID) (responses.ProductResponse, error) {
	prod, err := pq.Repo.GetProduct(ctx, pq.DB, req.ID)
	if err != nil {
		return responses.ProductResponse{}, fmt.Errorf("product with id %d not found", req.ID)
	}

	return responses.ProductResponse{
		ID:   prod.ID,
		Name: prod.Name,
	}, nil
}

func (pq *PostgresService) ListProducts(ctx context.Context, req requests.ListProductsRequest) ([]responses.ProductResponse, *responses.Pagination, error) {
	countProds, err := pq.Repo.CountProducts(ctx, pq.DB)
	if err != nil {
		return nil, nil, err
	}

	arg := repositories.ListProductsParams{
		Offset: (req.Page - 1) * req.PerPage,
		Limit:  req.PerPage,
	}

	prods, err := pq.Repo.ListProducts(ctx, pq.DB, arg)
	if err != nil {
		return nil, nil, err
	}

	var products []responses.ProductResponse
	for _, prod := range prods {
		p := responses.ProductResponse{
			ID:   prod.ID,
			Name: prod.Name,
		}

		products = append(products, p)
	}

	pagination := helpers.Paginate(
		int32(countProds),
		int32(len(products)),
		req.Page,
		req.PerPage,
		"/products",
	)

	return products, pagination, nil
}

func (pq *PostgresService) UpdateProduct(ctx context.Context, req requests.UpdateProductRequest) (responses.ProductResponse, error) {
	prod, err := pq.Repo.GetProduct(ctx, pq.DB, req.ID)
	if err != nil {
		return responses.ProductResponse{}, fmt.Errorf("product with id %d not found", req.ID)
	}

	arg := repositories.UpdateProductParams{
		ProductID:      prod.ID,
		NewProductName: req.Name,
	}

	updated, err := pq.Repo.UpdateProduct(ctx, pq.DB, arg)
	if err != nil {
		return responses.ProductResponse{}, err
	}

	return responses.ProductResponse{
		ID:   updated.ID,
		Name: updated.Name,
	}, nil
}
