package main

import (
	"net/http"
	"userservice/user"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	repo := user.NewInMemoryRepo()
	service := user.NewService(repo)
	handler := user.NewHandler(service)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Route("/users", func(r chi.Router) {
		r.Get("/", handler.GetAllUser)
		r.Get("/{id}", handler.GetUser)
		r.Post("/", handler.CreateUser)
		r.Delete("/{id}", handler.DeleteUser)
	})

	_ = http.ListenAndServe(":8080", r)
}
