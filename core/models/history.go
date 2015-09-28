package models

import (
	"time"
	"fmt"
)

type History struct {
	Id			int64		`json:"id"`
	Book_id		int64		`json: "book_id"`
	Date		time.Time	`json: "date"`
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

func HistoryGetPage(page, limit int) (history []*History, total_items int32, err error) {

	if page > 0 {
		page -= 1
	} else {
		page = 0
	}

	history = make([]*History, 0)

	// rows count
	total_rows, err := db.Query(`select * from history`)
	if err != nil {
		return
	}
	defer total_rows.Close()

	for total_rows.Next() {
		total_items++
	}

	query := fmt.Sprintf(`select * from history order by history.data LIMIT "%d" OFFSET "%d"`, limit, page)
	rows, err := db.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		story := new(History)
		err = rows.Scan(&story.Id, &story.Date, &story.Book_id)
		if err != nil {
			return
		}

		history = append(history, story)
	}

	return
}