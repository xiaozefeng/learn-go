package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"regexp"
	"time"
)

func main() {
	Crawl("https://blog.csdn.net/TH_NUM/article/details/56837852", 10, &RealFetcher{})
	time.Sleep(10*time.Second)
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLS found on that page
	Fetch(url string) (body string, urls []string, err error)
}

type fakeResult struct {
	body string
	urls []string
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {

	// todo Fetch URLS in parallel.
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
		if isDuplicated(u) {
			continue
		}
		go Crawl(u, depth-1, fetcher)
	}
}

// 获取页面内容，获取页面上的url
var linkReg = regexp.MustCompile(`(https://[^" ']+)`)

type RealFetcher struct {
}

func (f *RealFetcher) Fetch(url string) (body string, urls []string, err error) {
	resp, e := http.Get(url)
	if e != nil {
		return "", nil, e
	}
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", nil, e
	}
	matches := linkReg.FindAllSubmatch(contents, -1)
	for _, m := range  matches{
		fmt.Printf("got url:%s \n", string(m[1]))
	}
	return string(contents), nil, nil
}


var visitedUrls = make(map[string]bool)

// Duplicated
func isDuplicated(url string) bool {
	if visitedUrls[url] {
		return true
	}
	visitedUrls[url] = true
	return false
}

type fakeFetcher map[string]*fakeResult

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
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
