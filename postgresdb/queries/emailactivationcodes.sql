-- name: GetByEmailActivationCode :one
SELECT
    *
FROM
    EmailActivationCodes
WHERE
    activation_code=$1;

-- name: GetActivationCodeByUserID :one
SELECT
    *
FROM
    EmailActivationCodes
WHERE
    user_id=$1;

-- name: GenerateEmailActivationCode :one
INSERT INTO EmailActivationCodes (
    user_id, activation_code, expiration_time
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: VerifyEmailActivationCode :one
SELECT
    *
FROM
    EmailActivationCodes
WHERE
    activation_code=$1
AND
    expiration_time > now() at time zone 'utc';

-- name: RemoveEmailActivationsByUserID :exec
DELETE FROM
    EmailActivationCodes
WHERE
    user_id=$1;