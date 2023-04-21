package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	arguments := os.Args[1:]
	if len(arguments) > 1 {
		fmt.Println("Too many arguments")
	} else if len(arguments) == 0 {
		fmt.Println("File name missing")
	} else {
		content, _ := ioutil.ReadFile(arguments[0])
		fmt.Printf("%s", content)
	}
}
