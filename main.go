package main

import (
    "net/http"
    "./handlers"
)

func main() {
    http.HandleFunc("/shorten", handlers.HandleShorten)
    http.HandleFunc("/redirect/", handlers.HandleRedirect)
    http.HandleFunc("/metrics", handlers.HandleMetrics)
}
