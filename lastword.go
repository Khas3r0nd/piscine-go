package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}
	lastIndex := 0
	isFound := false
	sentence := args[0]
	for i := len(sentence) - 1; i >= 0; i-- {
		if sentence[i] != ' ' && !isFound {
			lastIndex = i
			isFound = true
		}
		if sentence[i] == ' ' && isFound {
			fmt.Println(sentence[i+1 : lastIndex+1])
			return
		}
	}
}
