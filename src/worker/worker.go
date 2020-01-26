package worker

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math"
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
	size := len(end)
	current := make([]int, size)
	password := ""
	tab := getCharacterRange()
	cpt := 0
	//TODO: voir pour le problème qu'il ne prendra jamais en compte le premier element du character range
	//Il faut trouver une autre méthode pour construire le password
	for cpt < int(math.Pow(float64(len(tab)), float64(size))) {
		i := size - 1
		password = ""
		toTake := false
		for i >= 0 {
			if current[i] != 0 {
				toTake = true
			}
			if toTake {
				password += string(tab[current[i]])
			}
			i--
		}
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
func increment(arr *[]int, length int, i int, limit int) {
	if i < 0 || i >= length {
		return
	}
	(*arr)[i]++
	if (*arr)[i]+1 == limit+1 {
		(*arr)[i] = 0
		increment(arr, length, i+1, limit)
	}
}
