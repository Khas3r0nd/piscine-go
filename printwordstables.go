package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for i := range a {
		for _, value := range a[i] {
			/*if value != ' ' {
				z01.PrintRune(value)
			} else {
				z01.PrintRune('\n')
			}*/
			z01.PrintRune(value)
		}
		z01.PrintRune('\n')
	}
}
