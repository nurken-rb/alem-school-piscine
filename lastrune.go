package piscine

func LastRune(s string) rune {
	r := []rune(s)
	return r[len(s)-1]
}
