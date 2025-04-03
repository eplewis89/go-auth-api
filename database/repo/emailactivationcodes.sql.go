// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: emailactivationcodes.sql

package repo

import (
	"context"
	"time"
)

const generateEmailActivationCode = `-- name: GenerateEmailActivationCode :one
INSERT INTO EmailActivationCodes (
    user_id, activation_code, expiration_time
) VALUES (
    $1, $2, $3
)
RETURNING id, activation_code, expiration_time, user_id
`

type GenerateEmailActivationCodeParams struct {
	UserID         int32
	ActivationCode string
	ExpirationTime time.Time
}

func (q *Queries) GenerateEmailActivationCode(ctx context.Context, arg GenerateEmailActivationCodeParams) (Emailactivationcode, error) {
	row := q.db.QueryRowContext(ctx, generateEmailActivationCode, arg.UserID, arg.ActivationCode, arg.ExpirationTime)
	var i Emailactivationcode
	err := row.Scan(
		&i.ID,
		&i.ActivationCode,
		&i.ExpirationTime,
		&i.UserID,
	)
	return i, err
}

const getActivationCodeByUserID = `-- name: GetActivationCodeByUserID :one
SELECT
    id, activation_code, expiration_time, user_id
FROM
    EmailActivationCodes
WHERE
    user_id=$1
`

func (q *Queries) GetActivationCodeByUserID(ctx context.Context, userID int32) (Emailactivationcode, error) {
	row := q.db.QueryRowContext(ctx, getActivationCodeByUserID, userID)
	var i Emailactivationcode
	err := row.Scan(
		&i.ID,
		&i.ActivationCode,
		&i.ExpirationTime,
		&i.UserID,
	)
	return i, err
}

const getByEmailActivationCode = `-- name: GetByEmailActivationCode :one
SELECT
    id, activation_code, expiration_time, user_id
FROM
    EmailActivationCodes
WHERE
    activation_code=$1
`

func (q *Queries) GetByEmailActivationCode(ctx context.Context, activationCode string) (Emailactivationcode, error) {
	row := q.db.QueryRowContext(ctx, getByEmailActivationCode, activationCode)
	var i Emailactivationcode
	err := row.Scan(
		&i.ID,
		&i.ActivationCode,
		&i.ExpirationTime,
		&i.UserID,
	)
	return i, err
}

const removeEmailActivationsByUserID = `-- name: RemoveEmailActivationsByUserID :exec
DELETE FROM
    EmailActivationCodes
WHERE
    user_id=$1
`

func (q *Queries) RemoveEmailActivationsByUserID(ctx context.Context, userID int32) error {
	_, err := q.db.ExecContext(ctx, removeEmailActivationsByUserID, userID)
	return err
}

const verifyEmailActivationCode = `-- name: VerifyEmailActivationCode :one
SELECT
    id, activation_code, expiration_time, user_id
FROM
    EmailActivationCodes
WHERE
    activation_code=$1
AND
    expiration_time > now() at time zone 'utc'
`

func (q *Queries) VerifyEmailActivationCode(ctx context.Context, activationCode string) (Emailactivationcode, error) {
	row := q.db.QueryRowContext(ctx, verifyEmailActivationCode, activationCode)
	var i Emailactivationcode
	err := row.Scan(
		&i.ID,
		&i.ActivationCode,
		&i.ExpirationTime,
		&i.UserID,
	)
	return i, err
}
