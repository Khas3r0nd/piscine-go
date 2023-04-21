package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x, y int
}

func convertToIntegerArray(number int) []int {
	var finalAnswer []int
	var numbers []int
	for number > 0 {
		numbers = append(numbers, number%10)
		number = number / 10
	}
	for i := len(numbers) - 1; i >= 0; i-- {
		finalAnswer = append(finalAnswer, numbers[i])
	}
	return finalAnswer
}

func printNums(numbers []int) {
	for _, value := range numbers {
		z01.PrintRune(rune(value) + 48)
	}
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}
	setPoint(points)

	numbersX := convertToIntegerArray(points.x)
	numbersY := convertToIntegerArray(points.y)
	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	printNums(numbersX)
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	printNums(numbersY)
	z01.PrintRune('\n')
}
