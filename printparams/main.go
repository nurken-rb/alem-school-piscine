package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		for k := 0; k < len(os.Args[i]); k++ {
			r := []rune(os.Args[i])
			z01.PrintRune(r[k])
		}
		z01.PrintRune('\n')
	}
}
