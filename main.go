package main

import (
	"net/http"
	"userservice/user"
)

func main() {
	repo := user.NewInMemoryRepo()
	service := user.NewService(repo)
	handler := user.NewHandler(service)

	http.HandleFunc("POST /users", handler.CreateUser)
	http.HandleFunc("GET /users", handler.GetAllUser)
	http.HandleFunc("GET /users/{id}", handler.GetUser)
	http.HandleFunc("DELETE /users/{id}", handler.DeleteUser)

	_ = http.ListenAndServe(":8080", nil)
}
