package ginserver

import (
	"os"
	"sqlc-rest-api/config"
	"sqlc-rest-api/services"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func newGinTestServer(t *testing.T, service services.Service) *GinServer {
	env := config.Environment{}
	server, err := NewGinServer(service, env)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
