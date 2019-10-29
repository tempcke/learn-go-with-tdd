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
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		srv := Server(store)

		request, cancel := requestWithCancelContext()
		time.AfterFunc(5*time.Millisecond, cancel)

		response := &SpyResponseWriter{}

		srv.ServeHTTP(response, request)

		if response.written {
			t.Error("a response should not have been written")
		}
	})
}

func requestWithCancelContext() (*http.Request, context.CancelFunc) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)

	cancellingCtx, cancel := context.WithCancel(request.Context())
	return request.WithContext(cancellingCtx), cancel
}