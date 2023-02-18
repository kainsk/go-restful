package requests

type CreateProductRequest struct {
	UserID int64  `json:"user_id" binding:"required,min=1"`
	Price  int64  `json:"price" binding:"required,min=1"`
	Name   string `json:"name" binding:"required"`
}

type BindUriID struct {
	ID int64 `json:"id" binding:"required,min=1" uri:"id"`
}

type GetUserProductsRequest struct {
	UserID  int64
	Page    int32 `form:"page,default=1" binding:"number,min=1"`
	PerPage int32 `form:"per_page,default=25" binding:"number,min=1,max=50"`
}

type UpdateProductRequest struct {
	ID    int64
	Name  string `json:"name" binding:"required"`
	Price int64  `json:"price" binding:"required"`
}
