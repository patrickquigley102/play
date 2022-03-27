package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type sqlConfig struct {
	user     string
	password string
	host     string
	port     string
	schema   string
}

func (s sqlConfig) connectionString() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		s.user,
		s.password,
		s.host,
		s.port,
		s.schema,
	)
}

func newSQLConfig(filePath string) *sqlConfig {
	file, err := os.Open(filePath)
	checkErr(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	user := strings.TrimPrefix(scanner.Text(), userTag)
	scanner.Scan()
	password := strings.TrimPrefix(scanner.Text(), passwordTag)
	scanner.Scan()
	host := strings.TrimPrefix(scanner.Text(), hostTag)
	scanner.Scan()
	port := strings.TrimPrefix(scanner.Text(), portTag)
	scanner.Scan()
	schema := strings.TrimPrefix(scanner.Text(), schemaTag)

	return &sqlConfig{
		user:     user,
		password: password,
		host:     host,
		port:     port,
		schema:   schema,
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

const (
	userTag     = "user: "
	passwordTag = "password: "
	hostTag     = "host: "
	portTag     = "port: "
	schemaTag   = "schema: "
)
