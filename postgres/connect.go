package postgres

import (
	"database/sql"
	"fmt"
)

const (
	pgHost = "localhost"
	pgPort = 5432
	pgUser = "pulat"
	pgPass = "9"
	pdDB   = "cafe"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open(
		"postgres",
		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			pgHost, pgPort, pgUser, pgPass, pdDB,
		),
	)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
