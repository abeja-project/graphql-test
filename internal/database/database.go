package database

import (
	"context"
	"encoding/json"
	"sync"
)

// Database saves data in memory.
type Database struct {
	mu   sync.RWMutex
	data map[string]map[int]json.RawMessage
}

// New returns an initialized database.
func New() *Database {
	return &Database{
		data: make(map[string]map[int]json.RawMessage),
	}
}

// NewID returns a free id for a model.
func (db *Database) NewID(ctx context.Context, model string) (int, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	var newID int
	for {
		newID++
		if _, ok := db.data[model][newID]; !ok {
			db.data[model][newID] = nil
			break
		}
	}
	return newID, nil
}

// Objects returns all objects for an model.
func (db *Database) Objects(ctx context.Context, model string) (map[int]json.RawMessage, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	return db.data[model], nil
}

// Update changes an object for an id. If the object does not exist, it is
// created.
func (db *Database) Update(ctx context.Context, model string, id int, data json.RawMessage) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if ids := db.data[model]; ids == nil {
		db.data[model] = make(map[int]json.RawMessage)
	}

	db.data[model][id] = data
	return nil
}
