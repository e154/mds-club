package webserver

import (
	"net/http"
)

func staticHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static_source"+r.URL.Path)
}