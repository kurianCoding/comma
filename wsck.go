package main

import (
	"flag"
	"github.com/gorilla/websocket"
	"net/url"
)

var addr = flag.String("addr", "localhost:12345", "http service address")

// this function allows transmitting of output
// and command as chat message strings
type CommandChat struct {
	Conn *websocket.Conn
}

func ChatConnect() (c *CommandChat, err error) {
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/echo"}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	c = &CommandChat{Conn: conn}
	return
}

func (c *CommandChat) Write(input []byte) (n int, err error) {
	err = c.Conn.WriteMessage(websocket.TextMessage, input)
	return
}
func (c *CommandChat) Rx(resp []byte) (err error) {
	return
}
func (c *CommandChat) Close() (err error) { return }
