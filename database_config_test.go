package main

import "testing"

func Test_sqlConfig_connectionString(t *testing.T) {
	s := sqlConfig{
		user:     "user",
		password: "password",
		host:     "host",
		port:     "port",
		schema:   "schema",
	}
	want := "user:password@tcp(host:port)/schema"
	got := s.connectionString()
	if got != want {
		t.Errorf("sqlConfig.connectionString() = %v, want %v", got, want)
	}
}
