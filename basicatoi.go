package piscine

func BasicAtoi(s string) int {
	answer := 0
	aString := []rune(s)
	for _, letter := range aString {
		answer = answer*10 + (int(letter) - 48)
	}
	return answer
}
