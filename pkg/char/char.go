package char

// CreateAlphabet creates the alphabet for the hash source possibilities
func CreateAlphabet() []rune {
	// Init char. arra
	chars := []rune{}

    // append numbers
	for i := 0x30; i <= 0x39; i++ {
		chars = append(chars, rune(i))
	}

	// append uppercase characters
	for i := 0x41; i <= 0x5a; i++ {
		chars = append(chars, rune(i))
	}

	// append lowercase characters
	for i := 0x61; i <= 0x7a; i++ {
		chars = append(chars, rune(i))
	}

    return chars
}

