package piscine

func ReverseMenuIndex(menu []string) []string {
	answer := make([]string, len(menu))
	counter := 0
	for i := len(menu) - 1; i >= 0; i-- {
		answer[counter] = menu[i]
		counter++
	}
	return answer
}
