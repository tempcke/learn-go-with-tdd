package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

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
	server := NewPlayerServer(&store)

	for _, tt := range cases {
		t.Run(tt.player, func(t *testing.T) {
			request := scoreRequest(tt.player)
			response := httptest.NewRecorder()

			server.ServeHTTP(response, request)

			assertStatus(t, response.Code, tt.status)
			assertResponseBody(t, response.Body.String(), tt.score)
		})
	}
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		scores:   map[string]int{},
		winCalls: []string{},
	}

	server := NewPlayerServer(&store)

	t.Run("it records wins on POST", func(t *testing.T) {
		player := "Pepper"

		request := postWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Errorf("winCalls: got %d, want %d", len(store.winCalls), 1)
		}

		storedPlayer := store.winCalls[0]
		if storedPlayer != player {
			t.Errorf("did not store correct winner got %q want %q", storedPlayer, player)
		}
	})

}

func TestLeague(t *testing.T) {
	store := StubPlayerStore{}
	server := NewPlayerServer(&store)

	t.Run("it returns 200 on /league", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
	})
}

func postWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("/players/%s", name),
		nil,
	)
	return req
}

func scoreRequest(player string) *http.Request {
	uri := fmt.Sprintf("/players/%s", player)
	req, _ := http.NewRequest(http.MethodGet, uri, nil)
	return req
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
