package racer

import (
	"gotest.tools/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("Racer should return the fastest url", func(t *testing.T) {
		fastServer := sleepyServer( 0 * time.Millisecond)
		slowServer := sleepyServer(10 * time.Millisecond)

		defer fastServer.Close()
		defer	slowServer.Close()

		fastURL := fastServer.URL
		slowURL := slowServer.URL

		assertWinner(t, fastURL, slowURL, fastURL)
		assertWinner(t, fastURL, fastURL, slowURL)
	})

	t.Run("timeout with error after timeout duration", func(t *testing.T) {
		timeout := 50 * time.Millisecond

		serverA := sleepyServer(timeout + (5 * time.Millisecond))
		serverB := sleepyServer(timeout + (6 * time.Millisecond))

		defer serverA.Close()
		defer serverB.Close()

		_, err := Racer(serverA.URL, serverB.URL, timeout)

		if err == nil {
			t.Errorf("Expected timeout error")
		}
	})
}

func assertWinner(t *testing.T, expected, a, b string) {
	winner, err := Racer(a, b, time.Second)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, winner)
}

func sleepyServer(sleepDuration time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(sleepDuration)
		w.WriteHeader(http.StatusOK)
	}))
}