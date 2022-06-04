package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i := len(os.Args) - 1; i > 0; i-- {
		for k := 0; k < len(os.Args[i]); k++ {
			r := []rune(os.Args[i])
			z01.PrintRune(r[k])
		}
		z01.PrintRune('\n')
	}
}
