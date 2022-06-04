package piscine

func BasicAtoi(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		res *= 10
		res += int(s[i]) - 48
	}
	return res
}
