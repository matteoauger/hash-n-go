package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"encoding/json"

	"github.com/gorilla/websocket"

	"gitlab.com/hacheurs/hash-n-go/pkg/net/ws/srv"
	"gitlab.com/hacheurs/hash-n-go/pkg/scal"
	"gitlab.com/hacheurs/hash-n-go/pkg/sys/swarm"
)

// TODO instructions suivantes :
// Distribuer le travail aux workers
// On return : envoi websocket au serveur node

// nDigits Max number of digits for the search space
const nDigits int = 6

var hash string
var schSpace []scal.SearchSpace
var cpt = -1

func main() {
	args := os.Args

	if len(args) < 2 || len(args) > 3 {
		hint := "<hash> <optionnal: worker count>"
		fmt.Printf("USAGE : %s %s \n", args[0], hint)
		os.Exit(1)
	}

	hash = args[1]

	// getting the worker count either from args or automatically
	var nWorkers int

	if len(args) > 2 {
		var err error
		nWorkers, err = strconv.Atoi(args[2])

		if err != nil {
			fmt.Printf("arg error : Worker count should be an integer")
			os.Exit(1)
		}
	} else {
		nWorkers = swarm.GetNodeCount()
	}

	schSpace = scal.ScaleWorkers(nWorkers, nDigits, hash)
	srv.Start("localhost:8080", connHandler)
}

func connHandler(c *websocket.Conn) {
	fmt.Println("Connected")
	cpt+=1

	json, err := json.Marshal(schSpace[cpt])
	if err != nil {
		log.Println(err)
		return
	}
	
	if err := c.WriteMessage(websocket.TextMessage, []byte(json)); err != nil {
		log.Println(err)
		return
	}
	
	// echo the password
	
	_, msg, err := c.ReadMessage()

	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(msg))
}
