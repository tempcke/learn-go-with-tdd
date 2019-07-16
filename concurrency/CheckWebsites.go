package concurrency

import "time"

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			time.Sleep(10000 * time.Millisecond)
			resultChannel <- result{u, wc(u)} // results[u] = wc(u)
		}(url)
	}

	for i := 0; i< len(urls); i++ {
		result := <-resultChannel
		results[result.string] = result.bool
	}

	return results
}