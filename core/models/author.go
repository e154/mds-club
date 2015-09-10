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

func (a *Author) Save() (id int64, err error) {

	stmt, err := db.Prepare("INSERT INTO author(name) values(?)")
	checkErr(err)

	res, err := stmt.Exec(a.Name)
	checkErr(err)

	return res.LastInsertId()
}

func (a *Author) Update() (err error) {

	stmt, err := db.Prepare("UPDATE author SET name=? where id=?")
	checkErr(err)

	res, err := stmt.Exec(a.Name, a.Id)
	checkErr(err)

	_, err = res.RowsAffected()

	return
}

func (a *Author) AddBook(b *Book) (int64, error) {

	b.Author_id = a.Id
	return  b.Save()
}

func (a *Author) Remove() error {
	return AuthorRemove(a.Id)
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

func AuthorRemove(id int64) (err error) {

	stmt, err := db.Prepare("DELETE FROM author WHERE id=?")
	checkErr(err)

	res, err := stmt.Exec(id)
	checkErr(err)

	_, err = res.RowsAffected()

	return
}

func AuthorGetAll() (authors []*Author, err error) {

	rows, err := db.Query("SELECT * FROM author")
	if err != nil {
		return
	}

	authors = make([]*Author, 0)	//[]

	for rows.Next() {
		author := new(Author)
		err = rows.Scan(&author.Id, &author.Name)
		if err != nil {
			return
		}

		authors = append(authors, author)
	}

	return
}