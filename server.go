package main

import (
	"fmt"
	"net/http"
)

// Server is a http.Handler
func Server(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "1")
}
