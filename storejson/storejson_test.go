package storejson

import (
	"reflect"
	"testing"

	"github.com/patrickquigley102/play/server"
)

func TestStoreJSON_GetLeague(t *testing.T) {
	tests := []struct {
		name string
		want []server.Player
	}{
		{"returns Players", []server.Player{{Name: "a", Score: 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := NewStoreJSON("../environments/test.json")
			defer store.database.Close()
			if got := store.GetLeague(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StoreJSON.GetLeague() = %v, want %v", got, tt.want)
			}
		})
	}
}
