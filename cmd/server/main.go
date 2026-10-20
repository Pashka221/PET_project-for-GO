package main

import (
	"context"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	connString := "postgres://pet_trigan:6Zkxaxsdfghj@localhost:5433/pet_project"
	pool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()
	err = pool.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	h := &Handler{pool: pool}
	http.HandleFunc("/hello", h.sayHello)
	http.HandleFunc("/register", h.register)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
