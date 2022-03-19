package main

import (
	"fmt"
	"net/http"
	"strings"
)

// Server is a http.Handler
func Server(w http.ResponseWriter, r *http.Request) {
	player := getPlayer(r)
	score := getScore(player)
	fmt.Fprint(w, score)
}

func getPlayer(r *http.Request) string {
	return strings.TrimPrefix(r.URL.Path, "/players/")
}

func getScore(player string) string {
	if player == "a" {
		return "1"
	}
	return "4"
}
