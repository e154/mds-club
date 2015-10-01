package models

import (
	"time"
	"reflect"
	"fmt"
	"strings"
)

type Book struct {
	Id				int64		`json: "id"`
	Name 			string		`json: "name"`
	Low_name 		string		`json: "low_name"`
	Author_id		int64		`json: "author_id"`
	Station_id		int64		`json: "station_id"`
	Date			time.Time	`json: "date"`
	Url				string		`json: "url"`
	Author_name 	string		`json: "author_name"`
	Station_name 	string		`json: "station_name"`
	Last_play	 	interface{}	`json: "last_play"`
	Play_count 		int64		`json: "play_count"`
}

func (b *Book) Save() (int64, error) {

	stmt, err := db.Prepare("INSERT INTO book(author_id, name, low_name, date, station_id, url) values(?,?,?,?,?,?)")
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
	var inject string

	switch reflect.TypeOf(val).Name() {
	case "int":
		id := val.(int)
		book.Id = int64(id)
		inject = fmt.Sprintf(`book.id="%d"`, id)

	case "string":
		var name string = strConv(val.(string))
		book.Name = name
		inject = fmt.Sprintf(`book.name="%s"`, name)
	}


	rows, err := db.Query(fmt.Sprintf(`
		select book.*, a.name as author_name, s.name as station_name, history.date as last_play, count(history.date) as play_count

		from book

		JOIN station as s on s.id = book.station_id
		JOIN author as a on a.id = book.author_id
		left JOIN history  history on history.book_id = book.id

		WHERE %s

		GROUP BY book.id
		order by book.id
	`, inject))
	if err != nil {
		checkErr(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if rows != nil {
			err = rows.Scan(&book.Author_id, &book.Date, &book.Id, &book.Name, &book.Station_id, &book.Url, &book.Low_name, &book.Author_name, &book.Station_name, &book.Last_play, &book.Play_count)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			if book.Last_play == nil {
				book.Last_play = ""
			}

			return book, nil
		}
	}

	return nil, fmt.Errorf("book not found")
}

func BookGetAll(arg interface{}, page, limit int) (books []*Book, err error) {

	var inject string
	books = make([]*Book, 0)	//[]

	switch reflect.TypeOf(arg).String() {
	case "*models.Author":
		inject = fmt.Sprintf(`book.author_id = "%d"`, arg.(*Author).Id)
	case "*models.Station":
		inject = fmt.Sprintf(`book.station_id = "%d"`, arg.(*Author).Id)
	default:
		break
	}

	if page > 0 {
		page = (page - 1) * limit
	} else {
		page = 0
	}

	rows, err := db.Query(fmt.Sprintf(`
		select book.*, a.name as author_name, s.name as station_name, history.date as last_play, count(history.date) as play_count

		from book

		JOIN station as s on s.id = book.station_id
		JOIN author as a on a.id = book.author_id
		left JOIN history  history on history.book_id = book.id

		WHERE %s

		GROUP BY book.id
		order by book.id

		LIMIT "%d" OFFSET "%d"
	`, inject, limit, page))
	if err != nil {
		checkErr(err)
		return
	}
	defer rows.Close()

	for rows.Next() {

		if rows != nil {
			book := new(Book)
			err = rows.Scan(&book.Author_id, &book.Date, &book.Id, &book.Name, &book.Station_id, &book.Url, &book.Low_name, &book.Author_name, &book.Station_name, &book.Last_play, &book.Play_count)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			if book.Last_play == nil {
				book.Last_play = ""
			}

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

func BookFind(book, author string, page, limit int) (books []*Book, total_items int32, err error) {

	if page > 0 {
		page = (page - 1) * limit
	} else {
		page = 0
	}

	books = make([]*Book, 0)	//[]

	query := fmt.Sprintf(`
		select result.*
		from
		(
			SELECT book.*
			from book, author
			WHERE book.low_name LIKE "%s" and author.low_name like "%s" AND book.author_id=author.id
		   order by book.id
		) result
	`, "%"+book+"%", "%"+author+"%")

	// rows count
	total_rows, err := db.Query(query)
	if err != nil {
		return
	}
	defer total_rows.Close()

	for total_rows.Next() {
		total_items++
	}

	// bookd page
	query = fmt.Sprintf(`
		select book.*, a.name as author_name, s.name as station_name, history.date as last_play, count(history.date) as play_count

		from book
		JOIN station as s on s.id = book.station_id
		JOIN author as a on a.id = book.author_id
		left JOIN history  history on history.book_id = book.id

		WHERE  book.low_name LIKE "%s"
		and a.low_name like "%s"
		AND book.author_id=a.id

		GROUP BY book.id
		order by book.id
		LIMIT "%d" OFFSET "%d"

	`, "%"+book+"%", "%"+author+"%", limit, page)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for rows.Next() {
		if rows.Err() != nil {
			fmt.Println(rows.Err())
		}

		book := new(Book)
		err = rows.Scan(&book.Author_id, &book.Date, &book.Id, &book.Name, &book.Station_id, &book.Url, &book.Low_name, &book.Author_name, &book.Station_name, &book.Last_play, &book.Play_count)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if book.Last_play == nil {
			book.Last_play = ""
		}

		books = append(books, book)
	}

	return
}