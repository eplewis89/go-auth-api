package main

import (
	"context"
	"log"

	"github.com/eplewis89/go-auth-api/database/repo"
	"github.com/jmoiron/sqlx"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx := context.Background()

	connStr := "some connection string"

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
