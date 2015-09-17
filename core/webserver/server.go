package webserver

import (
    "net/http"
    "time"
    "github.com/gorilla/websocket"
	"github.com/gorilla/mux"
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
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/ws", wsHandler)
	r.HandleFunc("/js/", staticHandler)
	r.HandleFunc("/css/", staticHandler)
	r.HandleFunc("/images/", staticHandler)
	r.HandleFunc("/templates/", staticHandler)
	http.Handle("/", r)

    go http.ListenAndServe(address, nil)
}
