package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(PlayerServer)
	port := 5000
	addr := fmt.Sprintf(":%d", port)
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatalf("could not listen on port %d %v", port, err)
	}
}
