package main

import (
	"context"
	"database/sql"
	_ "embed"
	"log"

	_ "modernc.org/sqlite"

	"github.com/eplewis89/go-auth-api/sqlitedb/repo"
)

//go:embed sqlitedb/migrations/schema.sql
var ddl string

func run_sqlite() error {
	ctx := context.Background()

	conn, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		return err
	}

	// create tables
	if _, err := conn.ExecContext(ctx, ddl); err != nil {
		return err
	}

	db := repo.New(conn)

	user, err := db.CreateUser(ctx, repo.CreateUserParams{
		FirstName: "test",
		LastName:  "user",
		Email:     "test@user.com",
	})

	if err != nil {
		log.Println(err)
	}

	log.Println(user)

	// list all authors
	users, err := db.ListUsers(ctx)
	if err != nil {
		return err
	}
	for _, user := range users {
		log.Println(user)
	}

	return nil
}
