DB_URL=postgres://postgres:ligmaballs@localhost:5432/sqlc_project?sslmode=disable

pqcreate:
	migrate create -ext sql  -dir db/postgres/schemas -seq $(name)

pqup:
	migrate -database "$(DB_URL)" -path db/postgres/schemas up $(n)

pqdown:
	migrate -database "$(DB_URL)" -path db/postgres/schemas down $(n)

pqclean:
	migrate -database "$(DB_URL)" -path db/postgres/schemas drop -f

pqmock:
	mockgen -package mocks -destination mocks/service_mock.go -source services/service.go Service

gqlgen:
	go run github.com/99designs/gqlgen generate