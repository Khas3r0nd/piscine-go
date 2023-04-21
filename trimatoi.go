package piscine

func TrimAtoi(s string) int {
	if s == "" {
		return 0
	}
	minusSignCounter := 0
	answer := 0
	aString := []rune(s)
	for _, letter := range aString {
		if letter == '-' {
			if answer == 0 {
				minusSignCounter++
			}
		} else {
			if letter >= '0' && letter <= '9' {
				answer = answer*10 + int(letter-'0')
			} else {
				continue
			}
		}
	}

	if minusSignCounter != 0 {
		return -answer
	}
	return answer
}
