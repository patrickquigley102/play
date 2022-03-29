package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	config := "./environments/development.yaml"
	store := newStoreSQL(config)
	defer store.DB.Close()

	handler := server{store: store}
	fmt.Println("Listening")
	log.Fatal(http.ListenAndServe(":3000", handler))
}
