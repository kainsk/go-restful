package ginserver

func (gs *GinServer) setupRoutes() {
	gs.Engine.POST("/products", gs.CreateProduct)
	gs.Engine.GET("/products", gs.ListProducts)
	gs.Engine.DELETE("/products/:id", gs.DeleteProduct)
	gs.Engine.GET("/products/:id", gs.GetProduct)
	gs.Engine.PUT("/products/:id", gs.UpdateProduct)
}
