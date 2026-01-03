-- name: CreateUser :one
INSERT INTO users (
    fullname, username, password, address, phone
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: Getuser :one
SELECT * FROM users 
WHERE id = $1 LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: ListUser :many
SELECT * FROM users
ORDER BY date_created DESC
LIMIT $1 OFFSET $2;

-- name: UpdateUser :one
UPDATE users
SET 
    fullname = $2, 
    username = $3,
    address = $4
WHERE id = $1
RETURNING *;

-- name: UpdatePassword :exec
UPDATE users
SET password = $2
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;