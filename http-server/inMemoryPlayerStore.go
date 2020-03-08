package main

// NewInMemoryPlayerStore constructs and returns an InMemoryPlayerStore
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

// InMemoryPlayerStore stores player scores in memory
type InMemoryPlayerStore struct {
	store map[string]int
}

// GetPlayerScore is used to get the score for the given player
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

// RecordWin is used to add a point for name
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}
