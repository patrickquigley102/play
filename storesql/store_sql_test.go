package storesql

import (
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/patrickquigley102/play/server"
)

func Test_storesql_GetPlayerScore(t *testing.T) {
	database, mock, _ := sqlmock.New()
	defer database.Close()
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
			sqlStore := StoreSQL{DB: database}
			mock.ExpectQuery(sql).
				WithArgs(tt.args.name).
				WillReturnRows(rows.AddRow(tt.want))

			if got := sqlStore.GetPlayerScore(tt.args.name); got != tt.want {
				t.Errorf("GetPlayerScore() = %v, want %v", got, tt.want)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func Test_storesql_updatePlayerScore(t *testing.T) {
	database, mock, _ := sqlmock.New()
	defer database.Close()
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
			sqlStore := StoreSQL{DB: database}
			mock.ExpectPrepare(sql).ExpectExec().
				WithArgs(tt.args.score, tt.args.name).
				WillReturnResult(sqlmock.NewResult(0, 1))

			sqlStore.UpdatePlayerScore(tt.args.name, tt.args.score)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func Test_storesql_GetLeague(t *testing.T) {
	database, mock, _ := sqlmock.New()
	defer database.Close()

	rows := sqlmock.NewRows([]string{"name", "score"})
	sql := "SELECT name, score FROM players"
	sqlStore := StoreSQL{DB: database}

	mock.
		ExpectQuery(sql).
		WillReturnRows(
			rows.AddRow("pq", 1),
			rows.AddRow("qp", 2),
		)

	want := []server.Player{
		{Name: "pq", Score: 1},
		{Name: "qp", Score: 2},
	}
	got := sqlStore.GetLeague()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("LeagueHandler() got = %v, want = %v", got, want)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
