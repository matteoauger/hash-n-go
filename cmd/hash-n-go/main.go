package main

import (
	"fmt"
	"os"
	"strconv"

	"gitlab.com/hacheurs/hash-n-go/pkg/sys/swarm"
    "gitlab.com/hacheurs/hash-n-go/pkg/scal"
)

// TODO instructions suivantes :
// Args : ["hash", "websocket-URI"]
// Récupérer n machines slaves
// Subdiviser l'espace de recherche avec n
// Distribuer le travail aux workers
// On return : envoi websocket au serveur node


func main() {
	args := os.Args

	if len(args) < 3 || len(args) > 4 {
		hint := "<hash> <websocket-URI> <optionnal: worker count>"
		fmt.Printf("USAGE : %s %s \n", args[0], hint)
		os.Exit(1)
	}

	//hash         := args[1]
	//websocketUri := args[2]

	// getting the worker count either from args or automatic
	var nWorkers int
	if len(args) > 3 {
		var err error
		nWorkers, err = strconv.Atoi(args[3])

		if err != nil {
			fmt.Printf("arg error : Worker count should be an integer")
			os.Exit(1)
		}
	} else {
		nWorkers = swarm.GetNodeCount()
	}

    scal.ScaleWorkers(nWorkers, 6)
	//fmt.Printf("%s %s", hash, websocketUri)
}
