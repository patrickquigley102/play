package main

import (
	"log"
	"net/http"

	"github.com/patrickquigley102/play/storesql"
)

func main() {
	config := "./environments/development.yaml"
	store := storesql.NewStoreSQL(config)
	defer store.DB.Close()

	server := newServer(store)

	log.Println("Listening")
	err := http.ListenAndServe(":3000", *server)
	if err != nil {
		log.Print(err)
	}
}
