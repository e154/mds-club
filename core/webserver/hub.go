package webserver

import (
	"github.com/e154/console"
	"time"
	"math/rand"
	"fmt"
)

type hub struct {
	// зарегистрированные соединения
	connections map[*Client]bool

	// входящие сообщения
	broadcast chan []byte

	// командный канал
	command chan map[*Client][]byte

	// запрос на регистранцию
	Register chan *Client

	// запрос Отменить регистрацию
	unregister chan *Client
}

func (h *hub) PushRoom() {}

var H = &hub{
	connections: make(map[*Client]bool),
	broadcast:   make(chan []byte, maxMessageSize),
	command:   	 make(chan map[*Client][]byte, maxMessageSize),
	Register:    make(chan *Client, 1),
	unregister:  make(chan *Client, 1),
}

func (h *hub) Output(text []byte) {
	H.broadcast <- text
}

func (h *hub) run() {

	console := console.GetPtr()
	console.Output(h)

	quitAll := func() {
		if len(h.connections) == 0 {
			fmt.Println("quit all")
			//...
			fmt.Println("ok")
		}
	}

	for {
		select {
			// запрос на регистранцию
			// отметить что он зарегистрирован
		case c := <-h.Register:
			h.connections[c] = true

            fmt.Printf("client register\n")
            fmt.Printf("total clients: %d\n", len(h.connections))

			// запрос Отменить регистрацию
			// удалить запись о регистрации из массива
			// закрываем канал
		case c := <-h.unregister:
			delete(h.connections, c)

			fmt.Printf("client unregister\n")
			fmt.Printf("total clients: %d\n", len(h.connections))

			close(c.Send)
			quitAll()

			// широковещательные сообщения
			// перебрать всех зарегистрированных,
			// и отправить каждому это сообщение
			// если в канале больше нет сообщений
			// ???? не понял
			// закрыть канал и отменить регистрацию
		case m := <-h.broadcast:
			for c := range h.connections {
				select {
				case c.Send <- m:
				default:
					// если слушателей больше нет, отключим сервис
					quitAll()
					close(c.Send)
					delete(h.connections, c)
				}
			}

		case m := <-h.command:
			for _, val := range m {
				console.Exec(string(val))
			}
		}
	}
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	go H.run()
}
