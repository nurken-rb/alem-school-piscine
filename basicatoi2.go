package piscine

func BasicAtoi2(s string) int {
	res := 0

	for i := 0; i < len(s); i++ {
		if 48 > s[i] || s[i] > 58 {
			return 0
		}
	}

	for i := 0; i < len(s); i++ {
		res *= 10
		res += int(s[i]) - 48
	}
	return res
}
