package webserver

import (
	"time"
	"encoding/json"
	"net/http"
	"strconv"
	models "../models"
)

var (
	last_book_id int
	last_time time.Time
)

func addHistoryHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	book_id, err := strconv.Atoi(r.Form[":book"][0])
	if err != nil {
		checkErr(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if last_book_id == book_id {
		return
	} else {
		last_book_id = book_id
	}

	history := new(models.History)
	history.Date = time.Now()
	history.Book_id = int64(book_id)
	id, err := history.Save()
	if err != nil {
		checkErr(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	msg, err := json.Marshal( &map[string]interface {}{
		"id": id,
	})

	w.Header().Set("Content-Type", "application/json")
	w.Write(msg)
}

func getHistoryHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	page, err := strconv.Atoi(r.Form[":page"][0])
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(r.Form[":limit"][0])
	if err != nil {
		limit = 24
	}

	history, total_items, err := models.HistoryGetPage(page, limit)
	if err != nil {
		checkErr(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	msg, err := json.Marshal( &map[string]interface {}{
		"total_items": total_items,
		"history": history,
	})

	if err != nil {
		checkErr(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(msg)
}