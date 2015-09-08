package models

import (
	"time"
	"reflect"
	"fmt"
)

type Book struct {
	Id			int			`json: "id"`
	Name 		string		`json: "name"`
	Author_id	int			`json: "author_id"`
	Station_id	int			`json: "station_id"`
	Datetime	time.Time	`json: "datetime"`
}

func (b *Book) Save() (id int, err error) {

	return
}

func BookGetById(id int) (book *Book, err error) {

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
