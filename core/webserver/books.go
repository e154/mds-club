package webserver

import (
	"net/http"
	models "../models"
	"encoding/json"
	"strconv"
	"strings"
)

func booksHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	search := r.Form[":search"][0]
	if search == "all" {
		search = "%"
	} else {
		search = strings.ToLower(search)
	}

	author := r.Form[":author"][0]
	if author == "all" {
		author = "%"
	} else {
		author = strings.ToLower(author)
	}

	page, err := strconv.Atoi(r.Form[":page"][0])
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(r.Form[":limit"][0])
	if err != nil {
		limit = 24
	}

	books, total_items, err := models.BookFind(search, author, page, limit)
	if err != nil {
		checkErr(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	msg, err := json.Marshal( &map[string]interface {}{
		"total_items": total_items,
		"books": books,
	})

	if err != nil {
		checkErr(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(msg)
}
