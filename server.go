package main

import (
	"fmt"
	"net/http"
	"strings"
)

// ServeHTTP serves http requests
func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	if r.Method == http.MethodPost {
		s.PostScore(w, player, 5)
	} else {
		s.GetScore(w, player)
	}
}

// GetScore of a player, write http response
func (s Server) GetScore(w http.ResponseWriter, player string) {
	score := s.store.getPlayerScore(player)
	if score > 0 {
		fmt.Fprint(w, score)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Score Not Found")
	}
}

// PostScore of a player, write http response
func (s Server) PostScore(w http.ResponseWriter, player string, score int) {
	s.store.updatePlayerScore(player, score)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Score Updated: %v", s.store.getPlayerScore(player))
}

// Server player information
type Server struct {
	store PlayerStore
}

// PlayerStore persists player data
type PlayerStore interface {
	getPlayerScore(string) int
	updatePlayerScore(string, int)
}
