package main

import (
	"log"
	"net/http"
)

func main() {
	server := NewPlayerServer(NewInMemoryPlayerStore())

	addr := ":5000"
	if err := http.ListenAndServe(addr, server); err != nil {
		log.Fatalf("could not listen on %s %v", addr, err)
	}
}
