package ginserver

import (
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

func (gs *GinServer) setupRoutes() {
	gs.Engine.POST("/products", gs.CreateProduct)
	gs.Engine.DELETE("/products/:id", gs.DeleteProduct)
	gs.Engine.GET("/products/:id", gs.GetProduct)
	gs.Engine.PUT("/products/:id", gs.UpdateProduct)

	gs.Engine.POST("/users", gs.CreateUser)
	gs.Engine.GET("/users/:id", gs.GetUser)
	gs.Engine.GET("/user/:id/products", gs.GetUserProducts)

	gs.Engine.GET("/playground", gs.graphPlayground())
	gs.Engine.POST("/graph", gs.graphQuery())
}

func (gs *GinServer) graphPlayground() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/graph")
	return func(ctx *gin.Context) {
		h.ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func (gs *GinServer) graphQuery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		gs.Graph.ServeHTTP(ctx.Writer, ctx.Request)
	}
}
