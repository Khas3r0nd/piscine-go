package main

import "fmt"

func StrLen(s string) int {
	aString := []rune(s)
	return len(aString)
}

func main() {
	fmt.Println(StrLen("Hello world!"))
}
