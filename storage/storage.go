package storage

import (
    "math/rand"
    "net/url"
    "strings"
    "sync"
)

var (
    urlMap       = make(map[string]string)
    domainCounts = make(map[string]int)
    mu           sync.RWMutex
)

func GetShortURL(originalURL string) string {
    mu.RLock()
    for short, url := range urlMap {
        if url == originalURL {
            mu.RUnlock()
            return short
        }
    }
    mu.RUnlock()

    mu.Lock()
    defer mu.Unlock()
    short := generateShortURL()
    urlMap[short] = originalURL
    incrementDomainCount(originalURL)
    return short
}

func ResolveURL(shortURL string) (string, bool) {
    mu.RLock()
    originalURL, exists := urlMap[shortURL]
    mu.RUnlock()
    return originalURL, exists
}

func GetTopDomains() map[string]int {
    mu.RLock()
    defer mu.RUnlock()
    return copyTopDomains()
}

func generateShortURL() string {
    b := make([]rune, 8)
    letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

func incrementDomainCount(originalURL string) {
    parsedURL, err := url.Parse(originalURL)
    if err != nil {
        return
    }
    domain := strings.Split(parsedURL.Hostname(), ".")[0]
    domainCounts[domain]++
}

func copyTopDomains() map[string]int {
    topDomains := make(map[string]int, len(domainCounts))
    for domain, count := range domainCounts {
        topDomains[domain] = count
    }
    return topDomains
}