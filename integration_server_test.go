package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/patrickquigley102/play/storesql"
)

func TestServer_Integration(t *testing.T) {
	config := "./environments/test.yaml"
	store := storesql.NewStoreSQL(config)
	defer store.DB.Close()
	server := newServer(store)
	responseWriter := httptest.NewRecorder()

	wantBody := "100"
	wantStatus := http.StatusOK

	server.ServeHTTP(httptest.NewRecorder(), postPlayer(t, "pq", "10"))
	server.ServeHTTP(httptest.NewRecorder(), postPlayer(t, "pq", "100"))

	server.ServeHTTP(responseWriter, getPlayer(t, "pq"))
	gotBody := responseWriter.Body.String()
	gotStatus := responseWriter.Code

	if gotBody != wantBody {
		t.Errorf("Body = %v, want %v", gotBody, wantBody)
	}
	if gotStatus != wantStatus {
		t.Errorf("Status = %v, want %v", gotStatus, wantStatus)
	}
}
