package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	aString := []rune(s)
	for _, letter := range aString {
		z01.PrintRune(letter)
	}
}
