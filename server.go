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

	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, "Score Updated")
		return
	}

	if score > 0 {
		fmt.Fprint(w, score)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Score Not Found")
	}
}

// Server player information
type Server struct {
	store PlayerStore
}

// PlayerStore persists player data
type PlayerStore interface {
	getPlayerScore(string) int
}
