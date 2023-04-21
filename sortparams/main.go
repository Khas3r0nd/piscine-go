package main

import (
	"os"

	"github.com/01-edu/z01"
)

func BubbleSorting(arr []string) []string {
	size := len(arr)
	for i := 1; i < size; i++ {
		for j := 1; j < size-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	arguments := os.Args
	sortedArguments := BubbleSorting(arguments)
	for i := range sortedArguments {
		if i > 0 {
			for _, value := range sortedArguments[i] {
				z01.PrintRune(value)
			}
			z01.PrintRune('\n')
		}
	}
}
