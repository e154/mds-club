package webserver

import (
	"net/http"
	"encoding/json"
	"strconv"
	models "../models"
)


func stationHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	id, err := strconv.Atoi(r.Form[":id"][0])
	if err != nil {
		checkErr(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	station, err := models.StationGet(id)
	if err != nil {
		checkErr(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	msg, err := json.Marshal( &map[string]interface {}{
		"station": station,
	})

	w.Header().Set("Content-Type", "application/json")
	w.Write(msg)
}