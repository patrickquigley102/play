package main

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestServer(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/players/Quigley", nil)
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
			"handler request",
			args{
				w: httptest.NewRecorder(),
				r: request,
			},
			"1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Server(tt.args.w, tt.args.r)

			got := tt.args.w.Body.String()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPostRenderer() = %v, want %v", got, tt.want)
			}
		})
	}
}
