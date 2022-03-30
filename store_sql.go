package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type storeSQL struct {
	DB *sql.DB
}

func newStoreSQL(path string) *storeSQL {
	config := newConfigSQL(path)
	db, err := sql.Open("mysql", config.connStr())
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to DB")
	return &storeSQL{DB: db}
}

func (db storeSQL) getPlayerScore(name string) int {
	var score int
	stmt := "SELECT score FROM players WHERE name = ?;"
	err := db.DB.QueryRow(stmt, name).Scan(&score)

	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
		return 0
	}

	return score
}

func (db storeSQL) updatePlayerScore(name string, score int) {
	stmt, err := db.DB.Prepare("UPDATE players SET score = ? WHERE name = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(score, name)
	if err != nil {
		log.Fatal(err)
	}
}
