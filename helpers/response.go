package helpers

import (
	"sqlc-rest-api/db/postgres/repositories"
	"sqlc-rest-api/responses"
)

func SuccessResponse(message string, data any) responses.ApiResponse {
	return responses.ApiResponse{
		Success: true,
		Message: message,
		Data:    data,
	}
}

func ProductResponse(source any) responses.Product {
	var product responses.Product
	switch p := source.(type) {
	case repositories.Product:
		product = responses.Product{
			ID:        p.ID,
			Name:      p.Name,
			Price:     p.Price,
			UserID:    p.UserID,
			CreatedAt: p.CreatedAt.Time,
		}
	default:
		panic("incompatible source")
	}

	return product
}

func UserResponse(source any) responses.User {
	var user responses.User
	switch u := source.(type) {
	case repositories.User:
		user = responses.User{
			ID:        u.ID,
			Name:      u.Name,
			Email:     u.Email,
			CreatedAt: u.CreatedAt.Time,
		}
	default:
		panic("incompatible source")
	}

	return user
}
