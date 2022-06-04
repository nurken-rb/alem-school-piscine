package piscine

func IsAlpha(s string) bool {
	for i := 0; i < len(s); i++ {
		if !(s[i] > 64 && s[i] < 91 || s[i] > 96 && s[i] < 123 || s[i] > 47 && s[i] < 58) {
			return false
		}
	}
	return true
}
