package main

import (
	"context"
	"database/sql"
	_ "embed"
	"errors"
	"log"
	"os"

	postgresdb "github.com/eplewis89/go-auth-api/postgresdb/repo"
	sqlitedb "github.com/eplewis89/go-auth-api/sqlitedb/repo"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

func main() {
	if err := run_sqlite(); err != nil {
		log.Fatal(err)
	}

	if err := run_postgres(); err != nil {
		log.Fatal(err)
	}
}

//go:embed sqlitedb/migrations/schema.sql
var ddl string

func run_sqlite() error {
	ctx := context.Background()

	conn, err := sql.Open("sqlite", "./sqlitedb/example.sqlite")
	if err != nil {
		return err
	}

	// create tables
	if _, err := conn.ExecContext(ctx, ddl); err != nil {
		return err
	}

	db := sqlitedb.New(conn)

	user, err := db.CreateUser(ctx, sqlitedb.CreateUserParams{
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

func run_postgres() error {
	ctx := context.Background()

	connStr := os.Getenv("GO_AUTH_DB_PG_CONN")

	if connStr == "" {
		return errors.New("no conn string")
	}

	conn, err := sqlx.Open("pgx", connStr)
	if err != nil {
		return err
	}
	defer conn.Close()

	db := postgresdb.New(conn)

	// list all authors
	user, err := db.FindUserByEmail(ctx, "some email")
	if err != nil {
		return err
	}
	log.Println(user)

	return nil
}
