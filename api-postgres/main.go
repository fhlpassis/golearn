package main

import (
	"fmt"
	"net/http"

	"github.com/fhlpassis/golang-postgres/configs"
	"github.com/fhlpassis/golang-postgres/handlers"
	"github.com/go-chi/chi/v5"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	r.Post("/todos", handlers.Create)
	r.Get("/todos", handlers.List)
	r.Get("/todos/{id}", handlers.Get)
	r.Put("/todos/{id}", handlers.Update)
	r.Delete("/todos/{id}", handlers.Delete)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}
