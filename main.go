package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	dev := sqlConfig{
		user:     "root",
		password: "",
		host:     "mysql",
		port:     "3306",
		schema:   "play",
	}
	store := NewSQLStore(dev)
	defer store.DB.Close()

	handler := Server{store: store}
	fmt.Println("Listening")
	log.Fatal(http.ListenAndServe(":3000", handler))
}
