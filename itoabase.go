package main

import "fmt"

func main() {
	fmt.Println(ItoaBase(-1225, 15))
}

func ItoaBase(value, base int) string {
	res := ""
	var count []rune

	if value < 0 { //-125<0
		res += "-"
		value = -value
	}

	for value > 0 {
		if (value % base) >= 10 {
			count = append(count, rune(value%base+55))
		} else {
			count = append(count, rune(value%base+48))
		}
		value /= base
	}
	for i := len(count) - 1; i >= 0; i-- {
		res += string(count[i])
	}
	return res
}
