package piscine

func IsLower(s string) bool {
	aString := []rune(s)
	for _, letter := range aString {
		if letter < 'a' || letter > 'z' {
			return false
		}
	}
	return true
}
