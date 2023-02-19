package ginserver

func (gs *GinServer) setupRoutes() {
	gs.Engine.POST("/products", gs.CreateProduct)
	gs.Engine.DELETE("/products/:id", gs.DeleteProduct)
	gs.Engine.GET("/products/:id", gs.GetProduct)
	gs.Engine.PUT("/products/:id", gs.UpdateProduct)

	gs.Engine.POST("/users", gs.CreateUser)
	gs.Engine.GET("/users/:id", gs.GetUser)
	gs.Engine.GET("/user/:id/products", gs.GetUserProducts)
}
