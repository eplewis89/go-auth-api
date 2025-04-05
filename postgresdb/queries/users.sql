-- name: CreateUser :one
INSERT INTO
    Users
    (first_name, last_name, email, encr_password, salt_password)
VALUES
    ($1,$2,$3,$4,$5)
RETURNING *;

-- name: FindUserByAccessToken :one
SELECT *
FROM Users u
WHERE u.id
IN
(
    SELECT act.user_id
    FROM AccessTokens act
    WHERE act.token=$1
);

-- name: FindUserById :one
SELECT *
FROM Users
WHERE id=$1;

-- name: FindUserByEmail :one
SELECT *
FROM Users
WHERE email=$1;

-- name: GetTempPasswordForUser :one
SELECT temp_password
FROM Users
WHERE id=$1;

-- name: UpdateUserProfile :exec
UPDATE Users
SET
    first_name=$1,
    last_name=$2
WHERE id=$3;

-- name: UpdateUserPassword :exec
UPDATE Users
SET encr_password=$1
WHERE id=$2;

-- name: UpdateUserTempPassword :exec
UPDATE Users
SET temp_password=$1
WHERE id=$2;

-- name: UpdateUserActivation :exec
UPDATE Users
SET email_activated=$1
WHERE id=$2;

-- name: SetUserDeleted :exec
UPDATE Users
SET
    deleted_at = now() at time zone 'utc',
    is_deleted = TRUE
WHERE id=$1;

-- name: ForceDeleteUser :exec
DELETE FROM Users
WHERE id=$1;

-- name: GetUserWithAccessToken :one
SELECT 
    u."id",
    u."first_name",
    u."last_name",
    u."email",
    u."email_activated",
    u."encr_password",
    u."salt_password",
    u."temp_password",
    a."id" AS "access_token_id",
    a."token" AS "access_token_value",
    a."expiration_time" AS "access_token_expiration",
    a."user_id" AS "access_token_user_id"
FROM Users AS u
INNER JOIN AccessTokens AS a
ON u.id = a.user_id
AND a.token=$1;