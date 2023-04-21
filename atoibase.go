package piscine

func AtoiBase(s string, base string) int {
	// Check if the base is valid
	if len(base) < 2 {
		// Return 0 if the base is invalid
		return 0
	}
	// Check if the base contains any invalid characters ('+' or '-')
	for i := 0; i < len(base); i++ {
		if base[i] == '+' || base[i] == '-' {
			// Return 0 if the base is invalid
			return 0
		}
		// Check if the base contains any duplicate characters
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				// Return 0 if the base is invalid
				return 0
			}
		}
	}
	result := 0
	bLen := len(base)
	// Iterate over each character in the string 's'
	for i := 0; i < len(s); i++ {
		found := false
		// Check if the current character is in the base
		for j := 0; j < bLen; j++ {
			if s[i] == base[j] {
				result = result*bLen + j
				found = true
				break
			}
		}
		if !found {
			// Return 0 if the string 's' contains a character that is not in the base
			return 0
		}
	}
	// Return the integer value of 's' in the given base
	return result
}
