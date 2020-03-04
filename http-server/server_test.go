package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var cases = []struct {
	player string
	score  string
}{
	{"Pepper", "20"},
	{"Floyd", "10"},
}

func TestGETPlayers(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.player, func(t *testing.T) {
			request, _ := scoreRequest(tt.player)
			response := httptest.NewRecorder()

			PlayerServer(response, request)

			assertResponseBody(t, response.Body.String(), tt.score)
		})
	}
}

func scoreRequest(player string) (*http.Request, error) {
	uri := fmt.Sprintf("/players/%s", player)
	return http.NewRequest(http.MethodGet, uri, nil)
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
