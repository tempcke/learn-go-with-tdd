package main

// InMemoryPlayerStore stores player scores in memory
type InMemoryPlayerStore struct{}

// GetPlayerScore is used to get the score for the given player
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}
