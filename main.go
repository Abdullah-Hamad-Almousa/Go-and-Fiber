package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func checkURL(url string, ch chan<- string,
	wg *sync.WaitGroup) {

	defer wg.Done()

	client := http.Client{Timeout: 3 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("[DOWN] %s (%v),"+
			url, err)
		return
	}
	defer resp.Body.Close()

	ch <- fmt.Sprintf("[%d] %s",
		resp.StatusCode, url)

}

func scrapeFirstFastest(urls []string) {
	ch := make(chan string)

	for _, url := range urls {
		go func(u string) {
			client := http.Client{
				Timeout: 2 * time.Second,
			}
			resp, err := client.Get(u)
			if err == nil {
				resp.Body.Close()
				ch <- fmt.Sprintf(
					"First to respond: %s [%d]",
					u, resp.StatusCode)
			}
		}(url)
	}

	select {
	case res := <-ch:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("Scraping timed out waiting for fastest response")
	}

}

func main() {
	urls := []string{
		"https://www.google.com",
		"https://golang.org",
		"https://httpbin.org/status/404",
	}

	results := make(chan string, len(urls))
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go checkURL(url, results, &wg)
	}

	wg.Wait()
	close(results)

	for res := range results {
		fmt.Println(res)
	}

	fmt.Println("\n--- Running Fastest Scraper ---")
	scrapeFirstFastest(urls)
	fmt.Printf("\n\n\n")

	fmt.Println("=== 1. defer ===")
	RunDeferExample()

	fmt.Println("\n=== 2. go ===")
	RunGoExample()

	fmt.Println("\n=== 3. chan ===")
	RunChanExample()

	fmt.Println("\n=== 4. <- ===")
	RunArrowExample()

	fmt.Println("\n=== 5. sync ===")
	RunSyncExample()

	fmt.Println("\n=== 6. select ===")
	RunSelectExample()

}
