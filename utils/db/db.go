package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var db *sqlx.DB

func New() {
	db = sqlx.MustConnect("sqlite3", "file:data.db?mode=rwc")
	db.SetMaxOpenConns(1)
	db.MustExec("PRAGMA journal_mode = WAL; PRAGMA foreign_keys = true; PRAGMA synchronous = 1")
}

func DB() *sqlx.DB {
	return db
}

func Close() {
	db.Close()
}
