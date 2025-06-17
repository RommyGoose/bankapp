package main

import (
    "log"
    "net/http"
    "os"

    "github.com/gorilla/mux"
    "github.com/joho/godotenv"

    "bankapp/internal/handler"
    "bankapp/internal/middleware"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    r := mux.NewRouter()

    // Public routes
    r.HandleFunc("/register", handler.Register).Methods("POST")
    r.HandleFunc("/login", handler.Login).Methods("POST")

    // Protected routes
    s := r.PathPrefix("/").Subrouter()
    s.Use(middleware.AuthMiddleware)
    s.HandleFunc("/accounts", handler.CreateAccount).Methods("POST")

    log.Println("Server started at :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
