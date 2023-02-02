package drivers

import (
	"database/sql"
	"fmt"
	"sqlc-rest-api/config"

	_ "github.com/lib/pq"
)

type Postgres struct {
	env config.Environment
}

func NewPostgres(env config.Environment) *Postgres {
	return &Postgres{
		env: env,
	}
}

func (pq *Postgres) Connect() (*sql.DB, error) {
	db, err := sql.Open(
		"postgres",
		fmt.Sprintf(
			"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
			pq.env.DBUsername,
			pq.env.DBPassword,
			pq.env.DBHost,
			pq.env.DBPort,
			pq.env.DBName,
		),
	)

	if err != nil {
		return nil, err
	}

	return db, err
}
