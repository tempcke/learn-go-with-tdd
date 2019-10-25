package context1

import (
	"context"
	"gotest.tools/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandler(t *testing.T) {
	data := "hello, world"

	t.Run("returns data from store", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		srv := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		srv.ServeHTTP(response, request)


		value := response.Body.String()

		msg := "\nexpected: %s \nactual:   %s"
		assert.Equal(
			t, data, value,
			msg, data, value)

		//store.assertNotCancelled()
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		t.Skipf("skipped")
		store := &SpyStore{response: data, t: t}
		srv := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)

		request = request.WithContext(cancellingCtx)
		response := httptest.NewRecorder()

		srv.ServeHTTP(response, request)

		//store.assertCancelled()
	})
}

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

//func (s *SpyStore) Cancel() {
//	s.cancelled = true
//}
//func (s *SpyStore) assertCancelled() {
//	s.t.Helper()
//	if !s.cancelled {
//		s.t.Errorf("store was not told to cancel")
//	}
//}
//
//func (s *SpyStore) assertNotCancelled() {
//	s.t.Helper()
//	if s.cancelled {
//		s.t.Errorf("store was told to cancel")
//	}
//}
