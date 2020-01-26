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

func getCharacterMap(tab []int) map[string]int {
	result := make(map[string]int)
	for i, character := range tab {
		result[string(character)] = i
	}
	return result
}

func getSteps(startClock []int, endClock []int, sizeAlphabet int) int {
	i := 0
	baseS := 0
	for i < len(startClock) {
		baseS += int(math.Pow(float64(sizeAlphabet), float64(i))) + startClock[i]
		i += 1
	}
	i = 0
	baseE := 0
	for i < len(endClock) {
		baseE += int(math.Pow(float64(sizeAlphabet), float64(i))) + endClock[i]
		i += 1
	}
	return baseE - baseS
}

func Md5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

// Match explique
func Match(message string, target string) bool {
	sum := Md5Hash(message)

	return strings.Compare(target, sum) == 0
}

func main() {
	fmt.Print("I'm the worker.")
}

//MainLoop hdihiz
func MainLoop(start string, end string, target string) string {
	size := len(end)
	startLength := len(start)
	endLength := len(end)
	current := make([]int, size)
	startClock := make([]int, startLength)
	endClock := make([]int, endLength)
	password := ""
	tab := getCharacterRange()
	characterMap := getCharacterMap(tab)
	k := 0
	for k < startLength {
		startClock[k] = characterMap[string(start[k])]
		current[k] = startClock[k]
		k += 1
	}
	k = 0
	for k < endLength {
		endClock[k] = characterMap[string(end[k])]
		k += 1
	}
	steps := getSteps(startClock, endClock, len(tab))
	cpt := 0
	for cpt < steps {
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
