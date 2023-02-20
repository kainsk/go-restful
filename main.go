package main

import (
	"sqlc-rest-api/config"
	"sqlc-rest-api/db/drivers"
	"sqlc-rest-api/db/postgres/repositories"
	"sqlc-rest-api/graph/generated"
	"sqlc-rest-api/services"

	gs "sqlc-rest-api/servers/gin"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()

	env, err := config.LoadEnv(".", "app")
	if err != nil {
		logger.Fatal("Failed to laod environment variables :", err)
	}

	db, err := drivers.NewPostgres(env).Connect()
	if err != nil {
		logger.Fatal("Failed to connect database :", err)
	}

	pqRepo := repositories.New()
	service := services.NewPostgresService(db, pqRepo)
	graph := handler.NewDefaultServer(
		generated.NewExecutableSchema(config.GraphConfig(service)),
	)

	ginserver, err := gs.NewGinServer(service, env, graph)
	if err != nil {
		logger.Fatal("Failed to create server :", err)
	}

	err = ginserver.Start()
	if err != nil {
		logger.Fatal("Failed to start server :", err)
	}
}
