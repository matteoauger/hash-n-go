package main

import (
	"fmt"
	"os"
	"strconv"

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

	if len(args) > 3 {
		var err error
		nWorkers, err = strconv.Atoi(args[2])

		if err != nil {
			fmt.Printf("arg error : Worker count should be an integer")
			os.Exit(1)
		}
	} else {
		nWorkers = swarm.GetNodeCount()
	}

	srv.Start("localhost:8080", connHandler, msgHandler)
	schSpace = scal.ScaleWorkers(nWorkers, nDigits)
}

func connHandler(c *websocket.Conn) {
	fmt.Println("Connected")
}

func msgHandler(message string) {
	fmt.Println(message)
}
