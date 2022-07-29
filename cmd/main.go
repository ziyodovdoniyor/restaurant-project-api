package main

import (
	"log"
	"restaurant/postgres"
	"restaurant/server"
)

func main() {
	db, err := postgres.Connect()
	if err != nil {
		panic(err)
	}

	repo := postgres.NewPostgresRepository(db)
	r := server.NewRouter(repo)

	if err = r.Run(); err != nil {
		log.Fatal("er: ", err)
		return
	}
}
