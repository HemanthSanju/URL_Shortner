package handlers

import (
    "net/http"
)

func HandleShorten(w http.ResponseWriter, r *http.Request) {
    originalURL := r.URL.Query().Get("url")
    if originalURL == "" {
        http.Error(w, "URL is required", http.StatusBadRequest)
        return
    }
}

func HandleRedirect(w http.ResponseWriter, r *http.Request) {
}

func HandleMetrics(w http.ResponseWriter, r *http.Request) {
}
