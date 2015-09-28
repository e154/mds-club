package webserver

import (
    "net/http"
    "time"
    "github.com/gorilla/websocket"
	"github.com/gorilla/pat"
	"fmt"
)

const (
// Time allowed to write a message to the client.
// Время разрешено написать сообщение клиенту.
    writeWait = 10 * time.Second

// Time allowed to read the next message from the client.
// Время разрешено читать следующее сообщение от клиента.
    readWait = 60 * time.Second

// Отправить пинги к клиенту с этим периодом.
// Должна быть меньше, чем readWait.
    pingPeriod = (readWait * 9) / 10

// Максимальный размер сообщений от клиента.
    maxMessageSize = 512
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    // разрешаем подключение к сокету с разных доменов
    CheckOrigin: func(r *http.Request) bool { return true },
}

func Run(address string) {

	// routes
	r := pat.New()
	r.Get("/ws", wsHandler)
	r.Get("/js/", staticHandler)
	r.Get("/css/", staticHandler)
	r.Get("/images/", staticHandler)
	r.Get("/templates/", staticHandler)
	r.Get("/fonts/", staticHandler)
	r.Get("/api/authors/page~{page:[0-9]+}/limit~{limit:[0-9]+}/search={search:[а-яА-Яa-zA-Z0-9]*}", authorsFindHandler)
	r.Get("/api/authors/id~{id:[0-9]+}", authorGetByIdHandler)
	r.Get("/api/station/id~{id:[0-9]+}", stationHandler)
	r.Get("/api/books/page~{page:[0-9]+}/limit~{limit:[0-9]+}/author~{author}/search={search:[а-яА-Яa-zA-Z0-9]*}", booksHandler)
	r.Get("/api/file/list/book~{book:[0-9]+}", getBookFileListHandler)
	r.Get("/", homeHandler)
	http.Handle("/", r)

    if err := http.ListenAndServe(address, nil); err != nil {
		fmt.Println(err.Error())
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
	}
}