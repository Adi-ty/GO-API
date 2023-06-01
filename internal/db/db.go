package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Databse struct {
	Client *sqlx.DB
}

func NewDatabase() (*Databse, error) {
	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_TABLE"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("SSL_MODE"),
	)

	dbConn, err := sqlx.Connect("postgres", connectionString)

	if err != nil {
		return &Databse{}, fmt.Errorf("could not connet to the database: %w", err)
	}

	return &Databse{
		Client: dbConn,
	}, nil
}

func (d *Databse) Ping(ctx context.Context) error {
	return d.Client.DB.PingContext(ctx)
}
