package piscine

func Atoi(s string) int {
	if s == "" {
		return 0
	}
	isFirstMinusSign := true
	minusSignCounter := 0
	plusSignCounter := 0
	answer := 0
	aString := []rune(s)
	for i, letter := range aString {
		if letter == '+' || letter == '-' {
			if letter == '+' {
				if i != 0 {
					return 0
				} else {
					plusSignCounter++
					if plusSignCounter > 1 {
						return 0
					} else {
						continue
					}
				}
			}
			if letter == '-' {
				if i != 0 {
					return 0
				} else {
					minusSignCounter++
					if minusSignCounter > 1 {
						return 0
					} else {
						if isFirstMinusSign && letter == '-' {
							continue
						} else if isFirstMinusSign && letter != '-' {
							answer = answer*(-10) + int(letter-'0')
							isFirstMinusSign = false
						} else {
							answer = answer*10 + int(letter-'0')
						}
					}
				}
			}
		} else {
			if letter >= '0' && letter <= '9' {
				answer = answer*10 + int(letter-'0')
			} else {
				return 0
			}
		}
	}
	if aString[0] == '-' {
		return -answer
	}
	return answer
}
