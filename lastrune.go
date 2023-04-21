package piscine

func LastRune(s string) rune {
	aString := []rune(s)
	return aString[len(s)-1]
}
