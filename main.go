package main

import (
	"log"
	"net/http"
)

func main() {
	handler := Server{store: tempStore{}}
	log.Fatal(http.ListenAndServe(":3000", handler))
}

type tempStore struct{}

func (s tempStore) getPlayerScore(name string) int {
	return 123
}
