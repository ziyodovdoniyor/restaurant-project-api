package main

import (
	"restaurant/postgres"
	"restaurant/server"

	_ "github.com/lib/pq"
)

func main() {
	db, err := postgres.Connect()
	if err != nil {
		panic(err)
	}

	repo := postgres.NewPostgresRepository(db)
	r := server.NewRouter(repo)

	r.Run()
}
