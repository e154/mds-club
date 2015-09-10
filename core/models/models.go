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
		return
	}
}

func init() {
	var err error
	db, err = sql.Open("sqlite3", "./db/mds.db")
	checkErr(err)

	// todo close
//	db.Close()
}