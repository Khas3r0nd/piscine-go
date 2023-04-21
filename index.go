package piscine

func Index(s string, toFind string) int {
	givenString := []rune(s)
	stringToFind := []rune(toFind)
	t := 0

	for i := 0; i < len(givenString); i++ {
		j := 0
		t = i
		for j < len(stringToFind) {
			if t < len(givenString) && givenString[t] == stringToFind[j] {
				j++
				t++
			} else {
				break
			}
		}
		if j == len(stringToFind) {
			return i
		}
	}
	return -1
}
