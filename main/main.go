package main

import (
	"fmt"
	"sync"
	"time"
)

// Fetcher TODO: comment
type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// FetchResult TODO: comment
type FetchResult struct {
	body string
	urls []string
}

// MyFetcher maps url to FetchResult
type MyFetcher map[string]*FetchResult

var fetcher = MyFetcher{
	"https://golang.org/": &FetchResult{
		body: "The Go Programming Language",
		urls: []string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &FetchResult{
		body: "Packages",
		urls: []string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &FetchResult{
		body: "Package fmt",
		urls: []string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &FetchResult{
		body: "Package os",
		urls: []string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

// Cache keep track of fetched urls
type Cache struct {
	urls map[string]int
	mux  sync.Mutex
}

// Inc TODO: comment
func (c *Cache) Inc(url string) {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.urls[url]++
}

// Value TODO: comment
func (c *Cache) Value(url string) int {
	c.mux.Lock()
	defer c.mux.Unlock()

	return c.urls[url]
}

var cache = Cache{
	urls: make(map[string]int),
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		return
	}

	cache.Inc(url)
	if cache.Value(url) > 1 {
		return
	}

	body, urls, err := fetcher.Fetch(url)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)

	for _, u := range urls {
		go Crawl(u, depth-1, fetcher)
	}

	return
}

// Fetch make MyFetcher implementing Fetcher interface
func (f *MyFetcher) Fetch(url string) (body string, urls []string, err error) {
	if res, ok := (*f)[url]; ok {
		return res.body, res.urls, nil
	}

	return "", nil, fmt.Errorf("not found: %s", url)
}

func main() {
	for i := 0; i < 1000; i++ {
		go Crawl("https://golang.org/", 10, &fetcher)
	}

	time.Sleep(time.Second)
}
