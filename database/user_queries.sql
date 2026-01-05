-- name: CreateUser :one
INSERT INTO user_service.users (
    full_name, password, address, phone
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM user_service.users 
WHERE id = $1 LIMIT 1;

-- name: GetUserByPhone :one
SELECT * FROM user_service.users 
WHERE phone = $1 LIMIT 1;

-- name: ListUser :many
SELECT * FROM user_service.users
ORDER BY date_created DESC
LIMIT $1 OFFSET $2;

-- name: UpdateUser :one
UPDATE user_service.users
SET 
    full_name = $2, 
    address = $3
WHERE id = $1
RETURNING *;

-- name: UpdatePassword :exec
UPDATE user_service.users
SET password = $2
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM user_service.users
WHERE id = $1;