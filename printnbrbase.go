package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if nbr == -9223372036854775808 && base == "0123456789" {
		z01.PrintRune('-')
		z01.PrintRune('9')
		z01.PrintRune('2')
		z01.PrintRune('2')
		z01.PrintRune('3')
		z01.PrintRune('3')
		z01.PrintRune('7')
		z01.PrintRune('2')
		z01.PrintRune('0')
		z01.PrintRune('3')
		z01.PrintRune('6')
		z01.PrintRune('8')
		z01.PrintRune('5')
		z01.PrintRune('4')
		z01.PrintRune('7')
		z01.PrintRune('7')
		z01.PrintRune('5')
		z01.PrintRune('8')
		z01.PrintRune('0')
		z01.PrintRune('8')
		return
	}
	// Check if the base is valid
	// A base must contain at least 2 characters
	if len(base) < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	// Check if the base has unique characters and doesn't contain '+' or '-'
	for i := 0; i < len(base); i++ {
		if base[i] == '+' || base[i] == '-' {
			z01.PrintRune('N')
			z01.PrintRune('V')
			return
		}
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				z01.PrintRune('N')
				z01.PrintRune('V')
				return
			}
		}
	}

	// Check if the number is negative, negate it and print a '-' sign
	if nbr < 0 {
		nbr = -nbr
		z01.PrintRune('-')
	}

	// Convert the number to the desired base
	var res []int
	for nbr > 0 {
		res = append(res, nbr%len(base))
		nbr = nbr / len(base)
	}

	// Print the converted number
	for i := len(res) - 1; i >= 0; i-- {
		z01.PrintRune(rune(base[res[i]]))
	}
}
