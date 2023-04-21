package main

import (
	"os"

	"github.com/01-edu/z01"
)

func BasicAtoi(s string) int {
	answer := 0
	aString := []rune(s)
	for _, letter := range aString {
		answer = answer*10 + (int(letter) - 48)
	}
	return answer
}

func main() {
	arguments := os.Args
	if len(arguments) < 2 {
		return
	}
	isUpper := false
	if arguments[1] == "--upper" {
		isUpper = true
	}
	for i := 1; i < len(arguments); i++ {
		letter := ""
		for _, value := range arguments[i] {
			letter += string(value)
		}
		if BasicAtoi(letter) >= 1 && BasicAtoi(letter) <= 26 && isUpper {
			z01.PrintRune(rune(BasicAtoi(letter) + 64))
		} else if BasicAtoi(letter) >= 1 && BasicAtoi(letter) <= 26 && isUpper == false {
			z01.PrintRune(rune(BasicAtoi(letter) + 96))
		} else if i == 1 && isUpper {
		} else {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
