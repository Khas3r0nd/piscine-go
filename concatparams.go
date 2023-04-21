package piscine

func ConcatParams(args []string) string {
	answer := ""
	for i := 0; i < len(args); i++ {
		if i != len(args)-1 {
			answer += args[i]
			answer += "\n"
		} else {
			answer += args[i]
		}
	}
	return answer
}
