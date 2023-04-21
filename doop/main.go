package main

import (
	"fmt"
	"os"
)

func Atoi(s string) int {
	if s == "" {
		return 0
	}
	isFirstMinusSign := true
	minusSignCounter := 0
	plusSignCounter := 0
	answer := 0
	aString := []rune(s)
	for i, letter := range aString {
		if letter == '+' || letter == '-' {
			if letter == '+' {
				if i != 0 {
					return 0
				} else {
					plusSignCounter++
					if plusSignCounter > 1 {
						return 0
					} else {
						continue
					}
				}
			}
			if letter == '-' {
				if i != 0 {
					return 0
				} else {
					minusSignCounter++
					if minusSignCounter > 1 {
						return 0
					} else {
						if isFirstMinusSign && letter == '-' {
							continue
						} else if isFirstMinusSign && letter != '-' {
							answer = answer*(-10) + int(letter-'0')
							isFirstMinusSign = false
						} else {
							answer = answer*10 + int(letter-'0')
						}
					}
				}
			}
		} else {
			if letter >= '0' && letter <= '9' {
				answer = answer*10 + int(letter-'0')
			} else {
				return 0
			}
		}
	}
	if aString[0] == '-' {
		return -answer
	}
	return answer
}

func Addition(a, b int) int {
	return a + b
}

func Substract(a, b int) int {
	return a - b
}

func Divide(a, b int) int {
	return a / b
}

func Multiplication(a, b int) int {
	return a * b
}

func Modulo(a, b int) int {
	return a % b
}

func isValid(s string) bool {
	for _, value := range s {
		if len(s) != 1 && s[0] == '0' {
			return false
		}
		if value < '0' || value > '9' {
			return false
		}
	}
	return true
}

func main() {
	arguments := os.Args[1:]
	if len(arguments) != 3 {
		return
	} else if arguments[1] == "*" || arguments[1] == "+" || arguments[1] == "-" || arguments[1] == "/" || arguments[1] == "%" {
		if isValid(arguments[0]) && isValid(arguments[2]) {
			switch arguments[1] {
			case "+":
				fmt.Print(Addition(Atoi(arguments[0]), Atoi(arguments[2])))
				break
			case "-":
				fmt.Print(Substract(Atoi(arguments[0]), Atoi(arguments[2])))
				break
			case "*":
				fmt.Print("hello")
				break
			case "/":
				fmt.Print(Divide(Atoi(arguments[0]), Atoi(arguments[2])))
				break
			case "%":
				fmt.Print(Modulo(Atoi(arguments[0]), Atoi(arguments[2])))
				break
			}
		}
	} else {
		return
	}
}
