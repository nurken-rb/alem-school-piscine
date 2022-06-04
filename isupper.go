package piscine

func IsUpper(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] < 64 || s[i] > 90 {
			return false
		}
	}
	return true
}
