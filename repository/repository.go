// This file contains the repository implementation layer.
package repository

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Repository struct {
	Db *sql.DB
}

type NewRepositoryOptions struct {
	Dsn string
}

func NewRepository(opts NewRepositoryOptions) *Repository {
	db, err := sql.Open("postgres", opts.Dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		fmt.Fprintf(os.Stdout, "Connection to the database failed")
	}

	return &Repository{
		Db: db,
	}
}
