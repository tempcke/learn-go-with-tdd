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
	status int
}{
	{"Pepper", "20", http.StatusOK},
	{"Floyd", "10", http.StatusOK},
	{"Apollo", "0", http.StatusNotFound},
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		scores: map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}
	server := &PlayerServer{store: &store}

	for _, tt := range cases {
		t.Run(tt.player, func(t *testing.T) {
			request, _ := scoreRequest(tt.player)
			response := httptest.NewRecorder()

			server.ServeHTTP(response, request)

			assertStatus(t, response.Code, tt.status)
			assertResponseBody(t, response.Body.String(), tt.score)
		})
	}
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		scores: map[string]int{},
	}

	server := &PlayerServer{&store}

	t.Run("it returns accepted on POST", func(t *testing.T) {
		request, _ := http.NewRequest(
			http.MethodPost,
			"/players/Pepper",
			nil,
		)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)
	})

}

func scoreRequest(player string) (*http.Request, error) {
	uri := fmt.Sprintf("/players/%s", player)
	return http.NewRequest(http.MethodGet, uri, nil)
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Satus: got %d, want %d", got, want)
	}
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("ResponseBody: got %q, want %q", got, want)
	}
}

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}
