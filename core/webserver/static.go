package webserver

import (
	"net/http"
	"fmt"
)

func staticHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("static_source/"+r.URL.Path)
	http.ServeFile(w, r, "static_source/"+r.URL.Path)
}