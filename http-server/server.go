package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	score := p.store.GetPlayerScore(player)

	fmt.Fprint(w, score)
}

func GetPlayerScore(player string) int {
	switch player {
	case "Pepper":
		return 20
	case "Floyd":
		return 10
	}
	return 0
}
