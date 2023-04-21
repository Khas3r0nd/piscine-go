package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args

	fileName := arguments[0]

	runeFileName := []rune(fileName)

	for i := range runeFileName {
		if i >= 2 {
			z01.PrintRune(runeFileName[i])
		}
	}
	z01.PrintRune('\n')
}
