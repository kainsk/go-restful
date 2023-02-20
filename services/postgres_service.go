package services

import (
	"context"
	"database/sql"
	"encoding/base64"
	"fmt"
	"sqlc-rest-api/db/postgres/repositories"
	"sqlc-rest-api/helpers"
	"sqlc-rest-api/requests"
	"sqlc-rest-api/responses"
	"time"
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

func (pq *PostgresService) CreateProduct(ctx context.Context, req requests.CreateProductRequest) (*responses.Product, error) {
	arg := repositories.CreateProductParams{
		UserID: req.UserID,
		Name:   req.Name,
		Price:  req.Price,
	}

	prod, err := pq.Repo.CreateProduct(ctx, pq.DB, arg)
	if err != nil {
		return &responses.Product{}, err
	}

	return helpers.ProductResponse(prod), nil
}

func (pq *PostgresService) DeleteProduct(ctx context.Context, req requests.BindUriID) error {
	prod, err := pq.Repo.GetProduct(ctx, pq.DB, req.ID)
	if err != nil {
		return fmt.Errorf("product with id %d not found", req.ID)
	}

	return pq.Repo.DeleteProduct(ctx, pq.DB, prod.ID)
}

func (pq *PostgresService) GetProduct(ctx context.Context, req requests.BindUriID) (*responses.Product, error) {
	prod, err := pq.Repo.GetProduct(ctx, pq.DB, req.ID)
	if err != nil {
		return &responses.Product{}, fmt.Errorf("product with id %d not found", req.ID)
	}

	return helpers.ProductResponse(prod), nil
}

func (pq *PostgresService) GetUserProducts(ctx context.Context, req requests.GetUserProductsRequest) (*responses.User, error) {
	u, err := pq.Repo.GetUser(ctx, pq.DB, req.UserID)
	if err != nil {
		return &responses.User{}, fmt.Errorf("user with id %d not found", req.UserID)
	}

	after := time.Now()
	if req.After != nil {
		after = helpers.DecodeCursor(*req.After)
	}

	arg := repositories.GetUserProductsParams{
		UserID: req.UserID,
		After:  sql.NullTime{Valid: true, Time: after},
		First:  int32(*req.First),
	}

	results, err := pq.Repo.GetUserProducts(ctx, pq.DB, arg)
	if err != nil {
		return &responses.User{}, err
	}

	if len(results) < 1 {
		user := helpers.UserResponse(u)
		user.Products = &responses.Products{}
		return user, nil
	}

	var hasNextPage bool
	edges := make([]*responses.ProductEdge, len(results))
	for i, result := range results {
		hasNextPage = result.Exists
		ds := result.CreatedAt.Time.String()
		edges[i] = &responses.ProductEdge{
			Cursor: base64.StdEncoding.EncodeToString([]byte(ds)),
			Node: &responses.Product{
				ID:        result.ID,
				Name:      result.Name,
				Price:     result.Price,
				UserID:    result.UserID,
				CreatedAt: result.CreatedAt.Time,
			},
		}
	}

	pageInfo := responses.PageInfo{
		StartCursor: base64.StdEncoding.EncodeToString([]byte(edges[0].Node.CreatedAt.String())),
		EndCursor:   base64.StdEncoding.EncodeToString([]byte(edges[len(edges)-1].Node.CreatedAt.String())),
		HasNextPage: hasNextPage,
	}

	products := responses.Products{
		Edges:    edges,
		PageInfo: &pageInfo,
	}

	user := helpers.UserResponse(u)
	user.Products = &products

	return user, nil
}

func (pq *PostgresService) UpdateProduct(ctx context.Context, req requests.UpdateProductRequest) (*responses.Product, error) {
	prod, err := pq.Repo.GetProduct(ctx, pq.DB, req.ID)
	if err != nil {
		return &responses.Product{}, fmt.Errorf("product with id %d not found", req.ID)
	}

	arg := repositories.UpdateProductParams{
		ID:    prod.ID,
		Name:  req.Name,
		Price: req.Price,
	}

	updated, err := pq.Repo.UpdateProduct(ctx, pq.DB, arg)
	if err != nil {
		return &responses.Product{}, err
	}

	return helpers.ProductResponse(updated), nil
}

func (pq *PostgresService) CreateUser(ctx context.Context, req requests.CreateUserRequest) (*responses.User, error) {
	arg := repositories.CreateUserParams{
		Name:  req.Name,
		Email: req.Email,
	}

	user, err := pq.Repo.CreateUser(ctx, pq.DB, arg)
	if err != nil {
		return &responses.User{}, err
	}

	return helpers.UserResponse(user), nil
}

func (pq *PostgresService) GetUser(ctx context.Context, req requests.BindUriID) (*responses.User, error) {
	user, err := pq.Repo.GetUser(ctx, pq.DB, req.ID)
	if err != nil {
		return &responses.User{}, err
	}

	return helpers.UserResponse(user), nil
}
