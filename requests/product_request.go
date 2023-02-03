package requests

type CreateProductRequest struct {
	Name string `json:"name" binding:"required"`
}

type BindUriID struct {
	ID int64 `json:"id" binding:"required,min=1" uri:"id"`
}

type ListProductsRequest struct {
	Offset int32 `json:"offset" binding:"number,min=0"`
	Limit  int32 `json:"limit" binding:"number,min=1"`
}

type UpdateProductRequest struct {
	ID   int64
	Name string `json:"name" binding:"required"`
}
