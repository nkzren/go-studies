package main

import (
	"fmt"
	"sync"
	"time"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, visited SafeMap) {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	defer wg.Wait()

	go innerCrawl(url, depth, fetcher, visited, wg)
}

func innerCrawl(url string, depth int, fetcher Fetcher, visited SafeMap, wg *sync.WaitGroup) {
	defer wg.Done()

	visited.Visit(url)
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		if !visited.HasVisited(u) {
			wg.Add(1)
			go innerCrawl(u, depth-1, fetcher, visited, wg)
		}
	}
	return
}

type SafeMap struct {
	mu sync.Mutex
	v  map[string]bool
}

func (m *SafeMap) Visit(url string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.v[url] = true
}

func (m *SafeMap) HasVisited(url string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.v[url]
}

func main() {
	visited := SafeMap{v: make(map[string]bool)}
	Crawl("https://golang.org/", 4, fetcher, visited)
}

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	time.Sleep(200 * time.Millisecond)
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
