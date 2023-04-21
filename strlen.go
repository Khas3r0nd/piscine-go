package piscine

func StrLen(s string) int {
	aString := []rune(s)
	counter := 0
	for range aString {
		counter += 1
	}
	return counter
}
