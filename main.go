package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/vanessanunes/frete-rapido/adapter/handlers"
	"github.com/vanessanunes/frete-rapido/adapter/postgres"
	"github.com/vanessanunes/frete-rapido/adapter/postgres/quoterepository"
	"github.com/vanessanunes/frete-rapido/configs"
)

func init() {
	err := configs.Load()
	if err != nil {
		log.Println(err)
	}
}

func main() {
	conn := postgres.OpenConnection()
	defer conn.Close()

	db := quoterepository.ConnectionRepository(conn)

	handler := handlers.Repository{
		DB: *db,
	}

	r := chi.NewRouter()
	r.Post("/quote", handler.Create)
	r.Get("/metrics", handler.Metrics)

	http.ListenAndServe(fmt.Sprintf(":%s", "8080"), r)
}
