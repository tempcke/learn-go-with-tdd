package racer

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	if fetchTime(a) < fetchTime(b) {
		return a
	}

	return b
}

func fetchTime(url string) time.Duration {
	start := time.Now()
	_, _ = http.Get(url)
	return time.Since(start)
}