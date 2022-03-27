package main

import (
	"reflect"
	"testing"
)

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

func Test_newSQLConfig(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		want     string
	}{
		{
			"valid yaml config",
			"environments/test.yaml",
			"root:@tcp(mysql:3306)/play",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlConfig := newSQLConfig(tt.filePath)
			got := sqlConfig.connectionString()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSqlConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
