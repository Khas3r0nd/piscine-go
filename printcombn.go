package piscine

import (
	"github.com/01-edu/z01"
)

func printComb(n int, comb []int, index int, firstime *bool) {
	if index == n {
		if !(*firstime) {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
		*firstime = false
		for _, i := range comb {
			z01.PrintRune(rune(i) + '0')
		}
		return
	}
	start := 0
	if index == 0 {
		start = comb[index]
	} else {
		start = comb[index-1] + 1
	}
	for i := start; i < 10; i++ {
		comb[index] = i
		printComb(n, comb, index+1, firstime)
	}
}

func PrintCombN(n int) {
	firstime := true
	if n == 0 {
		return
	}
	comb := make([]int, n)
	for i := range comb {
		comb[i] = i
	}
	printComb(n, comb, 0, &firstime)
	z01.PrintRune('\n')
}
