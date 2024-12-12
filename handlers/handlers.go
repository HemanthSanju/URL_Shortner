package handlers

import (
    "fmt"
    "net/http"
    "strings"
    "../storage"
)

func HandleShorten(w http.ResponseWriter, r *http.Request) {
    originalURL := r.URL.Query().Get("url")
    if originalURL == "" {
        http.Error(w, "URL is required", http.StatusBadRequest)
        return
    }
    shortURL := storage.GetShortURL(originalURL)
    fmt.Fprintf(w, "Shortened URL: http://%s/redirect/%s\n", r.Host, shortURL)
}

func HandleRedirect(w http.ResponseWriter, r *http.Request) {
    shortURL := strings.TrimPrefix(r.URL.Path, "/redirect/")
    originalURL, ok := storage.ResolveURL(shortURL)
    if !ok {
        http.Error(w, "URL not found", http.StatusNotFound)
        return
    }
    http.Redirect(w, r, originalURL, http.StatusFound)
}

func HandleMetrics(w http.ResponseWriter, r *http.Request) {
    metrics := storage.GetTopDomains()
    for domain, count := range metrics {
        fmt.Fprintf(w, "%s: %d\n", domain, count)
    }
}
