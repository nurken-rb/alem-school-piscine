package main

import (
	"fmt"
	"os"
)

func atoi(s string) int {
	i := 0
	result := 0
	sign := 1
	if s[0] == '-' {
		sign = -1
		i++
	}
	for i < len(s) {
		if s[i] < '0' || s[i] > '9' {
			return result
		} else {
			result *= 10
			result += int(s[i]) - '0'
		}
		i++
	}
	return result * sign
}

func isNumeric(s string) bool {
	for i := 0; i < len(s); i++ {
		if (48 > s[i] || s[i] > 57) && s[i] != '-' {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) == 4 {
		if !isNumeric(os.Args[1]) || !isNumeric(os.Args[3]) {
			return
		}
		if len(os.Args[1]) > 18 || len(os.Args[3]) > 18 {
			return
		}

		a := atoi(os.Args[1])
		b := atoi(os.Args[3])

		switch os.Args[2] {
		case "+":
			fmt.Println(a + b)
		case "-":
			fmt.Println(a - b)
		case "*":
			fmt.Println(a * b)
		case "/":
			if b == 0 {
				fmt.Println("No division by 0")
				return
			}
			fmt.Println(a / b)
		case "%":
			if b == 0 {
				fmt.Println("No modulo by 0")
				return
			}
			fmt.Println(a % b)
		}
	}
}
