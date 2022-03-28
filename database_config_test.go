package main

import (
	"reflect"
	"testing"
)

func Test_configSQL_connStr(t *testing.T) {
	s := configSQL{
		usr: "user",
		pwd: "password",
		hst: "host",
		prt: "port",
		sch: "schema",
	}
	want := "user:password@tcp(host:port)/schema"
	got := s.connStr()
	if got != want {
		t.Errorf("configSQL.connStr() = %v, want %v", got, want)
	}
}

func Test_newConfigSQL(t *testing.T) {
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
			configSQL := newConfigSQL(tt.filePath)
			got := configSQL.connStr()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSqlConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
