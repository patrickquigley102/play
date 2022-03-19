package main

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestServer(t *testing.T) {
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"get a score",
			args{
				w: httptest.NewRecorder(),
				r: buildRequest("a", t),
			},
			"1",
		},
		{
			"get b score",
			args{
				w: httptest.NewRecorder(),
				r: buildRequest("b", t),
			},
			"4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Server(tt.args.w, tt.args.r)

			got := tt.args.w.Body.String()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server() = %v, want %v", got, tt.want)
			}
		})
	}
}

func buildRequest(playerName string, t *testing.T) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/players/"+playerName, nil)
	return request
}
