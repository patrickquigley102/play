package storejson

import (
	"encoding/json"
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
	var players []server.Player
	err := json.NewDecoder(store.database).Decode(&players)
	if err != nil {
		return []server.Player{}
	}
	return players
}

// GetPlayerScore returns the score for a player
func (store StoreJSON) GetPlayerScore(name string) int {
	players := store.GetLeague()
	for _, player := range players {
		if player.Name == name {
			return player.Score
		}
	}
	return 0
}
