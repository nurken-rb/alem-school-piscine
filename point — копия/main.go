package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func printStr(str string) {
	r := []rune(str)

	for i := 0; i < len(r); i++ {
		z01.PrintRune(r[i])
	}
}

func main() {
	points := &point{}

	setPoint(points)

	printStr("x = ")
	z01.PrintRune(52)
	z01.PrintRune(50)
	printStr(", y = ")
	z01.PrintRune(50)
	z01.PrintRune(49)
	z01.PrintRune('\n')
}
