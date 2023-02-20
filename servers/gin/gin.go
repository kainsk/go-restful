package ginserver

import (
	"fmt"
	"sqlc-rest-api/config"
	"sqlc-rest-api/services"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
)

type GinServer struct {
	Service services.Service
	Engine  *gin.Engine
	Graph   *handler.Server
	Env     config.Environment
}

func NewGinServer(service services.Service, env config.Environment, graph *handler.Server) (*GinServer, error) {
	gs := &GinServer{
		Service: service,
		Env:     env,
		Engine:  gin.Default(),
		Graph:   graph,
	}

	gs.setupRoutes()

	return gs, nil
}

func (gs *GinServer) Start() error {
	return gs.Engine.Run(
		fmt.Sprintf(
			"%s:%s",
			gs.Env.ServerHost,
			gs.Env.ServerPort,
		),
	)
}
