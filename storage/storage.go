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
		urlMap       = make(map[string]string)  // Map to store the short URL to original URL mapping
		domainCounts = make(map[string]int)     // Map to keep track of the count of URLs shortened per domain
		mu           sync.RWMutex               // Mutex to ensure goroutine-safe access to maps
	)

	func init() {
		rand.Seed(time.Now().UnixNano())        // Seed the random number generator
	}

	// GetShortURL retrieves or creates a new short URL for a given original URL.
	// It ensures that each original URL maps to exactly one short URL and increments domain counts.
	func GetShortURL(originalURL string) string {
		mu.Lock()
		defer mu.Unlock()
		// Check if the URL has already been shortened
		for short, url := range urlMap {
			if url == originalURL {
				return short
			}
		}
		// Generate a new short URL
		short := generateShortURL()
		urlMap[short] = originalURL
		incrementDomainCount(originalURL)
		return short
	}

	// ResolveURL looks up the short URL and returns the corresponding original URL if it exists.
	func ResolveURL(shortURL string) (string, bool) {
		mu.RLock()
		originalURL, exists := urlMap[shortURL]
		mu.RUnlock()
		return originalURL, exists
	}

	// GetTopDomains returns a map of the top 3 domains and their counts.
	func GetTopDomains() map[string]int {
		mu.RLock()
		defer mu.RUnlock()
		return copyTopDomains()
	}

	// generateShortURL creates a random 8-character alphanumeric string.
	func generateShortURL() string {
		b := make([]rune, 8)
		letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
		for i := range b {
			b[i] = letters[rand.Intn(len(letters))]
		}
		return string(b)
	}

	// incrementDomainCount increments the count of URLs shortened for a specific domain.
	func incrementDomainCount(originalURL string) {
		parsedURL, err := url.Parse(originalURL)
		if err != nil {
			return  // Handle error appropriately depending on your application's requirements
		}
		domain := strings.Split(parsedURL.Hostname(), ".")[0]
		domainCounts[domain]++
	}

	// copyTopDomains creates a sorted list of domains by count and returns the top 3.
	func copyTopDomains() map[string]int {
		// Make a slice of all domains and their counts
		type kv struct {
			Key   string
			Value int
		}

		var ss []kv
		for k, v := range domainCounts {
			ss = append(ss, kv{k, v})
		}

		// Sort slice based on the count
		sort.Slice(ss, func(i, j int) bool {
			return ss[i].Value > ss[j].Value
		})

		// Pick top 3 domains
		topDomains := make(map[string]int)
		for i, kv := range ss {
			if i >= 3 {
				break
			}
			topDomains[kv.Key] = kv.Value
		}
		return topDomains
	}