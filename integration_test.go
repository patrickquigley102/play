package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/patrickquigley102/play/server"
	"github.com/patrickquigley102/play/storesql"
)

func Test_Integration(t *testing.T) {
	config := "./environments/test.yaml"
	store := storesql.NewStoreSQL(config)
	defer store.DB.Close()
	server := server.NewServer(store)
	responseWriter := httptest.NewRecorder()

	wantBody := "100"
	wantStatus := http.StatusOK

	request, _ := http.NewRequest(http.MethodPost, "/players/pq/10", nil)
	server.ServeHTTP(httptest.NewRecorder(), request)
	request, _ = http.NewRequest(http.MethodPost, "/players/pq/100", nil)
	server.ServeHTTP(httptest.NewRecorder(), request)

	request, _ = http.NewRequest(http.MethodGet, "/players/pq", nil)
	server.ServeHTTP(responseWriter, request)
	gotBody := responseWriter.Body.String()
	gotStatus := responseWriter.Code

	if gotBody != wantBody {
		t.Errorf("Body = %v, want %v", gotBody, wantBody)
	}
	if gotStatus != wantStatus {
		t.Errorf("Status = %v, want %v", gotStatus, wantStatus)
	}
}
