package main

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestServer_ServeHTTP(t *testing.T) {
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	type want struct {
		body string
		code int
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			"get a score",
			args{w: httptest.NewRecorder(), r: buildRequest("a", t)},
			want{body: "1", code: http.StatusOK},
		},
		{
			"player not found",
			args{w: httptest.NewRecorder(), r: buildRequest("c", t)},
			want{body: "Score Not Found", code: http.StatusNotFound},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Server{store: stubbedStore()}
			s.ServeHTTP(tt.args.w, tt.args.r)

			gotBody := tt.args.w.Body.String()
			if !reflect.DeepEqual(gotBody, tt.want.body) {
				t.Errorf("ServeHTTP() Body = %v, want %v", gotBody, tt.want.body)
			}

			gotCode := tt.args.w.Code
			if !reflect.DeepEqual(gotCode, tt.want.code) {
				t.Errorf("ServeHTTP() Code = %v, want %v", gotCode, tt.want.code)
			}
		})
	}
}

func buildRequest(playerName string, t *testing.T) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/players/"+playerName, nil)
	return request
}

func stubbedStore() PlayerStore {
	return &stubPlayerStore{
		scores: map[string]int{
			"a": 1,
			"b": 4,
		},
	}
}

type stubPlayerStore struct {
	scores map[string]int
}

func (s *stubPlayerStore) getPlayerScore(name string) int {
	score := s.scores[name]
	return score
}
