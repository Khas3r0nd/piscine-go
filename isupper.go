package piscine

func IsUpper(s string) bool {
	aString := []rune(s)
	for _, letter := range aString {
		if letter < 'A' || letter > 'Z' {
			return false
		}
	}
	return true
}
