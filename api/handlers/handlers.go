package handlers

import "github.com/godofprodev/tally/api/storage"

type Handlers struct {
	store *storage.MongoStore
}

func NewHandlers(store *storage.MongoStore) *Handlers {
	return &Handlers{store: store}
}
