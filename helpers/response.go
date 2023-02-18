package helpers

import (
	"fmt"
	"math"
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

func Paginate(total, items, page, perPage int32, url string) *responses.Pagination {
	lastPage := int32(math.Ceil(float64(total) / float64(perPage)))
	firstPage := int32(1)

	var hasNextPage bool
	var linkNextPage string
	if lastPage > page {
		hasNextPage = true
		linkNextPage = url + fmt.Sprintf("?page=%d&per_page=%d", page+1, perPage)
	}

	var hasPrevPage bool
	var linkPrevPage string
	if page > firstPage {
		hasPrevPage = true
		linkPrevPage = url + fmt.Sprintf("?page=%d&per_page=%d", page-1, perPage)
	}

	return &responses.Pagination{
		CurrentPage: page,
		FirstPage:   firstPage,
		LastPage:    lastPage,
		HasNextPage: hasNextPage,
		HasPrevPage: hasPrevPage,
		From:        (page * perPage) - (perPage - 1),
		To:          ((page - 1) * perPage) + items,
		Items: responses.PaginationItems{
			Count:   items,
			Total:   total,
			PerPage: perPage,
		},
		Links: responses.PaginationLinks{
			First: url + fmt.Sprintf("?page=%d&per_page=%d", 1, perPage),
			Last:  url + fmt.Sprintf("?page=%d&per_page=%d", lastPage, perPage),
			Next:  linkNextPage,
			Prev:  linkPrevPage,
		},
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
