package webserver

import (
	"net/http"
	"strconv"
	models "../models"
	"encoding/json"
)

func getBookFileListHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	book_id, err := strconv.Atoi(r.Form[":book"][0])
	if err != nil {
		checkErr(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	book, err := models.BookGet(book_id)
	if err != nil {
		checkErr(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file_list, err := book.Files()
	if err != nil {
		checkErr(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	msg, err := json.Marshal( &map[string]interface {}{
		"files": file_list,
	})

	if err != nil {
		checkErr(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(msg)
}