package models

import (
	"time"
	"reflect"
	"fmt"
)

type Book struct {
	Id			int64		`json: "id"`
	Name 		string		`json: "name"`
	Author_id	int64		`json: "author_id"`
	Station_id	int64		`json: "station_id"`
	Datetime	time.Time	`json: "datetime"`
}

func (b *Book) Save() (int64, error) {

	stmt, err := db.Prepare("INSERT INTO book(author_id, name, datetime, station_id) values(?,?,?,?)")
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(b.Author_id, b.Name, b.Datetime, b.Station_id)
	if err != nil {
		return 0, err
	}

	b.Id, err = res.LastInsertId()

	return b.Id, err
}

func (b *Book) Update() error {

	return  nil
}

func BookGetById(id int64) (book *Book, err error) {

	return
}

func BookGetAll(arg interface{}) (books []*Book, err error) {

	switch reflect.TypeOf(arg).String() {
	case "*models.Author":
		return  getAllByAuthor(arg.(*Author))
	case "*models.Station":
		return  getAllByStation(arg.(*Station))
	default:
		break
	}

	return
}

func getAllByAuthor(author *Author) (books []*Book, err error) {

	books = make([]*Book, 0)	//[]

	rows, err := db.Query(fmt.Sprintf("SELECT * FROM book WHERE author_id=%d", author.Id))
	checkErr(err)

	for rows.Next() {

		if rows != nil {
			book := new(Book)
			rows.Scan(&book.Id, &book.Author_id, &book.Name, &book.Datetime, &book.Station_id)
			books = append(books, book)
		}
	}

	return
}

func getAllByStation(station *Station) (books []*Book, err error) {

	books = make([]*Book, 0)	//[]

	rows, err := db.Query(fmt.Sprintf("SELECT * FROM book WHERE station_id=%d", station.Id))
	checkErr(err)

	for rows.Next() {

		if rows != nil {
			book := new(Book)
			rows.Scan(&book.Id, &book.Author_id, &book.Name, &book.Datetime, &book.Station_id)
			books = append(books, book)
		}
	}

	return
}
