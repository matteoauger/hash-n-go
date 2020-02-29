package ws

import "github.com/gorilla/websocket"

// ConnHandler websocket connection listener function
type ConnHandler func(*websocket.Conn)

// MsgHandler websocket message listener function
type MsgHandler func(string)
