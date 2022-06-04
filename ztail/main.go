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

func main() {
	isErr := false
	if os.Args[1] == "-c" {
		for i := 3; i < len(os.Args); i++ {
			file, err := os.ReadFile(os.Args[i])
			if err != nil {
				fmt.Println("open " + os.Args[i] + ": no such file or directory")
				fmt.Println()
				isErr = true
				continue
			}
			num := atoi(os.Args[2])
			if len(file)-num > 0 {
				fmt.Println("==> " + os.Args[i] + " <==")
				fmt.Print(string(file)[len(file)-num+1:])
			} else {
				fmt.Println("==> " + os.Args[i] + " <==")
				fmt.Print(string(file)[:])
			}
		}
		if isErr {
			os.Exit(1)
		}
	}
}
