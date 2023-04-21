package piscine

func IsAlphanumeric(char rune) bool {
	if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
		return true
	}
	return false
}

func Capitalize(s string) string {
	aString := []rune(s)

	first := true

	for i := range aString {
		if IsAlphanumeric(aString[i]) && first {
			if aString[i] >= 'a' && aString[i] <= 'z' {
				aString[i] -= 32
			}
			first = false
		} else if aString[i] >= 'A' && aString[i] <= 'Z' {
			aString[i] += 32
		} else if IsAlphanumeric(aString[i]) == false {
			first = true
		}
	}
	return string(aString)
}
