package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_Integration(t *testing.T) {
	store := &stubPlayerStore{scores: map[string]int{"a": 5}}
	server := Server{store: store}
	responseWriter := httptest.NewRecorder()

	wantBody := "10"
	wantStatus := http.StatusOK

	server.ServeHTTP(httptest.NewRecorder(), buildPostRequest("a", "10", t))

	server.ServeHTTP(responseWriter, buildGetRequest("a", t))
	gotBody := responseWriter.Body.String()
	gotStatus := responseWriter.Code

	if gotBody != wantBody {
		t.Errorf("Body = %v, want %v", gotBody, wantBody)
	}
	if gotStatus != wantStatus {
		t.Errorf("Status = %v, want %v", gotStatus, wantStatus)
	}
}
