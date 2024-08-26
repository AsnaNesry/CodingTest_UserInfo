package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "coding_test/handlers"
)

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
    r.HandleFunc("/users", handlers.GetUsers).Methods("GET")
    r.HandleFunc("/users/{username}", handlers.GetUser).Methods("GET")
    r.HandleFunc("/users/{username}", handlers.UpdateUser).Methods("PUT")
    r.HandleFunc("/users/{username}", handlers.DeleteUser).Methods("DELETE")

    log.Fatal(http.ListenAndServe(":8081", r))
}