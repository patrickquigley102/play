package main

import (
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(Server)
	log.Fatal(http.ListenAndServe(":3000", handler))
}
