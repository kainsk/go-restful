-- name: ListProducts :many
SELECT * FROM products
WHERE user_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: CreateProduct :one
INSERT INTO products(
    user_id,
    name,
    price
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetProduct :one
SELECT * FROM products
WHERE id = $1
LIMIT 1;

-- name: UpdateProduct :one
UPDATE products
SET
    name = $2,
    price = $3
WHERE id = $1
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1;

-- name: CountProductsByUserID :one
SELECT COUNT(*) FROM products
WHERE user_id = $1;