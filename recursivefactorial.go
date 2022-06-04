package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 || nb > 214748364 {
		return 0
	} else if nb == 1 || nb == 0 {
		return 1
	} else {
		return RecursiveFactorial(nb-1) * nb
	}
}
