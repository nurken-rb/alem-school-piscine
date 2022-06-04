package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := '0'; i <= '6'; i++ {
		for y := i + 1; y <= '8'; y++ {
			for z := y + 1; z <= '9'; z++ {
				z01.PrintRune(i)
				z01.PrintRune(y)
				z01.PrintRune(z)
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('7')
	z01.PrintRune('8')
	z01.PrintRune('9')
	z01.PrintRune('\n')
}
