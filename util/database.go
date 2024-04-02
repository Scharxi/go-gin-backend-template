package util

import (
	"database/sql"
	db "github.com/BxfferOverflow/gogintemplate/db/sqlc"
)

var Db db.Store

func GetStore() *db.Store {
	return &Db
}

func InitStore(conn *sql.DB) {
	Db = db.NewStore(conn)
}
