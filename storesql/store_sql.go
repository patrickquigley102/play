package storesql

import (
	"database/sql"
	"log"

	// this is a comment to justify the blank import.
	_ "github.com/go-sql-driver/mysql"
)

// StoreSQL interacts with the play database.
type StoreSQL struct {
	DB *sql.DB
}

// NewStoreSQL takes a path to a yaml configuration file and returns a
// StoreSQL connected to the play database.
// yaml file must be formatted like.
//
// user: user
// password: password
// host: host
// port: post
// schema: schema
func NewStoreSQL(path string) *StoreSQL {
	config := newConfigSQL(path)
	database, err := sql.Open("mysql", config.connStr())
	if err != nil {
		log.Print(err)
	}

	err = database.Ping()
	if err != nil {
		log.Print(err)
	}

	log.Println("Connected to DB")
	return &StoreSQL{DB: database}
}

// GetPlayerScore takes a name and returns the score for the first player
// matching that name.
func (db StoreSQL) GetPlayerScore(name string) int {
	var score int
	stmt := "SELECT score FROM players WHERE name = ?;"
	err := db.DB.QueryRow(stmt, name).Scan(&score)

	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
		return 0
	}

	return score
}

// UpdatePlayerScore takes a name and score, updating the score.
func (db StoreSQL) UpdatePlayerScore(name string, score int) {
	stmt, err := db.DB.Prepare("UPDATE players SET score = ? WHERE name = ?")
	if err != nil {
		log.Print(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(score, name)
	if err != nil {
		log.Print(err)
	}
}
