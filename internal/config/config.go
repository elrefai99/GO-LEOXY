package config

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func ConnectDatabase() (*sql.DB, error) {

	uri, err := Load()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("pgx", uri.DATABASE_URL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
