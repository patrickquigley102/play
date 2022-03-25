package main

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
)

func Test_database_getPlayerScore(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	type args struct {
		name string
	}
	type dbData struct {
		sql  string
		rows *sqlmock.Rows
	}
	tests := []struct {
		name   string
		args   args
		want   int
		dbData dbData
	}{
		{
			"susan found",
			args{name: "susan"},
			10,
			dbData{
				sql:  "SELECT score FROM players WHERE name = ?",
				rows: sqlmock.NewRows([]string{"score"}).AddRow(10),
			},
		},
		{
			"bob found",
			args{name: "bob"},
			1,
			dbData{
				sql:  "SELECT score FROM players WHERE name = ?",
				rows: sqlmock.NewRows([]string{"score"}).AddRow(1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlStore := SQLStore{DB: db}
			mock.ExpectQuery(tt.dbData.sql).
				WithArgs(tt.args.name).
				WillReturnRows(tt.dbData.rows)

			if got := sqlStore.getPlayerScore(tt.args.name); got != tt.want {
				t.Errorf("getPlayerScore() = %v, want %v", got, tt.want)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
