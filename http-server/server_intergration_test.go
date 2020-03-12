package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryPlayerStore()
	server := NewPlayerServer(store)
	player := "Pepper"

	for i := 0; i < 3; i++ {
		server.ServeHTTP(httptest.NewRecorder(), postWinRequest(player))
	}

	response := httptest.NewRecorder()
	server.ServeHTTP(response, scoreRequest(player))

	assertStatus(t, response.Code, http.StatusOK)
	assertResponseBody(t, response.Body.String(), "3")
}
