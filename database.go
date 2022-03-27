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

// NewSQLStore returns a new SQLStore. Tests database connection
func NewSQLStore(filePath string) *SQLStore {
	config := newSQLConfig(filePath)
	db, err := sql.Open("mysql", config.connectionString())
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to DB")
	return &SQLStore{DB: db}
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

func (db SQLStore) updatePlayerScore(name string, score int) {
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
