package responses

import "time"

type ProductResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Price     int64     `json:"price"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	User      *User     `json:"user,omitempty"`
}
