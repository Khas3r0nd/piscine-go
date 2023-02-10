package main

import "fmt"

func FirstRune(s string) rune {
	aString := []rune(s)
	return aString[0]
}

func main() {
	fmt.Printf("%c", FirstRune("Hello"))
}
