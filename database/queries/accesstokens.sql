-- name: GetAccessTokenByUserID :one
SELECT
    *
FROM
    AccessTokens
WHERE
    user_id=$1;

-- name: GetAccessTokenByValue :one
SELECT
    *
FROM
    AccessTokens
WHERE
    token=$1;

-- name: CountValidAccessTokens :one
SELECT
    *
FROM
    AccessTokens
WHERE
    user_id=$1
AND
    expiration_time > now() at time zone 'utc';

-- name: CreateAccessToken :one
INSERT INTO
    AccessTokens
    (token, user_id, expiration_time)
VALUES
    ($1,$2,$3)
RETURNING *;

-- name: UpdateAccessTokenExpirationTime :exec
UPDATE
    AccessTokens
SET
    expiration_time=$1
WHERE
    user_id=$2;

-- name: UpdateAccessToken :exec
UPDATE
    AccessTokens
SET
    token=$1
WHERE
    user_id=$2;

-- name: DeleteAccessToken :exec
DELETE FROM
    AccessTokens
WHERE
    id=$1;