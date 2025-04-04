// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: users.sql

package repo

import (
	"context"
	"database/sql"
	"time"
)

const createUser = `-- name: CreateUser :one
INSERT INTO
    Users
    (first_name, last_name, email, encr_password, salt_password)
VALUES
    ($1,$2,$3,$4,$5)
RETURNING id, first_name, last_name, email, email_activated, encr_password, salt_password, temp_password, created_at, updated_at, deleted_at, is_deleted
`

type CreateUserParams struct {
	FirstName    string
	LastName     string
	Email        string
	EncrPassword sql.NullString
	SaltPassword sql.NullString
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.EncrPassword,
		arg.SaltPassword,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.EmailActivated,
		&i.EncrPassword,
		&i.SaltPassword,
		&i.TempPassword,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.IsDeleted,
	)
	return i, err
}

const findUserByAccessToken = `-- name: FindUserByAccessToken :one
SELECT id, first_name, last_name, email, email_activated, encr_password, salt_password, temp_password, created_at, updated_at, deleted_at, is_deleted
FROM Users u
WHERE u.id
IN
(
    SELECT act.user_id
    FROM AccessTokens act
    WHERE act.token=$1
)
`

func (q *Queries) FindUserByAccessToken(ctx context.Context, token string) (User, error) {
	row := q.db.QueryRowContext(ctx, findUserByAccessToken, token)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.EmailActivated,
		&i.EncrPassword,
		&i.SaltPassword,
		&i.TempPassword,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.IsDeleted,
	)
	return i, err
}

const findUserByEmail = `-- name: FindUserByEmail :one
SELECT id, first_name, last_name, email, email_activated, encr_password, salt_password, temp_password, created_at, updated_at, deleted_at, is_deleted
FROM Users
WHERE email=$1
`

func (q *Queries) FindUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, findUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.EmailActivated,
		&i.EncrPassword,
		&i.SaltPassword,
		&i.TempPassword,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.IsDeleted,
	)
	return i, err
}

const findUserById = `-- name: FindUserById :one
SELECT id, first_name, last_name, email, email_activated, encr_password, salt_password, temp_password, created_at, updated_at, deleted_at, is_deleted
FROM Users
WHERE id=$1
`

func (q *Queries) FindUserById(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, findUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.EmailActivated,
		&i.EncrPassword,
		&i.SaltPassword,
		&i.TempPassword,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.IsDeleted,
	)
	return i, err
}

const forceDeleteUser = `-- name: ForceDeleteUser :exec
DELETE FROM Users
WHERE id=$1
`

func (q *Queries) ForceDeleteUser(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, forceDeleteUser, id)
	return err
}

const getTempPasswordForUser = `-- name: GetTempPasswordForUser :one
SELECT temp_password
FROM Users
WHERE id=$1
`

func (q *Queries) GetTempPasswordForUser(ctx context.Context, id int32) (sql.NullString, error) {
	row := q.db.QueryRowContext(ctx, getTempPasswordForUser, id)
	var temp_password sql.NullString
	err := row.Scan(&temp_password)
	return temp_password, err
}

const getUserWithAccessToken = `-- name: GetUserWithAccessToken :one
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
AND a.token=$1
`

type GetUserWithAccessTokenRow struct {
	ID                    int32
	FirstName             string
	LastName              string
	Email                 string
	EmailActivated        bool
	EncrPassword          sql.NullString
	SaltPassword          sql.NullString
	TempPassword          sql.NullString
	AccessTokenID         int32
	AccessTokenValue      string
	AccessTokenExpiration time.Time
	AccessTokenUserID     sql.NullInt32
}

func (q *Queries) GetUserWithAccessToken(ctx context.Context, token string) (GetUserWithAccessTokenRow, error) {
	row := q.db.QueryRowContext(ctx, getUserWithAccessToken, token)
	var i GetUserWithAccessTokenRow
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.EmailActivated,
		&i.EncrPassword,
		&i.SaltPassword,
		&i.TempPassword,
		&i.AccessTokenID,
		&i.AccessTokenValue,
		&i.AccessTokenExpiration,
		&i.AccessTokenUserID,
	)
	return i, err
}

const setUserDeleted = `-- name: SetUserDeleted :exec
UPDATE Users
SET
    deleted_at = now() at time zone 'utc',
    is_deleted = TRUE
WHERE id=$1
`

func (q *Queries) SetUserDeleted(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, setUserDeleted, id)
	return err
}

const updateUserActivation = `-- name: UpdateUserActivation :exec
UPDATE Users
SET email_activated=$1
WHERE id=$2
`

type UpdateUserActivationParams struct {
	EmailActivated bool
	ID             int32
}

func (q *Queries) UpdateUserActivation(ctx context.Context, arg UpdateUserActivationParams) error {
	_, err := q.db.ExecContext(ctx, updateUserActivation, arg.EmailActivated, arg.ID)
	return err
}

const updateUserPassword = `-- name: UpdateUserPassword :exec
UPDATE Users
SET encr_password=$1
WHERE id=$2
`

type UpdateUserPasswordParams struct {
	EncrPassword sql.NullString
	ID           int32
}

func (q *Queries) UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) error {
	_, err := q.db.ExecContext(ctx, updateUserPassword, arg.EncrPassword, arg.ID)
	return err
}

const updateUserProfile = `-- name: UpdateUserProfile :exec
UPDATE Users
SET
    first_name=$1,
    last_name=$2
WHERE id=$3
`

type UpdateUserProfileParams struct {
	FirstName string
	LastName  string
	ID        int32
}

func (q *Queries) UpdateUserProfile(ctx context.Context, arg UpdateUserProfileParams) error {
	_, err := q.db.ExecContext(ctx, updateUserProfile, arg.FirstName, arg.LastName, arg.ID)
	return err
}

const updateUserTempPassword = `-- name: UpdateUserTempPassword :exec
UPDATE Users
SET temp_password=$1
WHERE id=$2
`

type UpdateUserTempPasswordParams struct {
	TempPassword sql.NullString
	ID           int32
}

func (q *Queries) UpdateUserTempPassword(ctx context.Context, arg UpdateUserTempPasswordParams) error {
	_, err := q.db.ExecContext(ctx, updateUserTempPassword, arg.TempPassword, arg.ID)
	return err
}
