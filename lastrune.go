package main

import "fmt"

func LastRune(s string) rune {
	aString := []rune(s)
	return aString[len(aString)-1]
}

func main() {
	fmt.Printf("%c", LastRune("Hello"))
}
