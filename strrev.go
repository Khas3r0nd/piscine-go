package piscine

func StrRev(s string) string {
	answer := ""
	aString := []rune(s)
	for i := len(aString) - 1; i >= 0; i-- {
		answer = answer + string(aString[i])
	}
	return answer
}
