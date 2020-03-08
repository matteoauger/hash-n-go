package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"

	"encoding/json"
	"log"
	"strings"

	"gitlab.com/hacheurs/hash-n-go/pkg/char"
	"gitlab.com/hacheurs/hash-n-go/pkg/net/ws/cli"
	"gitlab.com/hacheurs/hash-n-go/pkg/scal"

	"github.com/gorilla/websocket"
)

// Constants

var chars []rune
var charsLen int
var revChars map[rune]int

// Main

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("USAGE: %s %s", os.Args[0], "<websocket URI>\n")
		os.Exit(1)
	}

	var wsURI = os.Args[1]
	cli.Connect(wsURI, connHandler)
}

// Initialisation

func init() {
	chars = char.CreateAlphabet()
	// Init reversed char. array
	revChars = make(map[rune]int)
	for i, r := range chars {
		revChars[r] = i
	}
	charsLen = len(chars)
}

// Utils

func md5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func search(start string, end string, hash string) string {
	var current = stringToRefs(start)
	var password = ""
	for password != end {
		password = string(refsToString(current))
		if match(password, hash) {
			return password
		}
		current = increment(current, 0)
	}
	return ""
}

func match(message string, target string) bool {
	sum := md5Hash(message)
	return strings.Compare(target, sum) == 0
}

func stringToRefs(str string) []int {
	var refs = make([]int, len(str))
	for i, r := range str {
		refs[i] = revChars[r]
	}
	return refs
}

func refsToString(refs []int) []rune {
	var runes = make([]rune, len(refs))
	for i, ref := range refs {
		runes[i] = chars[ref]
	}
	return runes
}

func increment(arr []int, i int) []int {
	if i < 0 {
		return arr
	}
	if i >= len(arr) {
		return append(arr, 0)
	}
	arr[i]++
	if arr[i] >= charsLen {
		arr[i] = 0
		return increment(arr, i+1)
	}
	return arr
}

func connHandler(conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read: ", err)
			return
		}

		var schSpace scal.SearchSpace

		json.Unmarshal(message, &schSpace)

		password := search(schSpace.Begin, schSpace.End, schSpace.Hash)

		conn.WriteMessage(websocket.TextMessage, []byte(password))
	}
}
