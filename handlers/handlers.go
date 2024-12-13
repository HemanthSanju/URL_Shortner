package handlers

import (
    "fmt"
    "net/http"
    "strings"
    "github.com/HemanthSanju/URL_Shortner/storage"
)

func HandleShorten(w http.ResponseWriter, r *http.Request) {
    originalURL := r.URL.Query().Get("url")
    if originalURL == "" {
        http.Error(w, "URL is required", http.StatusBadRequest)
        return
    }
    
    if !strings.HasPrefix(originalURL, "http://") && !strings.HasPrefix(originalURL, "https://") {
        originalURL = "http://" + originalURL
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

    w.Header().Set("Content-Type", "text/plain") // Set content type as plain text for simple output
    for domain, count := range metrics {
        // Remove the top-level domain part (assuming format "secondlevel.tld")
        parts := strings.Split(domain, ".")
        sld := parts[0]  // Get only the second-level domain part
        fmt.Fprintf(w, "%s: %d\n", sld, count)
    }
}