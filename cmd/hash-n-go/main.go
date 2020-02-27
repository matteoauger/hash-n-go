package main

import (
    "fmt"
    "os"
    "strconv"

    "gitlab.com/hacheurs/hash-n-go/pkg/sys/swarm"
    "gitlab.com/hacheurs/hash-n-go/pkg/scal"
)

// TODO instructions suivantes :
// Distribuer le travail aux workers
// On return : envoi websocket au serveur node

// nDigits Max number of digits for the search space
const nDigits int = 6

func main() {
    args := os.Args

    if len(args) < 3 || len(args) > 4 {
        hint := "<hash> <websocket-URI> <optionnal: worker count>"
        fmt.Printf("USAGE : %s %s \n", args[0], hint)
        os.Exit(1)
    }

    //hash         := args[1]
    //websocketUri := args[2]

    // getting the worker count either from args or automatically
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

    schSpace := scal.ScaleWorkers(nWorkers, nDigits)
    fmt.Println(schSpace)
}

