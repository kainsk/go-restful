package requests

type CreateProductRequest struct {
	UserID int64  `json:"user_id" binding:"required,min=1"`
	Price  int64  `json:"price" binding:"required,min=1"`
	Name   string `json:"name" binding:"required"`
}

type BindUriID struct {
	ID int64 `json:"id" binding:"required,min=1" uri:"id"`
}

type UpdateProductRequest struct {
	ID    int64
	Name  string `json:"name" binding:"required"`
	Price int64  `json:"price" binding:"required"`
}
