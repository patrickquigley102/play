package main

import (
	"fmt"
	"net/http"
	"strings"
)

// ServeHTTP serves http requests
func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	score := s.store.getPlayerScore(player)
	fmt.Fprint(w, score)
}

// Server player information
type Server struct {
	store PlayerStore
}

// PlayerStore persists player data
type PlayerStore interface {
	getPlayerScore(string) int
}
