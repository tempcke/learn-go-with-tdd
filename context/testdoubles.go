package context1

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	t        *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)
	go s.slowFetch(ctx, data)
	return s.fetchedResponse(ctx, data)
}

func (s *SpyStore) slowFetch(ctx context.Context, data chan string) {
	var result string
	for _, c := range s.response {
		select {
		case <-ctx.Done():
			s.t.Log("spy store got cancelled")
			return
		default:
			time.Sleep(10 * time.Millisecond)
			result += string(c)
		}
	}
	data <- result
}

func (s *SpyStore) fetchedResponse(ctx context.Context, data chan string) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

