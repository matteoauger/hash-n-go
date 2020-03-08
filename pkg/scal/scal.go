package scal

import (
	"math"

	"gitlab.com/hacheurs/hash-n-go/pkg/char"
)

// ScaleWorkload returns the scaled work for the given number of wokers & digits
func ScaleWorkload(nbWorkers int, nDigit int, hash string) []SearchSpace {
	chars := char.CreateAlphabet()
	base := len(chars)
	nChars := int(math.Pow(float64(base), float64(nDigit))) - 1
	workAmount := nChars / nbWorkers
	schSpaces := make([]SearchSpace, nbWorkers)

	for i := 0; i < nbWorkers; i++ {
		// calculate the begin string and end string for the given worker
		beginIdx := (workAmount * i) + 1 // begin index
		endIdx := workAmount * (i + 1)   // end index

		// compensating the poor precision of the integer division
		// as we can't cut exactly the workload
		if i == nbWorkers-1 && endIdx != nChars {
			endIdx = nChars
		}

		beginStr := convertBase(beginIdx, chars)
		endStr := convertBase(endIdx, chars)

		schSpaces[i] = SearchSpace{Begin: beginStr, End: endStr, Hash: hash}
	}

	return schSpaces
}

// convertBase converts the base of the given decimal number
func convertBase(nbr int, symbol []rune) string {
	newBase := len(symbol)
	res := ""

	n := nbr

	for n > 0 {
		res = string(symbol[n%newBase]) + res
		n = n / newBase
	}

	return res
}
