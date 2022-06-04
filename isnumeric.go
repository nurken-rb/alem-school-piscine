package piscine

func IsNumeric(s string) bool {
	for i := 0; i < len(s); i++ {
		if 48 > s[i] || s[i] > 57 {
			return false
		}
	}
	return true
}
