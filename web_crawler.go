package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string, cache *SafeMap) (body string, urls []string, err error)
}

type SafeMap struct {
	mu *sync.Mutex
	v  map[string]string
}

func (s *SafeMap) Get(key string) (string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	res, ok := s.v[key]
	return res, ok
}

func (s *SafeMap) Update(key string, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.v[key] = value
}

// crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func crawl(url string, depth int, fetcher Fetcher, ch chan string, s *SafeMap) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		ch <- url
		return
	}

	body, urls, err := fetcher.Fetch(url, s)

	if err != nil {
		ch <- url
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)
	subCh := make(chan string)

	for _, u := range urls {
		go crawl(u, depth-1, fetcher, subCh, s)
	}

	for i := 0; i < len(urls); i++ {
		<-subCh
	}

	ch <- url
	return
}

func Crawl(url string, depth int, fetcher Fetcher) {
	s := &SafeMap{mu: &sync.Mutex{}, v: make(map[string]string)}
	ch := make(chan string)
	go crawl(url, depth, fetcher, ch, s)
	<-ch
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string, cache *SafeMap) (string, []string, error) {
	if _, ok := cache.Get(url); ok {
		return "", nil, fmt.Errorf("in cache: %q", url)
	}
	if res, ok := f[url]; ok {
		cache.Update(url, res.body)
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}
