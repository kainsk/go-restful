-- name: ListProducts :many
SELECT * FROM products
ORDER BY id
LIMIT sqlc.arg('limit')
OFFSET sqlc.arg('offset');

-- name: CreateProduct :one
INSERT INTO products(
    name
) VALUES (
    sqlc.arg('product_name')
) RETURNING *;

-- name: GetProduct :one
SELECT * FROM products
WHERE id = sqlc.arg('product_id')
LIMIT 1;

-- name: UpdateProduct :one
UPDATE products
SET name = sqlc.arg('new_product_name')
WHERE id = sqlc.arg('product_id')
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = sqlc.arg('product_id');

-- name: CountProducts :one
SELECT COUNT(*) FROM products;