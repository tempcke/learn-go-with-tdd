package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	server := &PlayerServer{
		store: &InMemoryPlayerStore{},
	}
	port := 5000
	addr := fmt.Sprintf(":%d", port)
	if err := http.ListenAndServe(addr, server); err != nil {
		log.Fatalf("could not listen on port %d %v", port, err)
	}
}
