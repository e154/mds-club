package models

import (
	"time"
	"reflect"
	"fmt"
	"strings"
)

type Book struct {
	Id			int64		`json: "id"`
	Name 		string		`json: "name"`
	Low_name 	string		`json: "low_name"`
	Author_id	int64		`json: "author_id"`
	Station_id	int64		`json: "station_id"`
	Date		time.Time	`json: "date"`
	Url			string		`json: "url"`
}

func (b *Book) Save() (int64, error) {

	stmt, err := db.Prepare("INSERT INTO book(author_id, name, low_name, date, station_id, url) values(?,?,?,?,?)")
	if err != nil {
		checkErr(err)
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(b.Author_id, strConv(b.Name), strConv(b.Low_name), b.Date, b.Station_id, b.Url)
	if err != nil {
		checkErr(err)
		return 0, err
	}

	b.Id, err = res.LastInsertId()

	return b.Id, err
}

func (b *Book) Update() (err error) {

	_, err = db.Exec(fmt.Sprintf(`UPDATE book SET author_id=%d, date=%s, name="%s", low_name="%s", station_id=%d, url="%s" WHERE id=%d`, b.Author_id, b.Date, strConv(b.Name), strConv(b.Low_name), b.Station_id, b.Url,  b.Id))
	if err != nil {
		checkErr(err)
		return
	}

	return
}

func (b *Book) AddFile(file *File) (err error) {

	file.Book_id = b.Id
	return file.Update()
}

func (b *Book) Remove() (err error) {
	return BookRemove(b.Id)
}

func (b *Book) FileExist(file *File) bool {

	if file != nil {
		return b.Id == file.Book_id
	}

	return false
}

func (b *Book) Files() ([]*File, error) {

	return FileGetAllByBook(b)
}

func BookRemove(id int64) (err error) {

	stmt, err := db.Prepare(`DELETE FROM book WHERE id=?`)
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

func BookGet(val interface{}) (book *Book, err error) {

	book = new(Book)

	switch reflect.TypeOf(val).Name() {
	case "int64":
		var id int64 = val.(int64)
		rows, err := db.Query(fmt.Sprintf(`SELECT * FROM book WHERE id=%d LIMIT 1`, id))
		if err != nil {
			checkErr(err)
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			if rows != nil {
				rows.Scan(&book.Author_id, &book.Date, &book.Id, &book.Name, &book.Low_name, &book.Station_id, &book.Url)
				book.Id = id
				return book, nil
			}
		}

	case "string":
		var name string = strConv(val.(string))
		rows, err := db.Query(fmt.Sprintf(`SELECT * FROM book WHERE name="%s" LIMIT 1`, name))
		if err != nil {
			checkErr(err)
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			if rows != nil {
				rows.Scan(&book.Author_id, &book.Date, &book.Id, &book.Name, &book.Low_name, &book.Station_id, &book.Url)
				book.Name = name
				return book, nil
			}
		}
	}

	return nil, fmt.Errorf("book not found")
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

	rows, err := db.Query(fmt.Sprintf(`SELECT * FROM book WHERE author_id=%d`, author.Id))
	if err != nil {
		checkErr(err)
		return
	}
	defer rows.Close()

	for rows.Next() {

		if rows != nil {
			book := new(Book)
			rows.Scan(&book.Id, &book.Author_id, &book.Name, &book.Low_name, &book.Date, &book.Station_id, &book.Url)
			books = append(books, book)
		}
	}

	return
}

func getAllByStation(station *Station) (books []*Book, err error) {

	books = make([]*Book, 0)	//[]

	rows, err := db.Query(fmt.Sprintf(`SELECT * FROM book WHERE station_id=%d`, station.Id))
	if err != nil {
		checkErr(err)
		return
	}
	defer rows.Close()

	for rows.Next() {

		if rows != nil {
			book := new(Book)
			rows.Scan(&book.Id, &book.Author_id, &book.Name, &book.Low_name, &book.Date, &book.Station_id, &book.Url)
			books = append(books, book)
		}
	}

	return
}

func strConv(str string) string {
	return strings.Replace(str, "\"", "", -1)
}

func (b *Book) Play() (err error) {

	return nil
}