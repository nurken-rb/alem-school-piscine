package piscine

func Atoi(s string) int {
	if s == "" {
		return 0
	}
	res := 0
	sign := 1
	j := 0

	if s[0] == '-' {
		sign *= -1
		j++
	} else if s[0] == '+' {
		j++
	}
	for i := j; i < len(s); i++ {
		if s[i] < 48 || s[i] > 58 {
			return 0
		}
	}
	for ; j < len(s); j++ {
		res *= 10
		res += int(s[j]) - 48
	}
	return res * sign
}
