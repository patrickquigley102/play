package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func newServer(store playerStorer) *server {
	mux := http.NewServeMux()
	server := server{store: store, mux: mux}
	mux.Handle("/league", http.HandlerFunc(server.LeagueHandler))
	mux.Handle("/players/", http.HandlerFunc(server.PlayerHandler))
	return &server
}

func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s server) LeagueHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (s server) PlayerHandler(w http.ResponseWriter, r *http.Request) {
	name, score, err := parseURLParams(r.URL.Path)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if r.Method == http.MethodPost {
		s.postScore(w, name, score)
	} else {
		s.getScore(w, name)
	}
}

func (s server) getScore(w http.ResponseWriter, name string) {
	log.Printf("getScore for %s", name)

	score := s.store.getPlayerScore(name)
	if score > 0 {
		fmt.Fprint(w, score)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Score Not Found")
	}
}

func (s server) postScore(w http.ResponseWriter, name string, score string) {
	log.Printf("postScore for %s", name)

	scoreInt, err := strconv.Atoi(score)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	s.store.updatePlayerScore(name, scoreInt)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Score Updated: %v", s.store.getPlayerScore(name))
}

func parseURLParams(path string) (string, string, error) {
	errMsg := "Invalid Route"
	bits := strings.Split(path, "/")

	if bits[1] != "players" || len(bits) > 4 {
		return "", "", errors.New(errMsg)
	}

	var score string
	if len(bits) > 3 {
		score = bits[3]
	}

	return bits[2], score, nil
}

type server struct {
	store playerStorer
	mux   *http.ServeMux
}

type playerStorer interface {
	getPlayerScore(string) int
	updatePlayerScore(string, int)
}
