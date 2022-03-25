package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// SQLStore is a PlayerStore backed by a relational database.
type SQLStore struct {
	DB *sql.DB
}

func (db SQLStore) getPlayerScore(name string) int {

	var score int
	sqlStatement := "SELECT score FROM players WHERE name = ?;"
	err := db.DB.QueryRow(sqlStatement, name).Scan(&score)

	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
		return 0
	}

	return score
}

func (db SQLStore) updatePlayerScore(name string) {
}
