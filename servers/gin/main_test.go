package ginserver

import (
	"os"
	"sqlc-rest-api/config"
	graphconfig "sqlc-rest-api/graph/config"
	"sqlc-rest-api/graph/generated"
	"sqlc-rest-api/services"
	"testing"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func newGinTestServer(t *testing.T, service services.Service) *GinServer {
	graph := handler.NewDefaultServer(generated.NewExecutableSchema(
		graphconfig.GraphConfig(service),
	))

	env := config.Environment{ComplexityLimit: 100}
	graph.Use(extension.FixedComplexityLimit(env.ComplexityLimit))
	server, err := NewGinServer(service, env, graph)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
