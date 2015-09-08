package models

import (
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"fmt"
	"encoding/json"
	"reflect"
)

type Author struct  {
	Id		int			`json: "id"`
	Name	string		`json: "name"`
}

func AuthorAddNew(name string) (res sql.Result, err error) {

	stmt, err := db.Prepare("INSERT INTO author(name) values(?)")
	checkErr(err)

	return stmt.Exec(name)
}

func (a *Author) Update() (error) {

	stmt, err := db.Prepare("UPDATE author SET name=? where id=?")
	checkErr(err)

	res, err := stmt.Exec(a.Name, a.Id)
	checkErr(err)

	_, err = res.RowsAffected()

	return  err
}

func AuthorGet(val interface{}) (author *Author, err error) {

	author = new(Author)

	var rows *sql.Rows
	switch reflect.TypeOf(val).Name() {
	case "int":
		id := val.(int)
		author.Id = id
		rows, err = db.Query(fmt.Sprintf("SELECT name FROM author WHERE id=%d LIMIT 1", id))
		checkErr(err)

		for rows.Next() {

			if rows != nil {
				var name string
				rows.Scan(&name)
				author.Name = name
			}
		}

	case "string":
		name := val.(string)
		author.Name = name
		rows, err = db.Query(fmt.Sprintf("SELECT id FROM author WHERE name='%s' LIMIT 1", name))
		checkErr(err)

		for rows.Next() {

			if rows != nil {
				var id int
				rows.Scan(&id)
				author.Id = id
			}
		}

	default:
		break
	}

	j, err := json.Marshal(author)
	fmt.Println(string(j))

	return
}