package main

import (
	"fmt"

	"./worker"
)

func main() {
	testStr := "t"
	matched := worker.MainLoop("", "aaaaaa", worker.Md5Hash(testStr))
	fmt.Println(matched, testStr)
}
