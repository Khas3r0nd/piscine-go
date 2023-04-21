package piscine

func ToLower(s string) string {
	aString := []rune(s)

	for i := range aString {
		if aString[i] >= 65 && aString[i] <= 90 {
			aString[i] += 32
		}
	}
	return string(aString)
}
