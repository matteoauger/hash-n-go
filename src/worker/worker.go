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

func md5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

// Match explique
func Match(message string, target string) bool {
	sum := md5Hash(message)

	if strings.Compare(target, sum) == 0 {
		return true
	}

	return false
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
