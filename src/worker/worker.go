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
	//TODO: Rendre taille égale à taille de end
	const size = 3
	var current [size]int
	password := ""
	tab := getCharacterRange()
	cpt := 0
	// TODO: Condition d'arret à changer
	for cpt < len(tab)*len(tab)*len(tab) {
		password = string(tab[current[0]]) + string(tab[current[1]]) + string(tab[current[2]])
		if Match(password, target) {
			fmt.Println("found!")
			return password
		}
		cpt += 1
		increment(&current, size, 0, len(tab))
	}
	return ""
}

// Increment the given index of the clock array.
// TODO: enlever la taille fixe de arr
func increment(arr *[3]int, length int, i int, limit int) {
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
