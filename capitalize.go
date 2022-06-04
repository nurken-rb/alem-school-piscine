package piscine

func isAlpha(s rune) bool {
	if !(s > 64 && s < 91 || s > 96 && s < 123 || s > 47 && s < 58) {
		return false
	}
	return true
}

func Capitalize(s string) string {
	r := []rune(s)
	if 96 < r[0] && r[0] < 123 {
		r[0] -= 32
	}
	for i := 1; i < len(r); i++ {
		if 96 < r[i] && r[i] < 123 && !isAlpha(r[i-1]) {
			r[i] -= 32
		} else if 64 < r[i] && r[i] < 91 && isAlpha(r[i-1]) {
			r[i] += 32
		}
	}
	return string(r)
}
