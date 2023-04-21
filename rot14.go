package piscine

func Rot14(s string) string {
	answer := ""
	for _, letter := range s {
		if letter >= 'a' && letter <= 'z' {
			if letter+14 <= 'z' {
				answer += string(letter + 14)
			} else {
				answer += string(96 + 14 - ('z' - letter))
			}
		} else if letter >= 'A' && letter <= 'Z' {
			if letter+14 <= 'Z' {
				answer += string(letter + 14)
			} else {
				answer += string(64 + 14 - ('Z' - letter))
			}
		} else {
			answer += string(letter)
		}
	}
	return answer
}
