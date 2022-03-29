package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_Integration(t *testing.T) {
	config := "./environments/test.yaml"
	store := newStoreSQL(config)
	defer store.DB.Close()
	server := server{store: store}
	responseWriter := httptest.NewRecorder()

	wantBody := "100"
	wantStatus := http.StatusOK

	server.ServeHTTP(httptest.NewRecorder(), postPlayer("pq", "10", t))
	server.ServeHTTP(httptest.NewRecorder(), postPlayer("pq", "100", t))

	server.ServeHTTP(responseWriter, getPlayer("pq", t))
	gotBody := responseWriter.Body.String()
	gotStatus := responseWriter.Code

	if gotBody != wantBody {
		t.Errorf("Body = %v, want %v", gotBody, wantBody)
	}
	if gotStatus != wantStatus {
		t.Errorf("Status = %v, want %v", gotStatus, wantStatus)
	}
}
