package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/gorilla/websocket"

	"gitlab.com/hacheurs/hash-n-go/pkg/net/ws/srv"
	"gitlab.com/hacheurs/hash-n-go/pkg/scal"
	"gitlab.com/hacheurs/hash-n-go/pkg/sys/swarm"
)

// nDigits Max number of digits for the search space
const nDigits int = 6
const flagHash string = "hash"
const workerHash string = "workers"

var hash string
var schSpace []scal.SearchSpace
var cpt = -1

var mutex = &sync.Mutex{}

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
	go srv.Start("localhost:8080", connHandler)
	swarm.InitSwarm("ws://localhost:8080")
}

func connHandler(c *websocket.Conn) {
	fmt.Println("Connected")

	// locking the counter modification
	mutex.Lock()
	cpt++
	var currentSchSpace = schSpace[cpt]
	mutex.Unlock()

	json, err := json.Marshal(currentSchSpace)
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

	fmt.Println("FOUND : " + string(msg))
	swarm.ClearSwarm()
	os.Exit(0)
}
