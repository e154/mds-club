package webserver

import (
	"time"
	"io/ioutil"
	"encoding/json"
	"github.com/gorilla/websocket"
)

// Client посредник в соединении между websocket и hub.
type Client struct {
	// websocket соединение.
	Ws *websocket.Conn

	// буферизированный канал сообщений.
	Send chan []byte
}

// проталкивание сообщений от клиента до hub
func (c *Client) ReadPump() {
	defer func() {
		H.unregister <- c
		c.Ws.Close()
	}()

	c.Ws.SetReadLimit(maxMessageSize)
	c.Ws.SetReadDeadline(time.Now().Add(readWait))
	c.Ws.SetPongHandler(func(string) error { c.Ws.SetReadDeadline(time.Now().Add(readWait)); return nil })

	for {
		op, r, err := c.Ws.NextReader()
		if err != nil {
			break
		}
		switch op {
		case websocket.TextMessage:
			message, err := ioutil.ReadAll(r)
			if err != nil {
				break
			}
			//c.username + "_" + time.Now().Format("15:04:05") + ":" + string(message)
			j, _ := json.Marshal(&map[string]interface {}{"message": string(message)})
		H.broadcast <- []byte(j)
		}
	}
}

// write writes a message with the given opCode and payload.
func (c *Client) Write(opCode int, payload []byte) error {
	c.Ws.SetWriteDeadline(time.Now().Add(writeWait))
	return c.Ws.WriteMessage(opCode, payload)
}

// проталкивание сообщений от hub до клиента.
func (c *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)

	defer func() {
		ticker.Stop()
		c.Ws.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.Write(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.Write(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.Write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}
