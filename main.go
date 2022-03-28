package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	devConfig := "./environments/development.yaml"
	store := newStoreSQL(devConfig)
	defer store.DB.Close()

	handler := server{store: store}
	fmt.Println("Listening")
	log.Fatal(http.ListenAndServe(":3000", handler))
}
