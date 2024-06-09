// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package repo

import (
	"database/sql"
	"time"
)

type Accesstoken struct {
	ID             int32
	ExpirationTime time.Time
	Token          string
	UserID         sql.NullInt32
}

type DbConfig struct {
	Key   string
	Value sql.NullString
}

type Emailactivationcode struct {
	ID             int32
	ActivationCode string
	ExpirationTime time.Time
	UserID         int32
}

type User struct {
	ID             int32
	FirstName      string
	LastName       string
	Email          string
	EmailActivated bool
	EncrPassword   sql.NullString
	SaltPassword   sql.NullString
	TempPassword   sql.NullString
	CreatedAt      time.Time
	UpdatedAt      sql.NullTime
	DeletedAt      sql.NullTime
	IsDeleted      bool
}
