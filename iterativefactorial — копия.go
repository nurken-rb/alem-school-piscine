package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 2147483647 {
		return 0
	}
	fact := 1
	for i := 1; i <= nb; i++ {
		fact *= i
	}
	return fact
}
