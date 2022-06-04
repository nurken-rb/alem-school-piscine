package piscine

import "github.com/01-edu/z01"

func horizontalB1(x int) {
	for i := 1; i <= x; i++ {
		if i == 1 {
			z01.PrintRune('/')
		} else if i == x {
			z01.PrintRune('\\')
		} else {
			z01.PrintRune('*')
		}
	}
	z01.PrintRune('\n')
}

func horizontalB2(x int) {
	for i := 1; i <= x; i++ {
		if i == 1 {
			z01.PrintRune('\\')
		} else if i == x {
			z01.PrintRune('/')
		} else {
			z01.PrintRune('*')
		}
	}
	z01.PrintRune('\n')
}

func verticalB(y int) {
	for i := 1; i <= y; i++ {
		if i == 1 || i == y {
			z01.PrintRune('*')
		} else {
			z01.PrintRune(32)
		}
	}
	z01.PrintRune('\n')
}

func QuadB(x, y int) {
	if x == 0 || y == 0 {
		printStr("Rectangle with size 0!\n")
		return
	}
	if x < 0 || y < 0 {
		printStr("Size cannot be negative number!\n")
		return
	}
	for i := 1; i <= y; i++ {
		if i == 1 {
			horizontalB1(x)
		} else if i == y {
			horizontalB2(x)
		} else {
			verticalB(x)
		}
	}
	z01.PrintRune('\n')
}
