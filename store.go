package main

import (
	"sync"
)

// pointsStore is a global in-memory store for receipt point values.
// It satisfies the requirement that data does not persist after application shutdown.
type pointsStore struct {
	sync.RWMutex
	data map[string]int
}

// singleton instance for global use
var store = &pointsStore{
	data: make(map[string]int),
}

// savePoints stores the points associated with a generated receipt ID.
func savePoints(id string, points int) {
	store.Lock()
	defer store.Unlock()
	store.data[id] = points
}

// getPoints retrieves the points associated with a receipt ID.
// Returns the points and a boolean indicating if the ID was found.
func getPoints(id string) (int, bool) {
	store.RLock()
	defer store.RUnlock()
	points, ok := store.data[id]
	return points, ok
}
