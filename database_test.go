package main

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
)

func Test_database_getPlayerScore(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	sql := "SELECT score FROM players WHERE name = ?"
	rows := sqlmock.NewRows([]string{"score"})

	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"susan found",
			args{name: "susan"},
			10,
		},
		{
			"bob found",
			args{name: "bob"},
			1,
		},
		{
			"player not found",
			args{name: "does not exist"},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlStore := storeSQL{DB: db}
			mock.ExpectQuery(sql).
				WithArgs(tt.args.name).
				WillReturnRows(rows.AddRow(tt.want))

			if got := sqlStore.getPlayerScore(tt.args.name); got != tt.want {
				t.Errorf("getPlayerScore() = %v, want %v", got, tt.want)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func Test_database_updatePlayerScore(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	sql := "UPDATE players SET score = ?"

	type args struct {
		name  string
		score int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"player exists",
			args{name: "susan", score: 1},
		},
		{
			"player exists 2",
			args{name: "susan", score: 15},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlStore := storeSQL{DB: db}
			mock.ExpectPrepare(sql).ExpectExec().
				WithArgs(tt.args.score, tt.args.name).
				WillReturnResult(sqlmock.NewResult(0, 1))

			sqlStore.updatePlayerScore(tt.args.name, tt.args.score)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
