package piscine

import "github.com/01-edu/z01"

func horizontalE1(x int) {
	for i := 1; i <= x; i++ {
		if i == 1 {
			z01.PrintRune('A')
		} else if i == x {
			z01.PrintRune('C')
		} else {
			z01.PrintRune('B')
		}
	}
	z01.PrintRune('\n')
}

func horizontalE2(x int) {
	for i := 1; i <= x; i++ {
		if i == 1 {
			z01.PrintRune('C')
		} else if i == x {
			z01.PrintRune('A')
		} else {
			z01.PrintRune('B')
		}
	}
	z01.PrintRune('\n')
}

func verticalE(y int) {
	for i := 1; i <= y; i++ {
		if i == 1 || i == y {
			z01.PrintRune('B')
		} else {
			z01.PrintRune(32)
		}
	}
	z01.PrintRune('\n')
}

func QuadE(x, y int) {
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
			horizontalE1(x)
		} else if i == y {
			horizontalE2(x)
		} else {
			verticalE(x)
		}
	}
	z01.PrintRune('\n')
}
