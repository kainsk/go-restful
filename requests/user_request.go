package requests

type CreateUserRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type GetUserProductsRequest struct {
	UserID int64   `json:"user_id" uri:"id"`
	First  *int    `json:"first" form:"first,default=5" binding:"number,min=1"`
	After  *string `json:"after" form:"after"`
}
