package main

import (
	"fmt"
    "bytes"
    "io"
    "os"
	"os/exec"
)

// TODO instructions suivantes : 
// Args : ["hash", "websocket-URI"]
// Récupérer n machines slaves
// Subdiviser l'espace de recherche avec n 
// Distribuer le travail aux workers
// On return : envoi websocket au serveur node

// 26 * 2 + 10 = 62 * 6 = 372 caractères 

func main() {
    args := os.Args

    if len(args) != 3 {
        fmt.Printf("USAGE : %s <hash> <websocket-URI>\n", args[0])
        os.Exit(1)
    }

    //hash         := args[1]
    //websocketUri := args[2]
    getNodeCount();

    //fmt.Printf("%s %s", hash, websocketUri)
}

func getNodeCount() {
        //create command
        catCmd := exec.Command("docker", "node", "ls")
        wcCmd := exec.Command( "wc", "-l" )

        //make a pipe
        reader, writer := io.Pipe()
        var buf bytes.Buffer

        //set the output of "cat" command to pipe writer
        catCmd.Stdout = writer
        //set the input of the "wc" command pipe reader

        wcCmd.Stdin = reader

        //cache the output of "wc" to memory
        wcCmd.Stdout = &buf

        //start to execute "cat" command
        catCmd.Start()

        //start to execute "wc" command
        wcCmd.Start()

        //waiting for "cat" command complete and close the writer
        catCmd.Wait()
        writer.Close()

        //waiting for the "wc" command complete and close the reader
        wcCmd.Wait()
        reader.Close()
        //copy the buf to the standard output
        io.Copy( os.Stdout, &buf )

        fmt.Println(string(buf.Bytes()))
}

//func scaleWorkers() {
    //var nbAvailableSlaves: int64 = 6
    //var workAmmount : int64 = 372 / nbAvailableSlaves
//}

