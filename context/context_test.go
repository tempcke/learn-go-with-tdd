package context1

import (
	"gotest.tools/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandler(t *testing.T) {
	data := "hello, world"

	srv := Server(&SpyStore{response: data})
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()
	srv.ServeHTTP(response, request)
	value := response.Body.String()

	msg := "\nexpected: %s \nactual:   %s"
	assert.Equal(
		t, data, value,
		msg, data, value)
}


type SpyStore struct {
	response string
	cancelled bool
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}