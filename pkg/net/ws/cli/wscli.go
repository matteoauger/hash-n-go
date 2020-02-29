package cli

import (
	"fmt"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

func main() {
	var url = url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/"}

	ws, _, err := websocket.DefaultDialer.Dial(url.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	ws.WriteMessage(websocket.TextMessage, []byte("Hello world"))

	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("read: ", err)
			return
		}

		fmt.Println("received :", string(message))
	}
}
