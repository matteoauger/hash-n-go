package main

import (
	"fmt"

	"./worker"
)

func main() {
	testStr := "test"
	matched := worker.Match(testStr, "098f6bcd4621d373cade4e832627b4f6")
	fmt.Println(matched, testStr)
}
