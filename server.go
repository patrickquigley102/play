package main

import (
	"fmt"
	"net/http"
	"strings"
)

// ServeHTTP serves http requests
func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		s.PostScore(w, r)
	} else {
		s.GetScore(w, r)
	}
}

// GetScore of a player, write http response
func (s Server) GetScore(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	score := s.store.getPlayerScore(player)

	if score > 0 {
		fmt.Fprint(w, score)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Score Not Found")
	}
}

// PostScore of a player, write http response
func (s Server) PostScore(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Score Updated")
}

// Server player information
type Server struct {
	store PlayerStore
}

// PlayerStore persists player data
type PlayerStore interface {
	getPlayerScore(string) int
}
