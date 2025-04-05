-- name: GetUser :one
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY email;

-- name: CreateUser :one
INSERT INTO users (
  first_name, last_name, email
) VALUES (
  ?, ?, ?
)
RETURNING *;

-- name: UpdateUserName :exec
UPDATE users
set first_name = ?,
last_name = ?
WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?;