package piscine

func ToUpper(s string) string {
	aString := []rune(s)

	for i := range aString {
		if aString[i] >= 97 && aString[i] <= 122 {
			aString[i] -= 32
		}
	}
	return string(aString)
}
