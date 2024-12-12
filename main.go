package main

import (
    "log"
    "net/http"
    "github.com/HemanthSanju/URL_Shortner/handlers"
)

func main() {
    http.HandleFunc("/shorten", handlers.HandleShorten)
    http.HandleFunc("/redirect/", handlers.HandleRedirect)
    http.HandleFunc("/metrics", handlers.HandleMetrics)
    log.Println("Starting server on port 8080...")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatalf("Error starting server: %s\n", err)
    }
}