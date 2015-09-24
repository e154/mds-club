package models

import (
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"fmt"
	"reflect"
)

type Author struct  {
	Id			int64		`json: "id"`
	Name		string		`json: "name"`
	Low_name	string		`json: "low_name"`
}

func (a *Author) Save() (id int64, err error) {

	stmt, err := db.Prepare("INSERT INTO author(name, low_name) values(?,?)")
	if err != nil {
		checkErr(err)
		return
	}
	defer stmt.Close()

	res, err := stmt.Exec(a.Name, a.Low_name)
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

	stmt, err := db.Prepare("UPDATE author SET name=?, low_name=? where id=?")
	if err != nil {
		checkErr(err)
		return
	}
	defer stmt.Close()

	res, err := stmt.Exec(a.Name, a.Low_name, a.Id)
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
		rows, err = db.Query(fmt.Sprintf(`SELECT * FROM author WHERE id=%d LIMIT 1`, id))
		if err != nil {
			checkErr(err)
			return
		}
		defer rows.Close()

		author.Id = id
		for rows.Next() {
			if rows != nil {
				rows.Scan(&author.Id, &author.Name, &author.Low_name)
				return
			}
		}

	case "string":
		name := val.(string)
		rows, err = db.Query(fmt.Sprintf(`SELECT * FROM author WHERE name="%s" LIMIT 1`, name))
		if err != nil {
			checkErr(err)
			return
		}
		defer rows.Close()

		author.Name = name
		for rows.Next() {
			if rows != nil {
				rows.Scan(&author.Id, &author.Name, &author.Low_name)
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
		err = rows.Scan(&author.Id, &author.Name, &author.Low_name)
		if err != nil {
			return
		}

		authors = append(authors, author)
	}

	return
}

func AuthorFind(name string, page, items_per_page int) (authors []*Author, total_items int32, err error) {

	if page > 0 {
		page -= 1
	} else {
		page = 0
	}

	authors = make([]*Author, 0)	//[]

	total_rows, err := db.Query(fmt.Sprintf(`SELECT * FROM "author" WHERE "author"."name" LIKE '%s'`, "%"+name+"%"))
	if err != nil {
		return
	}
	defer total_rows.Close()

	for total_rows.Next() {
		total_items++
	}

	rows, err := db.Query(fmt.Sprintf(`SELECT * FROM "author" WHERE "author"."name" LIKE '%s' LIMIT %d OFFSET %d`, "%"+name+"%", items_per_page, page))
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		author := new(Author)
		err = rows.Scan(&author.Id, &author.Name, &author.Low_name)
		if err != nil {
			return
		}

		authors = append(authors, author)
	}

	return
}