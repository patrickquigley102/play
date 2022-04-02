package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/patrickquigley102/play/storesql"
)

func main() {
	config := "./environments/development.yaml"
	store := storesql.NewStoreSQL(config)
	defer store.DB.Close()

	server := newServer(store)

	fmt.Println("Listening")
	log.Fatal(http.ListenAndServe(":3000", *server))
}
