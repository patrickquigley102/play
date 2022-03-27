package main

import "fmt"

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
