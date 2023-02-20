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

-- name: DeleteProduct :one
DELETE FROM products
WHERE id = $1
RETURNING id;

-- name: GetUserProducts :many
SELECT *
FROM products
WHERE user_id = sqlc.arg('user_id') AND created_at < sqlc.arg('after')
ORDER BY created_at DESC
LIMIT sqlc.arg('first');

-- name: UserProductsHasNextPage :one
SELECT EXISTS(
    SELECT 1
    FROM products
    WHERE user_id = sqlc.arg('user_id') AND created_at < sqlc.arg('after')
);