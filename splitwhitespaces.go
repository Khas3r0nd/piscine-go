package piscine

func SplitWhiteSpaces(s string) []string {
	var answer []string
	tempWord := ""
	for i, value := range s {
		if value != ' ' && i != len(s)-1 {
			tempWord += string(value)
		} else if i == len(s)-1 {
			tempWord += string(value)
			if tempWord != "" {
				answer = append(answer, tempWord)
				return answer
			}
		} else {
			if tempWord != "" {
				answer = append(answer, tempWord)
				tempWord = ""
			}
		}
	}
	return answer
}
