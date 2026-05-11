//Упражнение: веб-краулер
//В этом упражнении вы воспользуетесь возможностями параллельной обработки в Go, чтобы распараллелить работу веб-краулера.
//
//Измените функцию Crawl, чтобы она параллельно запрашивала URL-адреса, не повторяя один и тот же запрос дважды.
//
//Подсказка: вы можете хранить кэшированные URL-адреса на карте, но сами по себе карты небезопасны для параллельной обработки!

package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	var (
		data map[string]*fakeResult
		mu   sync.Mutex
		wg   sync.WaitGroup
	)

	data = make(map[string]*fakeResult)

	var crawl func(string, int)

	crawl = func(url string, depth int) {
		if depth <= 0 {
			return
		}

		mu.Lock()
		if _, ok := data[url]; ok {
			mu.Unlock()
			return
		}

		data[url] = &fakeResult{}
		mu.Unlock()

		urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			return
		}

		//fmt.Printf("found: %s %v\n", url, urls)
		fmt.Printf("fetched: %s \n", url)
		for _, u := range urls {
			wg.Add(1)
			go func() {
				defer wg.Done()
				Crawl(u, depth-1, fetcher)
			}()
		}
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		crawl(url, depth)
	}()

	wg.Wait()

	return
}

func main() {
	start := time.Now()

	Crawl("https://golang.org/", 3, fetcher)
	fmt.Println("Elapsed time:", time.Since(start))
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher struct {
}

type fakeResult struct {
	body string
	urls []string
}

// fetcher is a populated fakeFetcher.
var fetcher = &fakeFetcher{}

func (f *fakeFetcher) Fetch(pageUrl string) ([]string, error) {
	res, err := http.Get(pageUrl)
	if err != nil {
		return nil, fmt.Errorf("not found: %s", pageUrl)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, fmt.Errorf("ошибка парсинга HTML (url: %s): %w", pageUrl, err)
	}

	base, _ := url.Parse(pageUrl)
	var urls []string
	doc.Find("a[href]").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if !exists {
			return
		}

		href = strings.TrimSpace(href)
		if href == "" || strings.HasPrefix(href, "#") ||
			strings.HasPrefix(href, "javascript:") {
			return
		}

		absolute := resolveURL(base, href) // та же функция, что выше
		if absolute != "" {
			urls = append(urls, absolute)
		}
	})

	return urls, nil
}

// resolveURL преобразует href в абсолютный URL относительно base.
func resolveURL(base *url.URL, href string) string {
	// Пропускаем явно не веб-адреса: javascript:, mailto:, tel: и т.п.
	if strings.HasPrefix(href, "javascript:") || strings.HasPrefix(href, "mailto:") || strings.HasPrefix(href, "tel:") {
		return ""
	}

	ref, err := url.Parse(href)
	if err != nil {
		return ""
	}
	return base.ResolveReference(ref).String()
}
