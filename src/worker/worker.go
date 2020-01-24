package worker

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func getCharacterRange() []int {
	tab := []int{}

	// numbers
	for i := 0x30; i <= 0x39; i++ {
		tab = append(tab, i)
	}

	// uppercase characters
	for i := 0x41; i <= 0x5a; i++ {
		tab = append(tab, i)
	}

	// lowercase characters
	for i := 0x61; i <= 0x7a; i++ {
		tab = append(tab, i)
	}

	return tab
}

func Md5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

// Match explique
func Match(message string, target string) bool {
	sum := Md5Hash(message)

	if strings.Compare(target, sum) == 0 {
		return true
	}

	return false
}

//MainLoop hdihiz
func MainLoop(start string, end string, target string) string {
	current := start
	tab := getCharacterRange()
	cpt := 0
	for cpt < len(tab) {
		current = string(tab[cpt])
		fmt.Println(current)
		if Match(current, target) {
			fmt.Println("found!")
			return current
		}
		cpt += 1
	}
	return ""
}

// Test csiojfise
func Test() {
	tab := getCharacterRange()

	for i := 0; i < len(tab); i++ {
		fmt.Printf("%c ", tab[i])
	}
}

// Increment the given index of the clock array.
func increment(arr []int, length int, i int, limit int) {
	if i < 0 || i >= length {
		return
	}
	arr[i]++
	if arr[i] == 0 {
		return
	}
	if arr[i]%limit == 0 {
		arr[i] = 0
		increment(arr, length, i+1, limit)
	}
}
