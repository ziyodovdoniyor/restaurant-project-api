package db

import (
	"database/sql"
	"fmt"
	"io"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "pulat"
	password = "9"
	dbname   = "restaurant"
)

func connect() *sql.DB {
	db, er := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
	if er != nil {
		log.Fatal("open sql: ", er)
	}

	if er = db.Ping(); er != nil && er != io.EOF {
		log.Fatal("ping: ", er)
	}
	return db
}
