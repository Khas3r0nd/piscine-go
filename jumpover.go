package piscine

func JumpOver(str string) string {
	var answer string
	for i, letter := range str {
		if i != 0 && i%3 == 2 {
			answer += string(letter)
		}
	}
	answer += "\n"
	return answer
}
