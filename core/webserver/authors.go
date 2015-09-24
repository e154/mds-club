package webserver

import (
	"net/http"
	models "../models"
	"encoding/json"
	"strconv"
	"strings"
)

func authorsFindHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	search := r.Form[":search"][0]
	if search == "" {
		search = "%"
	} else {
		search = strings.ToLower(search)
	}

	page, err := strconv.Atoi(r.Form[":page"][0])
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(r.Form[":limit"][0])
	if err != nil {
		limit = 24
	}

	authors, total_items, err := models.AuthorFind(search, page, limit)
	if err != nil {
		checkErr(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	msg, err := json.Marshal( &map[string]interface {}{
		"total_items": total_items,
		"authors": authors,
	})

	if err != nil {
		checkErr(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(msg)
}

func authorGetByIdHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	id, err := strconv.Atoi(r.Form[":id"][0])
	if err != nil {
		checkErr(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	author, err := models.AuthorGet(id)
	if err != nil {
		checkErr(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	msg, err := json.Marshal( &map[string]interface {}{
		"author": author,
	})

	w.Header().Set("Content-Type", "application/json")
	w.Write(msg)
}