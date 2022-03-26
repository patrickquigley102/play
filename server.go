package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// ServeHTTP serves http requests
func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player, score, err := parseURLParams(r.URL.Path)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if r.Method == http.MethodPost {
		s.PostScore(w, player, score)
	} else {
		s.GetScore(w, player)
	}
}

// GetScore of a player, write http response
func (s Server) GetScore(w http.ResponseWriter, player string) {
	log.Printf("GetScore for %s", player)

	score := s.store.getPlayerScore(player)
	if score > 0 {
		fmt.Fprint(w, score)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Score Not Found")
	}
}

// PostScore of a player, write http response
func (s Server) PostScore(w http.ResponseWriter, player string, score string) {
	log.Printf("PostScore for %s", player)

	scoreInt, err := strconv.Atoi(score)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	s.store.updatePlayerScore(player, scoreInt)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Score Updated: %v", s.store.getPlayerScore(player))
}

func parseURLParams(path string) (string, string, error) {
	errMessage := "Invalid Route"
	bits := strings.Split(path, "/")

	if bits[1] != "players" || len(bits) > 4 {
		return "", "", errors.New(errMessage)
	}

	var score string
	if len(bits) > 3 {
		score = bits[3]
	}

	return bits[2], score, nil
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
