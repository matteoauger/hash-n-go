package main

import (
	"fmt"

	"./worker"
)

func main() {
	matched, message := worker.Match("test", "098f6bcd4621d373cade4e832627b4f6")
	fmt.Println(matched, message)
}
