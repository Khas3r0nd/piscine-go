package piscine

func BasicAtoi2(s string) int {
	answer := 0
	aString := []rune(s)
	for _, letter := range aString {
		if letter < '0' || letter > '9' {
			answer := 0
			return answer
		}
		answer = answer*10 + (int(letter) - 48)
	}
	return answer
}
