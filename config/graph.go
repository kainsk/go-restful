package config

import (
	"sqlc-rest-api/graph/generated"
	"sqlc-rest-api/graph/resolvers"
	"sqlc-rest-api/services"
)

func GraphConfig(service services.Service) generated.Config {
	resolver := resolvers.NewResolver(service)
	config := generated.Config{
		Resolvers: resolver,
	}

	return config
}
