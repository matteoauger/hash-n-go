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

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true } // allow any connection

	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
	}

	connHandler(ws)
	//reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", wsEndpoint)
}

// Start starts the websocket API
func Start(addr string, connectionHandler ws.ConnHandler) {
	connHandler = connectionHandler
	fmt.Println("Go websockets")
	setupRoutes()
	log.Fatal(http.ListenAndServe(addr, nil))
}
