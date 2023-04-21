package piscine

func AlphaCount(s string) int {
	counter := 0
	aString := []rune(s)
	for _, letter := range aString {
		if (letter >= 'a' && letter <= 'z') || (letter >= 'A' && letter <= 'Z') {
			counter++
		}
	}
	return counter
}
