package main

import (
	"fmt"
	"os"
)

func Rot13(s string) string {
	answer := ""
	for _, letter := range s {
		if letter >= 'a' && letter <= 'z' {
			if letter+13 <= 'z' {
				answer += string(letter + 13)
			} else {
				answer += string(96 + 13 - ('z' - letter))
			}
		} else if letter >= 'A' && letter <= 'Z' {
			if letter+13 <= 'Z' {
				answer += string(letter + 13)
			} else {
				answer += string(64 + 13 - ('Z' - letter))
			}
		} else {
			answer += string(letter)
		}
	}
	return answer
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}
	fmt.Println(Rot13(args[0]))
}
