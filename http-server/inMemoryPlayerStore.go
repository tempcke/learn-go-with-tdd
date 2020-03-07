package main

// InMemoryPlayerStore stores player scores in memory
type InMemoryPlayerStore struct{}

// GetPlayerScore is used to get the score for the given player
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

// RecordWin is used to add a point for name
func (i *InMemoryPlayerStore) RecordWin(name string) {

}
