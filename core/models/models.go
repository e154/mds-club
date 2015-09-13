package models

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
)

var (
	db *sql.DB
)

func checkErr(err error) {
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
	}
}

func init() {
	var err error
	db, err = sql.Open("sqlite3", "file:db/mds.db?cache=shared")
	if err != nil {
		checkErr(err)
		return
	}

	// todo close
//	db.Close()
}