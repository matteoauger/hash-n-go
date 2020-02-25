package main

import (
	"fmt"
    "bytes"
    "io"
    "os"
	"os/exec"
    "strconv"
    "strings"
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
    nWorkers := getNodeCount()
    fmt.Println(nWorkers)
    //fmt.Printf("%s %s", hash, websocketUri)
}

func getNodeCount() int {
        //create command
        ncCmd := exec.Command("docker", "node", "ls")
        wcCmd := exec.Command("wc", "-l")

        //make a pipe
        reader, writer := io.Pipe()
        var buf bytes.Buffer

        //set the output of "node ls" command to pipe writer
        ncCmd.Stdout = writer
        //set the input of the "wc" command pipe reader

        wcCmd.Stdin = reader

        //cache the output of "wc" to memory
        wcCmd.Stdout = &buf

        //start to execute "node ls" command
        ncCmd.Start()

        //start to execute "wc" command
        wcCmd.Start()

        ncCmd.Wait()
        writer.Close()

        wcCmd.Wait()
        reader.Close()

        strStdout := string(buf.Bytes())
        firstLine := strings.Split(strStdout, "\n")[0]
        res, err := strconv.Atoi(firstLine)

        if err != nil {
            panic(err)
        }

        // return number of available workers
        // minus 1 bc ignoring header
        return res - 1
}

//func scaleWorkers() {
    //var nbAvailableSlaves: int64 = 6
    //var workAmmount : int64 = 372 / nbAvailableSlaves
//}

