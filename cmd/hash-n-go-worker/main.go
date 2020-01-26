package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

// Constants

var chars []rune
var charsLen int
var revChars map[rune]int

// Initialisation

func init() {
	// Init char. arra
	chars = []rune{}
	// append numbers
	for i := 0x30; i <= 0x39; i++ {
		chars = append(chars, rune(i))
	}
	// append uppercase characters
	for i := 0x41; i <= 0x5a; i++ {
		chars = append(chars, rune(i))
	}
	// append lowercase characters
	for i := 0x61; i <= 0x7a; i++ {
		chars = append(chars, rune(i))
	}

	// Init reversed char. array
	revChars = make(map[rune]int)
	for i, r := range chars {
		revChars[r] = i
	}
	charsLen = len(chars)
}

// Main

func main() {
	var args = os.Args
	if len(args) < 4 {
		fmt.Fprintln(os.Stderr, "Usage:", args[0], "<start>", "<end>", "<hash>")
		os.Exit(1)
	}
	var start = os.Args[1]
	var end = os.Args[2]
	var hash = os.Args[3]
	var lStart = len(start)
	var lEnd = len(end)
	if lStart > lEnd {
		fmt.Fprintf(os.Stderr, "'%s' greater than '%s'\n", start, end)
		os.Exit(2)
	}
	var rStart = stringToRefs(start)
	var rEnd = stringToRefs(end)
	for i := 0; i < lStart; i++ {
		if rStart[i] < rEnd[i] {
			break
		}
		if rStart[i] > rEnd[i] {
			fmt.Fprintf(os.Stderr, "'%s' greater than '%s'\n", start, end)
			os.Exit(2)
		}
	}
	var pass = search(start, end, hash)
	fmt.Println(pass)
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
