package ginserver

import (
	"sqlc-rest-api/helpers"
	"sqlc-rest-api/requests"

	"github.com/gin-gonic/gin"
)

func (gs *GinServer) CreateProduct(c *gin.Context) {
	var req requests.CreateProductRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	product, err := gs.Service.CreateProduct(c, req)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	data := gin.H{
		"product": product,
	}

	resp := helpers.SuccessResponse("product created successfully", data)
	c.JSON(201, resp)
}

func (gs *GinServer) DeleteProduct(c *gin.Context) {
	var req requests.BindUriID
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := gs.Service.DeleteProduct(c, req)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	resp := helpers.SuccessResponse("product deleted successfully", nil)
	c.JSON(200, resp)
}

func (gs *GinServer) GetProduct(c *gin.Context) {
	var req requests.BindUriID
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	product, err := gs.Service.GetProduct(c, req)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	data := gin.H{
		"product": product,
	}

	resp := helpers.SuccessResponse("get one product successfully", data)
	c.JSON(200, resp)
}

func (gs *GinServer) ListProducts(c *gin.Context) {
	var req requests.ListProductsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	prods, pagination, err := gs.Service.ListProducts(c, req)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	data := gin.H{
		"products":   prods,
		"pagination": pagination,
	}

	resp := helpers.SuccessResponse("list products successfully", data)
	c.JSON(200, resp)
}

func (gs *GinServer) UpdateProduct(c *gin.Context) {
	var req requests.UpdateProductRequest
	var uri requests.BindUriID

	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	req.ID = uri.ID
	prod, err := gs.Service.UpdateProduct(c, req)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	data := gin.H{
		"products": prod,
	}

	resp := helpers.SuccessResponse("update product successfully", data)
	c.JSON(200, resp)
}
