package ginserver

import (
	"sqlc-rest-api/helpers"
	"sqlc-rest-api/requests"

	"github.com/gin-gonic/gin"
)

func (gs *GinServer) CreateProduct(c *gin.Context) {
	var req requests.CreateProductRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, gin.H{
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
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	deletedProduct, err := gs.Service.DeleteProduct(c, req)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	data := gin.H{
		"product": deletedProduct,
	}

	resp := helpers.SuccessResponse("product deleted successfully", data)
	c.JSON(200, resp)
}

func (gs *GinServer) GetProduct(c *gin.Context) {
	var req requests.BindUriID
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(400, gin.H{
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

func (gs *GinServer) GetUserProducts(c *gin.Context) {
	var req requests.GetUserProductsRequest
	var uri requests.BindUriID

	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	req.UserID = uri.ID
	user, err := gs.Service.GetUserProducts(c, req)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	data := gin.H{
		"user": user,
	}

	resp := helpers.SuccessResponse("list user products successfully", data)
	c.JSON(200, resp)
}

func (gs *GinServer) UpdateProduct(c *gin.Context) {
	var req requests.UpdateProductRequest
	var uri requests.BindUriID

	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, gin.H{
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
		"product": prod,
	}

	resp := helpers.SuccessResponse("update product successfully", data)
	c.JSON(200, resp)
}
