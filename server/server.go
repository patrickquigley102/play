package server

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Server of http traffic. Embeds a http.Handler.
type Server struct {
	store playerStorer
	http.Handler
}

// NewServer returns a Server.
func NewServer(store playerStorer) *Server {
	mux := http.NewServeMux()
	server := Server{store: store, Handler: mux}
	mux.Handle("/league", http.HandlerFunc(server.leagueHandler))
	mux.Handle("/players/", http.HandlerFunc(server.playerHandler))
	return &server
}

func (s Server) leagueHandler(writer http.ResponseWriter, r *http.Request) {
	writer.WriteHeader(http.StatusOK)
}

func (s Server) playerHandler(writer http.ResponseWriter, req *http.Request) {
	name, score, err := parseURLParams(req.URL.Path)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.Method == http.MethodPost {
		s.postScore(writer, name, score)
	} else {
		s.getScore(writer, name)
	}
}

func (s Server) getScore(writer http.ResponseWriter, name string) {
	log.Printf("getScore for %s", name)

	score := s.store.GetPlayerScore(name)
	if score > 0 {
		fmt.Fprint(writer, score)
	} else {
		writer.WriteHeader(http.StatusNotFound)
		fmt.Fprint(writer, "Score Not Found")
	}
}

func (s Server) postScore(
	writer http.ResponseWriter,
	name string,
	score string,
) {
	log.Printf("postScore for %s", name)

	scoreInt, err := strconv.Atoi(score)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	s.store.UpdatePlayerScore(name, scoreInt)
	writer.WriteHeader(http.StatusCreated)
	fmt.Fprintf(writer, "Score Updated: %v", s.store.GetPlayerScore(name))
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

type playerStorer interface {
	GetPlayerScore(string) int
	UpdatePlayerScore(string, int)
}
