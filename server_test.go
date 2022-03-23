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
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"get score",
			args{w: httptest.NewRecorder(), r: buildGetRequest("a", t)},
			http.StatusOK,
		},
		{
			"post score",
			args{w: httptest.NewRecorder(), r: buildPostRequest("a", t)},
			http.StatusCreated,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Server{store: stubbedStore()}
			s.ServeHTTP(tt.args.w, tt.args.r)

			gotCode := tt.args.w.Code
			if !reflect.DeepEqual(gotCode, tt.want) {
				t.Errorf("ServeHTTP() Code = %v, want %v", gotCode, tt.want)
			}
		})
	}
}

func TestServer_GetScore(t *testing.T) {
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
			"get score",
			args{w: httptest.NewRecorder(), r: buildGetRequest("a", t)},
			want{body: "1", code: http.StatusOK},
		},
		{
			"get score, player not found",
			args{w: httptest.NewRecorder(), r: buildGetRequest("c", t)},
			want{body: "Score Not Found", code: http.StatusNotFound},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Server{store: stubbedStore()}
			s.GetScore(tt.args.w, tt.args.r)

			gotBody := tt.args.w.Body.String()
			if !reflect.DeepEqual(gotBody, tt.want.body) {
				t.Errorf("GetScore() Body = %v, want %v", gotBody, tt.want.body)
			}

			gotCode := tt.args.w.Code
			if !reflect.DeepEqual(gotCode, tt.want.code) {
				t.Errorf("GetScore() Code = %v, want %v", gotCode, tt.want.code)
			}
		})
	}
}

func TestServer_PostScore(t *testing.T) {
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	type want struct {
		body        string
		code        int
		updateCalls []string
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			"post score",
			args{w: httptest.NewRecorder(), r: buildPostRequest("a", t)},
			want{
				body:        "Score Updated",
				code:        http.StatusCreated,
				updateCalls: []string{"a"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := &stubPlayerStore{}
			s := Server{store: store}
			s.PostScore(tt.args.w, tt.args.r)

			gotBody := tt.args.w.Body.String()
			if !reflect.DeepEqual(gotBody, tt.want.body) {
				t.Errorf("PostScore() Body = %v, want %v", gotBody, tt.want.body)
			}

			gotCode := tt.args.w.Code
			if !reflect.DeepEqual(gotCode, tt.want.code) {
				t.Errorf("PostScore() Code = %v, want %v", gotCode, tt.want.code)
			}

			gotUpdateCalls := store.updateCalls
			if !reflect.DeepEqual(gotUpdateCalls, tt.want.updateCalls) {
				t.Errorf(
					"PostScore() updateCalls = %v, want %v",
					gotUpdateCalls,
					tt.want.updateCalls,
				)
			}
		})
	}
}

func stubbedStore() PlayerStore {
	return &stubPlayerStore{scores: map[string]int{"a": 1}}
}

type stubPlayerStore struct {
	scores      map[string]int
	updateCalls []string
}

func (s *stubPlayerStore) getPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *stubPlayerStore) updatePlayerScore(name string, score int) {
	s.updateCalls = append(s.updateCalls, name)
}

func buildGetRequest(playerName string, t *testing.T) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/players/"+playerName, nil)
	return request
}

func buildPostRequest(playerName string, t *testing.T) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, "/players/"+playerName, nil)
	return request
}
