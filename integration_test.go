package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/patrickquigley102/play/server"
	"github.com/patrickquigley102/play/storesql"
)

func Test_Integration(t *testing.T) {
	config := "./environments/test.yaml"
	store := storesql.NewStoreSQL(config)
	defer store.DB.Close()
	svr := server.NewServer(store)

	t.Run("test get and post score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/players/pq/10", nil)
		svr.ServeHTTP(httptest.NewRecorder(), request)
		request, _ = http.NewRequest(http.MethodPost, "/players/pq/100", nil)
		svr.ServeHTTP(httptest.NewRecorder(), request)

		responseWriter := httptest.NewRecorder()
		request, _ = http.NewRequest(http.MethodGet, "/players/pq", nil)
		svr.ServeHTTP(responseWriter, request)

		wantBody := "100"
		wantStatus := http.StatusOK

		gotBody := responseWriter.Body.String()
		gotStatus := responseWriter.Code

		if gotBody != wantBody {
			t.Errorf("Body = %v, want %v", gotBody, wantBody)
		}
		if gotStatus != wantStatus {
			t.Errorf("Status = %v, want %v", gotStatus, wantStatus)
		}
	})

	t.Run("test get league", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/players/pq/14", nil)
		svr.ServeHTTP(httptest.NewRecorder(), request)

		responseWriter := httptest.NewRecorder()
		request, _ = http.NewRequest(http.MethodGet, "/league", nil)
		svr.ServeHTTP(responseWriter, request)

		wantStatus := http.StatusOK
		wantPlayers := []server.Player{{Name: "pq", Score: 14}}

		var gotPlayers []server.Player
		err := json.NewDecoder(responseWriter.Body).Decode(&gotPlayers)
		gotStatus := responseWriter.Code

		if gotStatus != wantStatus {
			t.Errorf("Status = %v, want %v", gotStatus, wantStatus)
		}

		if err != nil {
			t.Errorf("Unable to parse %q into JSON. err: %v", request.Body, err)
		}

		if !reflect.DeepEqual(gotPlayers, wantPlayers) {
			t.Errorf("LeagueHandler() got = %v, want = %v", gotPlayers, wantPlayers)
		}
	})
}
