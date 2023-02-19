package responses

import "time"

type Product struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Price     int64     `json:"price"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	User      *User     `json:"user,omitempty"`
}

type Products struct {
	Edges    []*ProductEdge `json:"edges"`
	PageInfo *PageInfo      `json:"page_info"`
}

type ProductEdge struct {
	Cursor string   `json:"cursor"`
	Node   *Product `json:"node"`
}
