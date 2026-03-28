package main

import (
	"net/http"
	"userservice/user"
)

func main() {
	repo := user.NewInMemoryRepo()
	service := user.NewService(repo)
	handler := user.NewHandler(service)

	http.HandleFunc("POST /user/add", handler.CreateUser)
	http.HandleFunc("GET /user/get", handler.GetAllUser)
	http.HandleFunc("GET /user/get/{id}", handler.GetUser)
	http.HandleFunc("DELETE /user/delete/{id}", handler.DeleteUser)

	_ = http.ListenAndServe(":8080", nil)
}
