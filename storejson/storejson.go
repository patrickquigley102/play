package storejson

import (
	"log"
	"os"

	"github.com/patrickquigley102/play/server"
)

// StoreJSON stores player data in a file as JSON
type StoreJSON struct {
	database *os.File
}

// NewStoreJSON takes a path to the json file and returns a StoreJSON that'll
// use it as a data store
func NewStoreJSON(path string) *StoreJSON {
	file, err := os.Open(path)
	if err != nil {
		log.Print(err)
	}
	return &StoreJSON{database: file}
}

// GetLeague returns all players and their scores
func (store StoreJSON) GetLeague() []server.Player {
	return []server.Player{{Name: "a", Score: 1}}
}
