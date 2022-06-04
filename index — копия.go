package piscine

func contains(s1 string, s2 string) bool {
	for i := 0; i <= len(s1)-len(s2); i++ {
		count := 0
		for j := 0; j < len(s2); j++ {
			if s2[j] == s1[i+j] {
				count++
			}
		}
		if count == len(s2) {
			return true
		}
	}
	return false
}

func Index(s string, toFind string) int {
	if len(s) == 0 || len(toFind) == 0 {
		return 0
	}
	if contains(s, toFind) {
		for i := 0; i < len(s); i++ {
			if s[i] == toFind[0] {
				return i
			}
		}
	}
	return -1
}
