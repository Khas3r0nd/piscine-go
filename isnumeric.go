package piscine

func IsNumeric(s string) bool {
	aString := []rune(s)
	for _, letter := range aString {
		if letter < '0' || letter > '9' {
			return false
		}
	}
	return true
}
