package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	for i := range arguments {
		if i > 0 {
			for _, value := range arguments[i] {
				z01.PrintRune(value)
			}
			z01.PrintRune('\n')
		}
	}
}
