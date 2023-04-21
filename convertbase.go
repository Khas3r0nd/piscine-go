package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	// Convert baseFrom and baseTo to slices of runes
	from := []rune(baseFrom)
	to := []rune(baseTo)
	baseFromLen := len(from)
	baseToLen := len(to)

	// Initialize decimal value
	var decValue int

	// Get the length of the number string
	nbrLen := len(nbr)

	// Convert the number string to decimal value
	for i := 0; i < nbrLen; i++ {
		for j := 0; j < baseFromLen; j++ {
			// Find the index of the current digit in the baseFrom
			if from[j] == rune(nbr[i]) {
				decValue = decValue*baseFromLen + j
				break
			}
		}
	}

	// If the decimal value is 0, return the first character in baseTo
	if decValue == 0 {
		return string(to[0])
	}

	// Initialize the result slice
	var res []rune

	// Convert decimal value to the target base
	for decValue > 0 {
		res = append(res, to[decValue%baseToLen])
		decValue /= baseToLen
	}

	// Reverse the result slice to get the correct order
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	// Return the result as a string
	return string(res)
}
