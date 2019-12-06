package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func md5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func match(message string, target string) bool {
	sum := md5Hash(message)

	return strings.Compare(target, sum) == 0
}

func main() {
	fmt.Print("I'm the worker.")
}
