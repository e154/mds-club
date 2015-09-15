package models

import (
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"fmt"
	"reflect"
)

type Author struct  {
	Id		int64		`json: "id"`
	Name	string		`json: "name"`
}

func (a *Author) Save() (id int64, err error) {

	stmt, err := db.Prepare("INSERT INTO author(name) values(?)")
	if err != nil {
		checkErr(err)
		return
	}
	defer stmt.Close()

	res, err := stmt.Exec(a.Name)
	if err != nil {
		checkErr(err)
		return
	}

	id, err = res.LastInsertId()
	if err != nil {
		checkErr(err)
		return
	}

	a.Id = id

	return
}

func (a *Author) Update() (err error) {

	stmt, err := db.Prepare("UPDATE author SET name=? where id=?")
	if err != nil {
		checkErr(err)
		return
	}
	defer stmt.Close()

	res, err := stmt.Exec(a.Name, a.Id)
	if err != nil {
		checkErr(err)
		return
	}

	_, err = res.RowsAffected()

	return
}

func (a *Author) AddBook(b *Book) error {

	b.Author_id = a.Id
	return  b.Update()
}

func (a *Author) Remove() error {
	return AuthorRemove(a.Id)
}

func (a *Author) IsAssigned(b *Book) bool {

	if b != nil {
		return b.Author_id == a.Id
	}

	return false
}

func AuthorGet(val interface{}) (author *Author, err error) {

	author = new(Author)

	var rows *sql.Rows
	switch reflect.TypeOf(val).Name() {
	case "int64":
		id := val.(int64)
		rows, err = db.Query(fmt.Sprintf(`SELECT name FROM author WHERE id=%d LIMIT 1`, id))
		if err != nil {
			checkErr(err)
			return
		}
		defer rows.Close()

		author.Id = id
		for rows.Next() {
			if rows != nil {
				rows.Scan(&author.Name)
				return
			}
		}

	case "string":
		name := val.(string)
		rows, err = db.Query(fmt.Sprintf(`SELECT id FROM author WHERE name="%s" LIMIT 1`, name))
		if err != nil {
			checkErr(err)
			return
		}
		defer rows.Close()

		author.Name = name
		for rows.Next() {
			if rows != nil {
				rows.Scan(&author.Id)
				return
			}
		}

	}

	return nil, fmt.Errorf("author not found")
}

func AuthorRemove(id int64) (err error) {

	stmt, err := db.Prepare("DELETE FROM author WHERE id=?")
	if err != nil {
		checkErr(err)
		return
	}
	defer stmt.Close()

	res, err := stmt.Exec(id)
	if err != nil {
		checkErr(err)
		return
	}

	_, err = res.RowsAffected()

	return
}

func AuthorGetAll() (authors []*Author, err error) {

	rows, err := db.Query("SELECT * FROM author")
	if err != nil {
		return
	}
	defer rows.Close()

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