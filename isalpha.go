package piscine

func IsAlpha(s string) bool {
	aString := []rune(s)
	for _, letter := range aString {
		if (letter < 'A' || letter > 'Z') && (letter < 'a' || letter > 'z') && (letter < '0' || letter > '9') {
			return false
		}
	}
	return true
}
