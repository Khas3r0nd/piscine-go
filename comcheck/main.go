package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args[1:]
	for _, value := range arguments {
		if value == "01" || value == "galaxy" || value == "galaxy 01" {
			fmt.Println("Alert!!!")
			return
		}
	}
}
