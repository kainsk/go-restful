package repositories

import (
	"database/sql"
	"fmt"
	"os"
	"sqlc-rest-api/config"
	"testing"

	"github.com/sirupsen/logrus"
)

var testDB *sql.DB
var testRepo *Queries

func TestMain(m *testing.M) {
	var err error
	logger := logrus.New()

	env, err := config.LoadEnv("../../../", "app")
	if err != nil {
		logger.Error(err)
	}

	testDB, err = sql.Open(
		"postgres",
		fmt.Sprintf(
			"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
			env.DBUsername,
			env.DBPassword,
			env.DBHost,
			env.DBPort,
			env.DBName,
		),
	)

	if err != nil {
		logger.Error(err)
	}

	testRepo = New()
	os.Exit(m.Run())
}
