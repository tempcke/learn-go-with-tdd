package main

import (
	"fmt"
	"net/http"
)

// PlayerStore interface defines the requirments to be a PlayerStore
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

// PlayerServer implements serveHTTP and has a PlayerStore
type PlayerServer struct {
	store PlayerStore
	http.Handler
}

// NewPlayerServer constructs and returns a PlayerServer
func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := &PlayerServer{store: store}

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playerHandler))

	p.Handler = router

	return p
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (p *PlayerServer) playerHandler(w http.ResponseWriter, r *http.Request) {

	player := r.URL.Path[len("/players/"):]

	switch r.Method {
	case http.MethodGet:
		p.showScore(w, player)

	case http.MethodPost:
		p.processWin(w, player)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}
