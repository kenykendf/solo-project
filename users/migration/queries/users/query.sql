-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: GetUserByPhone :one
SELECT * FROM users
WHERE phone = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id;

-- name: CreateUser :exec
INSERT INTO users (
    username
    ,email
    ,password
    ,phone
    ,first_name
    ,last_name
    ,created_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
);

-- name: UpdateUser :exec
UPDATE users SET
    username = $2
    ,email = $3
    ,password = $4
    ,phone = $5
    ,first_name = $6
    ,last_name = $7
    ,updated_at = $8
WHERE id = $1;

-- name: SoftDeleteUser :exec
UPDATE users SET updated_at = $2, deleted_at = $3
WHERE id = $1;

-- name: HardDeleteUser :exec
DELETE FROM users
WHERE id = $1;