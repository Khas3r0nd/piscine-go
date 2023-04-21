package piscine

func NRune(s string, n int) rune {
	aString := []rune(s)
	if n > 0 && n <= len(s) {
		return aString[n-1]
	}
	return (rune(0))
}
