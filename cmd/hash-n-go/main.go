package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

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
const flagHash string = "hash"
const workerHash string = "workers"

var hash string
var schSpace []scal.SearchSpace
var cpt = -1

func main() {
	hashPtr := flag.String(flagHash, "", "MANDATORY : Hash to decrypt")
	nWorkersPtr := flag.Int(workerHash, 0, "Number of workers to assign")
	flag.Parse()

	hash := *hashPtr
	nWorkers := *nWorkersPtr

	// getting the worker count either from args or automatically
	if nWorkers <= 0 {
		nWorkers = swarm.GetNodeCount()
	}

	// Managing empty hash
	if hash == "" {
		fmt.Printf("Error : Missing hash\nUSAGE :\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// If at this point no workers are available, send an error and exit the program
	if nWorkers <= 0 {
		fmt.Println("Error : no workers available.")
		os.Exit(1)
	}

	// Scale the workload and start the websocket endpoint
	schSpace = scal.ScaleWorkload(nWorkers, nDigits, hash)
	srv.Start("localhost:8080", connHandler)
}

func connHandler(c *websocket.Conn) {
	fmt.Println("Connected")
	cpt++

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
