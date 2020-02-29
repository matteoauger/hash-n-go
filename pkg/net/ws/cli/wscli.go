package cli

import (
	"log"

	"github.com/gorilla/websocket"
	"gitlab.com/hacheurs/hash-n-go/pkg/net/ws"
)

// Connect connects to the websocket
func Connect(url string, connHandler ws.ConnHandler) {
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	connHandler(ws)
}
