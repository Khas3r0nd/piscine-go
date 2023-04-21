package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) < 2 {
		bytes, _ := io.ReadAll(os.Stdin)
		os.Stdout.Write(bytes)
	}
	fileNames := os.Args[1:]
	for _, fileName := range fileNames {
		data, err := os.ReadFile(fileName)
		if err != nil {
			PrintStr("ERROR: ")
			PrintStr(err.Error())
			z01.PrintRune('\n')
			os.Exit(1)
			return
		}
		os.Stdout.Write(data)
	}
}

func PrintStr(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}
