package piscine

func issPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	if nb == 2 {
		return true
	}
	if nb%2 == 0 {
		return false
	}
	for i := 3; i <= nb/3; i += 2 {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func FindNextPrime(nb int) int {
	for true {
		if issPrime(nb) {
			return nb
		}
		nb++
	}
	return 0
}
