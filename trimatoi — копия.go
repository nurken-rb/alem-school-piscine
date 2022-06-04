package piscine

func TrimAtoi(s string) int {
	res := 0
	sign := 1

	if len(s) == 0 {
		return 0
	}

	for i := 0; !(47 < s[i] && s[i] < 58); i++ {
		if i == len(s)-1 {
			break
		}
		if s[i] == '-' {
			sign *= -1
			break
		}
	}

	for i := 0; i < len(s); i++ {
		if 47 < s[i] && s[i] < 58 {
			res *= 10
			res += int(s[i]) - 48
		}
	}

	return res * sign
}
