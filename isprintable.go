package piscine

func IsPrintable(s string) bool {
	aString := []rune(s)
	for _, letter := range aString {
		if letter < 32 || letter > 127 {
			return false
		}
	}
	return true
}
