package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for i := 9; i >= 0; i-- {
		for j := 9; j >= 0; j-- {
			for k := 9; k >= 0; k-- {
				for l := 9; l >= 0; l-- {
					if i*10+j > k*10+l {
						z01.PrintRune(rune(i + 48))
						z01.PrintRune(rune(j + 48))
						z01.PrintRune(' ')
						z01.PrintRune(rune(k + 48))
						z01.PrintRune(rune(l + 48))
						if i == 0 && j == 1 && k == 0 && l == 0 {
							break
						}
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
}
