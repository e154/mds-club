package webserver

import (
	"html/template"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	var homeTempl = template.Must(template.ParseFiles("static_source/templates/index.html"))
	data := struct {
		Host       string
	}{r.Host}
	homeTempl.Execute(w, data)
}
