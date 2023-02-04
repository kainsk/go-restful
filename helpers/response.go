package helpers

import (
	"fmt"
	"math"
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
