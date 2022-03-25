package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(mysql:3306)/play")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	store := SQLStore{DB: db}

	handler := Server{store: store}
	fmt.Println("Listening")
	log.Fatal(http.ListenAndServe(":3000", handler))
}
