package main

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/eplewis89/go-auth-api/database/repo"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
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

	db := repo.New(conn)

	// list all authors
	user, err := db.FindUserByEmail(ctx, "some email")
	if err != nil {
		return err
	}
	log.Println(user)

	return nil
}
