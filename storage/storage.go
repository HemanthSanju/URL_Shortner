package storage

import (
    "math/rand"
    "net/url"
    "strings"
    "sync"
    "sort"
    "time"
)

var (
    urlMap       = make(map[string]string)
    domainCounts = make(map[string]int)
    mu           sync.RWMutex
)

func init() {
    rand.Seed(time.Now().UnixNano())
}

func GetShortURL(originalURL string) string {
    mu.Lock()
    defer mu.Unlock()
    
    for short, url := range urlMap {
        if url == originalURL {
            return short
        }
    }
    
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
    parts := strings.Split(parsedURL.Hostname(), ".")
    if len(parts) >= 2 {
        domain := parts[len(parts)-2] + "." + parts[len(parts)-1]
        domainCounts[domain]++
    } else {
        domainCounts[parsedURL.Hostname()]++
    }
}

type kv struct {
    Key   string
    Value int
}

func copyTopDomains() map[string]int {
    var ss []kv
    for k, v := range domainCounts {
        ss = append(ss, kv{k, v})
    }
    sort.Slice(ss, func(i, j int) bool {
        return ss[i].Value > ss[j].Value
    })
    topDomains := make(map[string]int)
    for i, kv := range ss {
        if i >= 3 {
            break
        }
        topDomains[kv.Key] = kv.Value
    }
    return topDomains
}
