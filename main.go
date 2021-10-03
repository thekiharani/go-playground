package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/users", GetUsers).Methods(http.MethodGet)
	r.HandleFunc("/users/{id}", ViewUser).Methods(http.MethodGet)
	r.HandleFunc("/users", CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/users{id}", UpdateUser).Methods(http.MethodPut)
	r.HandleFunc("/users{id}", DeleteUser).Methods(http.MethodDelete)

	r.HandleFunc("/posts", GetPosts).Methods("GET")
	r.HandleFunc("/posts/{id}", ViewPost).Methods("GET")
	r.HandleFunc("/posts", CreatePost).Methods("POST")
	r.HandleFunc("/posts{id}", UpdatePost).Methods("PUT")
	r.HandleFunc("/posts{id}", DeletePost).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {
	InitializeMigration()
	InitializeRouter()
}
