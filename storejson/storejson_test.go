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
	resetDatabase()

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

func resetDatabase() {
	filePath := "../environments/test.json"
	os.Truncate(filePath, 0)
	resetValue := []byte("[{\"name\": \"a\", \"score\": 1}]")
	os.WriteFile(filePath, resetValue, 0644)
}
