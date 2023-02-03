package ginserver

import (
	"fmt"
	"sqlc-rest-api/config"
	"sqlc-rest-api/services"

	"github.com/gin-gonic/gin"
)

type GinServer struct {
	Service services.Service
	Engine  *gin.Engine
	Env     config.Environment
}

func NewGinServer(service services.Service, env config.Environment) (*GinServer, error) {
	gs := &GinServer{
		Service: service,
		Env:     env,
		Engine:  gin.Default(),
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
