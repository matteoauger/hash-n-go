package srv

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"gitlab.com/hacheurs/hash-n-go/pkg/net/ws"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var connHandler ws.ConnHandler
var msgHandler ws.MsgHandler

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func reader(conn *websocket.Conn) {
	for {
		messageType, msg, err := conn.ReadMessage()

		if err != nil {
			log.Println(err)
			return
		}

		msgHandler(string(msg))

		// echo back the message

		if err := conn.WriteMessage(messageType, msg); err != nil {
			log.Println(err)
			return
		}
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true } // allow any connection

	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
	}

	connHandler(ws)
	reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", wsEndpoint)
}

// Start starts the websocket API
func Start(addr string, connectionHandler ws.ConnHandler, messageHandler ws.MsgHandler) {
	connHandler = connectionHandler
	msgHandler = messageHandler
	fmt.Println("Go websockets")
	setupRoutes()
	log.Fatal(http.ListenAndServe(addr, nil))
}
