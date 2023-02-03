package main

import (
	"sqlc-rest-api/config"
	"sqlc-rest-api/db/drivers"
	"sqlc-rest-api/services"

	gs "sqlc-rest-api/servers/gin"

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

	service := services.NewPostgresService(db)
	ginserver, err := gs.NewGinServer(service, env)
	if err != nil {
		logger.Fatal("Failed to create server :", err)
	}

	err = ginserver.Start()
	if err != nil {
		logger.Fatal("Failed to start server :", err)
	}
}
