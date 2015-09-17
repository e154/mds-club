package webserver

import (
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"
)

func wsHandler(w http.ResponseWriter, r *http.Request) {

	ws, err := upgrader.Upgrade(w, r, nil)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(w, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		fmt.Printf(err.Error())
		return
	}

	// регистрация нового клиента
	c := &Client{
		Send: make(chan []byte, 512),
		Ws: ws,
	}

	H.Register <- c
	go c.WritePump()
	c.ReadPump()
}