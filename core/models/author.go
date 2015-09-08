package models

import (
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"fmt"
	"encoding/json"
	"reflect"
)

type Author struct  {
	Id		int64		`json: "id"`
	Name	string		`json: "name"`
}

func (a *Author) Save() (int64, error) {

	stmt, err := db.Prepare("INSERT INTO author(name) values(?)")
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(a.Name)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()

	return id, err
}

func (a *Author) Update() (error) {

	stmt, err := db.Prepare("UPDATE author SET name=? where id=?")
	checkErr(err)

	res, err := stmt.Exec(a.Name, a.Id)
	checkErr(err)

	_, err = res.RowsAffected()

	return  err
}

func (a *Author) AddBook(b *Book) (int64, error) {

	b.Author_id = a.Id
	return  b.Save()
}

func (a *Author) Remove() error {
	return Remove(a.Id)
}

func AuthorGet(val interface{}) (author *Author, err error) {

	author = new(Author)

	var rows *sql.Rows
	switch reflect.TypeOf(val).Name() {
	case "int64":
		id := val.(int64)
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
				var id int64
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

func Remove(id int64) error {

	stmt, err := db.Prepare("DELETE FROM author WHERE id=?")
	checkErr(err)

	res, err := stmt.Exec(id)
	checkErr(err)

	_, err = res.RowsAffected()

	return err
}