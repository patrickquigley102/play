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
			args{w: httptest.NewRecorder(), r: buildPostRequest("a", "1", t)},
			http.StatusCreated,
		},
		{
			"invalid route",
			args{w: httptest.NewRecorder(), r: buildBadRequest()},
			http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := &stubPlayerStore{scores: map[string]int{"a": 1}}
			s := Server{store: store}
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
		w      *httptest.ResponseRecorder
		player string
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
			args{w: httptest.NewRecorder(), player: "a"},
			want{body: "1", code: http.StatusOK},
		},
		{
			"get score, player not found",
			args{w: httptest.NewRecorder(), player: "c"},
			want{body: "Score Not Found", code: http.StatusNotFound},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := &stubPlayerStore{scores: map[string]int{"a": 1}}
			s := Server{store: store}
			s.GetScore(tt.args.w, tt.args.player)

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
		w      *httptest.ResponseRecorder
		player string
		score  string
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
			args{w: httptest.NewRecorder(), player: "a", score: "1"},
			want{
				body:        "Score Updated: 1",
				code:        http.StatusCreated,
				updateCalls: []string{"a"},
			},
		},
		{
			"post score, invalid score",
			args{w: httptest.NewRecorder(), player: "a", score: "not an int"},
			want{
				body:        "",
				code:        http.StatusBadRequest,
				updateCalls: []string{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := &stubPlayerStore{
				scores:      map[string]int{"a": 0},
				updateCalls: []string{},
			}
			s := Server{store: store}

			s.PostScore(tt.args.w, tt.args.player, tt.args.score)

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

func Test_parseURLParams(t *testing.T) {
	tests := []struct {
		name       string
		path       string
		wantPlayer string
		wantScore  string
		wantErr    bool
	}{
		{"valid, no score", "/players/pq", "pq", "", false},
		{"valid", "/players/pq/100", "pq", "100", false},
		{"invalid, prefix", "/playerz/pq", "", "", true},
		{"invalid, too many", "/players/pq/foo/bar", "", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPlayer, gotScore, err := parseURLParams(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotPlayer != tt.wantPlayer {
				t.Errorf("() gotPlayer = %v, want %v", gotPlayer, tt.wantPlayer)
			}
			if gotScore != tt.wantScore {
				t.Errorf("() gotId = %v, want %v", gotScore, tt.wantScore)
			}
		})
	}
}

type stubPlayerStore struct {
	scores      map[string]int
	updateCalls []string
}

func (s *stubPlayerStore) getPlayerScore(name string) int {
	return s.scores[name]
}

func (s *stubPlayerStore) updatePlayerScore(name string, score int) {
	s.scores[name] = score
	s.updateCalls = append(s.updateCalls, name)
}

func buildGetRequest(playerName string, t *testing.T) *http.Request {
	t.Helper()
	request, _ := http.NewRequest(http.MethodGet, "/players/"+playerName, nil)
	return request
}

func buildBadRequest() *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/not/a/route", nil)
	return request
}

func buildPostRequest(
	playerName string,
	score string,
	t *testing.T,
) *http.Request {
	t.Helper()
	request, _ := http.NewRequest(
		http.MethodPost,
		"/players/"+playerName+"/"+score,
		nil,
	)
	return request
}
