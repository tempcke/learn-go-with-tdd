package racer

import (
	"gotest.tools/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	fastServer := sleepyServer(0 * time.Nanosecond)
	slowServer := sleepyServer(20 * time.Nanosecond)

	defer fastServer.Close()
	defer	slowServer.Close()

	fastURL := fastServer.URL
	slowURL := slowServer.URL

	assert.Equal(t, fastURL, Racer(slowURL, fastURL))
}

func sleepyServer(sleepDuration time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(sleepDuration)
		w.WriteHeader(http.StatusOK)
	}))
}