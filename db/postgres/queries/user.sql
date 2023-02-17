-- name: CreateUser :one
INSERT INTO users(
    name,
    email
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users 
WHERE id = $1
LIMIT 1;

-- name: GetBatchUsers :many
SELECT * FROM users
WHERE id = ANY(@ids::BIGINT[]);