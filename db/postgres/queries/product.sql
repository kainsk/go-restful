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

-- name: CountProductsByUserID :one
SELECT COUNT(*) FROM products
WHERE user_id = $1;

-- name: GetUserProducts :many
SELECT *,
    EXISTS(
        SELECT 1
        FROM products AS pp
        WHERE pp.user_id = sqlc.arg('user_id') AND pp.created_at < (
            SELECT created_at
            FROM products AS ppp
            WHERE ppp.user_id = sqlc.arg('user_id') AND ppp.created_at < sqlc.arg('after')
            ORDER BY created_at ASC
            LIMIT 1
        )       
    )
FROM products AS p
WHERE p.user_id = sqlc.arg('user_id') AND p.created_at < sqlc.arg('after')
ORDER BY p.created_at DESC
LIMIT sqlc.arg('first');