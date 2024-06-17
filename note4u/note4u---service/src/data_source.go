package src

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	Db, err = sql.Open("sqlite3", "./db/note4u1.db")
	if err != nil {
		panic(err)
	}
}
