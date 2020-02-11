package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os/exec"
)

// TODO instructions suivantes : 
// Args : ["hash", "websocket-URI"]
// Récupérer n machines slaves
// Subdiviser l'espace de recherche avec n 
// Distribuer le travail aux workers
// On return : envoi websocket au serveur node

func main() {
	worker := "gitlab.com/hacheurs/hash-n-go/cmd/hash-n-go-worker"
	start := "000"
	end := "999"
	target := "111"
	hash := md5Hash(target)
	cmd := exec.Command("go", "run", worker, start, end, hash)
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error %v\n", err)
	} else {
		fmt.Printf("Output: %s\n", out)
	}
}

func md5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
