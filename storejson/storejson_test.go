package storejson

import (
	"os"
	"reflect"
	"testing"

	"github.com/patrickquigley102/play/server"
)

func TestStoreJSON_GetLeague(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		want     []server.Player
	}{
		{
			"success, returns Players",
			"../environments/test.json",
			[]server.Player{{Name: "a", Score: 1}},
		},
		{
			"decode fails, empty array",
			"../environments/test-broken.json",
			[]server.Player{},
		},
	}
	resetDatabase(t)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := NewStoreJSON(tt.filePath)
			defer store.database.Close()
			if got := store.GetLeague(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StoreJSON.GetLeague() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStoreJSON_GetPlayerScore(t *testing.T) {
	tests := []struct {
		testName   string
		path       string
		playerName string
		want       int
	}{
		{"score found", "../environments/test.json", "a", 1},
		{"score not found", "../environments/test.json", "none", 0},
		{"invalid json", "../environments/test-broken.json", "none", 0},
	}
	resetDatabase(t)

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			store := NewStoreJSON(tt.path)
			if got := store.GetPlayerScore(tt.playerName); got != tt.want {
				t.Errorf("StoreJSON.GetPlayerScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func resetDatabase(t *testing.T) {
	t.Helper()
	filePath := "../environments/test.json"
	os.Truncate(filePath, 0)
	resetValue := []byte("[{\"name\": \"a\", \"score\": 1}]")
	os.WriteFile(filePath, resetValue, 0644)
}
