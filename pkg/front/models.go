package front

import (
	sq "github.com/elgris/sqrl"
	"scaffold/utils/db"
)

type DbLog struct {
	ID        int    `db:"id"`
	Username  string `db:"username"`
	CreatedAt int    `db:"created_at"`
}

func writeLog(username string) (*DbLog, error) {
	if sql, args, err := sq.Insert("logs").
		Columns("username").
		Values(username).
		Returning("*").ToSql(); err != nil {
		return nil, err
	} else {
		dblog := DbLog{}
		err := db.DB().Get(&dblog, sql, args...)
		if err != nil {
			return nil, err
		}
		return &dblog, nil
	}
}
