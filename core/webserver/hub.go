package webserver

import (
	"time"
	"math/rand"
    "fmt"
)

type hub struct {
	// зарегистрированные соединения
	connections map[*Client]bool

	// входящие сообщения
	broadcast chan []byte

	// запрос на регистранцию
	Register chan *Client

	// запрос Отменить регистрацию
	unregister chan *Client

	// выход для служб сбора информации
	quit chan bool

	// флаг, сигнализирует что служба сбора информации запущена
	isProcRead bool
}

func (h *hub) PushRoom() {}

var H = &hub{
	connections: make(map[*Client]bool),
	broadcast:   make(chan []byte, maxMessageSize),
	Register:    make(chan *Client, 1),
	unregister:  make(chan *Client, 1),
	quit:		 make(chan bool, 1),
	isProcRead: false,
}

func (h *hub) run() {

	quitAll := func() {
		if len(h.connections) == 0 {
			h.isProcRead = false
			h.quit <- true
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

			// при подулючении запустить, если не запущен сервис сбора информации
			if !h.isProcRead {
				h.isProcRead = true
//				go uptime()
//				go timeinfo();
			}

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

		}
	}
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	go H.run()
}
