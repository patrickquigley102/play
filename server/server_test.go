package server

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func Test_server_routing(t *testing.T) {
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
			args{w: httptest.NewRecorder(), r: getPlayer(t, "a")},
			http.StatusOK,
		},
		{
			"get league",
			args{w: httptest.NewRecorder(), r: getLeague(t)},
			http.StatusOK,
		},
		{
			"invalid route",
			args{w: httptest.NewRecorder(), r: invalidRoute(t)},
			http.StatusNotFound,
		},
	}
	for _, tt := range tests {
		args := tt.args
		t.Run(tt.name, func(t *testing.T) {
			store := &stubPlayerStore{scores: map[string]int{"a": 1}}
			s := NewServer(store)
			s.ServeHTTP(args.w, args.r)

			gotCode := args.w.Code
			if !reflect.DeepEqual(gotCode, tt.want) {
				t.Errorf("ServeHTTP() Code = %v, want %v", gotCode, tt.want)
			}
		})
	}
}

func Test_server_leagueHandler(t *testing.T) {
	store := &stubPlayerStore{scores: map[string]int{"a": 1}}
	server := NewServer(store)
	t.Run("returns 200", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := getLeague(t)

		server.leagueHandler(w, r)

		if !reflect.DeepEqual(w.Code, http.StatusOK) {
			t.Errorf("leagueHandler() Code = %v, want %v", w.Code, http.StatusOK)
		}
	})
}

func Test_server_playerHandler(t *testing.T) {
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
			args{w: httptest.NewRecorder(), r: getPlayer(t, "a")},
			http.StatusOK,
		},
		{
			"post score",
			args{w: httptest.NewRecorder(), r: postPlayer(t, "a", "1")},
			http.StatusCreated,
		},
		{
			"invalid request",
			args{w: httptest.NewRecorder(), r: invalidPlayers(t)},
			http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		args := tt.args
		t.Run(tt.name, func(t *testing.T) {
			store := &stubPlayerStore{scores: map[string]int{"a": 1}}
			s := NewServer(store)
			s.playerHandler(args.w, args.r)

			gotCode := args.w.Code
			if !reflect.DeepEqual(gotCode, tt.want) {
				t.Errorf("playerHandler() Code = %v, want %v", gotCode, tt.want)
			}
		})
	}
}

func TestServer_getScore(t *testing.T) {
	type args struct {
		w    *httptest.ResponseRecorder
		name string
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
			args{w: httptest.NewRecorder(), name: "a"},
			want{body: "1", code: http.StatusOK},
		},
		{
			"get score, name not found",
			args{w: httptest.NewRecorder(), name: "c"},
			want{body: "Score Not Found", code: http.StatusNotFound},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := &stubPlayerStore{scores: map[string]int{"a": 1}}
			s := NewServer(store)
			s.getScore(tt.args.w, tt.args.name)

			gotBody := tt.args.w.Body.String()
			if !reflect.DeepEqual(gotBody, tt.want.body) {
				t.Errorf("getScore() Body = %v, want %v", gotBody, tt.want.body)
			}

			gotCode := tt.args.w.Code
			if !reflect.DeepEqual(gotCode, tt.want.code) {
				t.Errorf("getScore() Code = %v, want %v", gotCode, tt.want.code)
			}
		})
	}
}

func TestServer_postScore(t *testing.T) {
	type args struct {
		w     *httptest.ResponseRecorder
		name  string
		score string
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
			args{w: httptest.NewRecorder(), name: "a", score: "1"},
			want{
				body: "Score Updated: 1", code: http.StatusCreated,
				updateCalls: []string{"a"},
			},
		},
		{
			"post score, invalid score",
			args{w: httptest.NewRecorder(), name: "a", score: "not an int"},
			want{
				body: "", code: http.StatusBadRequest, updateCalls: []string{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := &stubPlayerStore{
				scores: map[string]int{"a": 0}, updateCalls: []string{},
			}
			s := NewServer(store)

			s.postScore(tt.args.w, tt.args.name, tt.args.score)

			gotBody := tt.args.w.Body.String()
			if !reflect.DeepEqual(gotBody, tt.want.body) {
				t.Errorf("postScore() Body = %v, want %v", gotBody, tt.want.body)
			}

			gotCode := tt.args.w.Code
			if !reflect.DeepEqual(gotCode, tt.want.code) {
				t.Errorf("postScore() Code = %v, want %v", gotCode, tt.want.code)
			}

			gotUpdateCalls := store.updateCalls
			if !reflect.DeepEqual(gotUpdateCalls, tt.want.updateCalls) {
				t.Errorf(
					"postScore() updateCalls = %v, want %v",
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

func (s *stubPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}

func (s *stubPlayerStore) UpdatePlayerScore(name string, score int) {
	s.scores[name] = score
	s.updateCalls = append(s.updateCalls, name)
}

func getPlayer(t *testing.T, name string) *http.Request {
	t.Helper()
	request, _ := http.NewRequest(http.MethodGet, "/players/"+name, nil)
	return request
}

func getLeague(t *testing.T) *http.Request {
	t.Helper()
	request, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return request
}

func invalidPlayers(t *testing.T) *http.Request {
	t.Helper()
	request, _ := http.NewRequest(http.MethodGet, "/players/not/a/route", nil)
	return request
}

func invalidRoute(t *testing.T) *http.Request {
	t.Helper()
	request, _ := http.NewRequest(http.MethodGet, "/not/a/route", nil)
	return request
}

func postPlayer(t *testing.T, name string, score string) *http.Request {
	t.Helper()
	request, _ := http.NewRequest(
		http.MethodPost,
		"/players/"+name+"/"+score,
		nil,
	)
	return request
}
