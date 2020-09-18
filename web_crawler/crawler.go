package main

import (
	"fmt"
	"sync"
)

// Fetcher is fetch interface
type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

// FetchResult is structure
type FetchResult struct {
	result string
	url    string
	body   string
}

// Crawl is fake web Crawler
func Crawl(url string, depth int, fetcher Fetcher) {
	results := make(map[string]FetchResult)
	var w sync.WaitGroup
	w.Add(1)
	go crawl(url, depth, fetcher, results, &w)

	w.Wait()
	for _, result := range results {
		fmt.Println(result.result, ":", result.url, result.body)
	}
}

func crawl(url string, depth int, fetcher Fetcher, rslt map[string]FetchResult, w *sync.WaitGroup) {
	if depth <= 0 {
		w.Done()
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		rslt[url] = FetchResult{"not found", url, ""}
		w.Done()
		return
	}
	rslt[url] = FetchResult{"found", url, body}
	for _, u := range urls {
		w.Add(1)
		go crawl(u, depth-1, fetcher, rslt, w)
	}
	w.Done()
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

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
