package ginserver

import (
	"os"
	"sqlc-rest-api/config"
	"sqlc-rest-api/graph/generated"
	"sqlc-rest-api/graph/resolvers"
	"sqlc-rest-api/services"
	"testing"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func newGinTestServer(t *testing.T, service services.Service) *GinServer {
	resolver := resolvers.NewResolver(service)
	c := generated.Config{Resolvers: resolver}
	graph := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	env := config.Environment{}
	server, err := NewGinServer(service, env, graph)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
