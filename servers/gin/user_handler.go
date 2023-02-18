package ginserver

import (
	"sqlc-rest-api/helpers"
	"sqlc-rest-api/requests"

	"github.com/gin-gonic/gin"
)

func (gs *GinServer) CreateUser(c *gin.Context) {
	var req requests.CreateUserRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	user, err := gs.Service.CreateUser(c, req)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	data := gin.H{
		"user": user,
	}

	resp := helpers.SuccessResponse("user created successfully", data)
	c.JSON(201, resp)
}

func (gs *GinServer) GetUser(c *gin.Context) {
	var uri requests.BindUriID
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	user, err := gs.Service.GetUser(c, uri)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	data := gin.H{
		"user": user,
	}

	resp := helpers.SuccessResponse("get user successfully", data)
	c.JSON(200, resp)
}
