package racer

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out wanting for %s and %s", a, b)
	}
}

func ping(url string) chan time.Duration {
	ch := make(chan time.Duration)
	go func() {
		ch <- fetchTime(url)
	}()
	return ch
}

func fetchTime(url string) time.Duration {
	start := time.Now()
	_, _ = http.Get(url)
	return time.Since(start)
}