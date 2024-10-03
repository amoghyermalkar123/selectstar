package db

import (
	"database/sql"
	"log"

	_ "github.com/glebarez/go-sqlite"
)

func DB() *sql.DB {
	db, err := sql.Open("sqlite", "/opt/selectstar/queries.db")
	if err != nil {
		log.Fatal("open err", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("ping err", err)
	}
	return db
}
