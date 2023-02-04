package responses

type PaginationItems struct {
	Count   int32 `json:"count"`
	Total   int32 `json:"total"`
	PerPage int32 `json:"per_page"`
}

type PaginationLinks struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Next  string `json:"next"`
	Prev  string `json:"prev"`
}

type Pagination struct {
	CurrentPage int32           `json:"current_page"`
	FirstPage   int32           `json:"first_page"`
	LastPage    int32           `json:"last_page"`
	HasNextPage bool            `json:"has_next_page"`
	HasPrevPage bool            `json:"has_prev_page"`
	From        int32           `json:"from"`
	To          int32           `json:"to"`
	Items       PaginationItems `json:"items"`
	Links       PaginationLinks `json:"links"`
}
