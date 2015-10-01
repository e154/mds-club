package models

import (
	"time"
	"fmt"
)

type History struct {
	Id				int64		`json:"id"`
	Book_id			int64		`json: "book_id"`
	Date			time.Time	`json: "date"`
}

type HistoryRes struct {
	Id				int64		`json:"id"`
	Book_name		string		`json:"book_name"`
	Book_Id			int64		`json:"book_id"`
	Author_name		string		`json:"author_name"`
	Open_date		interface{}	`json:"open_date"`
	Station_name	string		`json:"station_name"`
}

func (h *History) Save() (int64, error) {

	stmt, err := db.Prepare("INSERT INTO history(book_id, date) values(?,?)")
	if err != nil {
		checkErr(err)
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(h.Book_id, h.Date)
	if err != nil {
		checkErr(err)
		return 0, err
	}

	h.Id, err = res.LastInsertId()

	return h.Id, err
}

func (h *History) Remove(id int64) (err error) {

	stmt, err := db.Prepare(`DELETE FROM history WHERE id=?`)
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

func HistoryGetPage(page, limit int) (history []*HistoryRes, total_items int32, err error) {

	if page > 0 {
		page = (page - 1) * limit
	} else {
		page = 0
	}

	history = make([]*HistoryRes, 0)

	// rows count
	total_rows, err := db.Query(`select * from history`)
	if err != nil {
		return
	}
	defer total_rows.Close()

	for total_rows.Next() {
		total_items++
	}

	query := fmt.Sprintf(`
		select history.id as id, book.id as book_id, book.name as book_name, author.name as author_name, history.date as open_date, station.name as station_name

		from history
		JOIN book on history.book_id = book.id
		join author on book.author_id = author.id
		join station on book.station_id = station.id

		order by id
		LIMIT "%d" OFFSET "%d"
	`, limit, page)
	rows, err := db.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		story := new(HistoryRes)
		err = rows.Scan(&story.Id, &story.Book_Id, &story.Book_name, &story.Author_name, &story.Open_date, &story.Station_name)
		if err != nil {
			return
		}

		history = append(history, story)
	}

	return
}